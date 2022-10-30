package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/cmwylie19/kubefs/pkg/utils"
	"go.uber.org/zap"
)

func EnableCors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		f(w, r)
	}
}

var wg sync.WaitGroup

type File struct {
	Name string
	Date string
	Path string
}
type Server struct {
	Port string
	Dir  string
}

func (s *Server) CascadeDelete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wg.Add(1)

	indexes, err := queryToRange(req.URL.Query())
	if err != nil {
		utils.Logger.Error("Error parsing query", zap.Error(err))
	}

	left, err := fileNameToInt(indexes[0])
	if err != nil {
		utils.Logger.Error("Error parsing query left parameter", zap.Error(err))
	}

	right, err := fileNameToInt(indexes[1])
	if err != nil {
		utils.Logger.Error("Error parsing query right parameter", zap.Error(err))
	}
	utils.Logger.Info("Cascade Delete", zap.Int("left", left), zap.Int("right", right))
	filechan := make(chan []File)
	go cascadeDelete(s.Dir, left, right, filechan)
	deletedFiles := <-filechan
	wg.Wait()
	utils.Logger.Info("Cascade Deletion", zap.Int("number", len(deletedFiles)))
	json.NewEncoder(w).Encode(deletedFiles)
}

func cascadeDelete(dir string, left, right int, filechan chan []File) {
	defer wg.Done()
	var count int

	file_array := []File{}
	files := make(chan []File)
	go listFiles(dir, files)
	file_list := <-files

	utils.Logger.Info("cascadeDelete - listed files", zap.Int("file count", len(file_list)))

	// find files between left and right

	for _, file := range file_list {
		// convert file name to int
		file_num, err := fileNameToInt(file.Name)
		if err != nil {
			utils.Logger.Error("Error parsing file name", zap.Error(err), zap.String("file", file.Name))
			return
		}

		if file_num >= left && file_num <= right {
			count++
			// delete file
			wg.Add(1)
			done := make(chan bool)
			go deleteFile(file.Path, done)
			<-done
			wg.Wait()
			file_array = append(file_array, file)
		}
	}

	utils.Logger.Info("Files Deleted: ", zap.Int("count", count))
	filechan <- file_array

}

func queryToRange(query url.Values) ([]string, error) {
	begin, found := query["begin"]
	if !found {
		utils.Logger.Fatal("No begin parameter found")
		return []string{}, fmt.Errorf("No begin parameter found")
	}

	end, found := query["end"]
	if !found {
		utils.Logger.Fatal("No end parameter found")
		return []string{}, fmt.Errorf("No end parameter found")
	}

	return []string{begin[0], end[0]}, nil
}
func fileNameToInt(name string) (int, error) {
	// Remove the file extension
	name = name[1 : len(name)-4]

	// Convert the file name to an int
	return strconv.Atoi(name)
}

func deleteFile(path string, done chan bool) {
	defer wg.Done()
	utils.Logger.Info("Delete File", zap.String("path", path))
	err := os.Remove(path)
	if err != nil {
		utils.Logger.Error("Error deleting file", zap.String("path", path), zap.Error(err))
	}
	utils.Logger.Info("Deleted File", zap.String("path", path))
	done <- true
}
func (s *Server) DeleteFile(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	wg.Add(1)

	done := make(chan bool)
	go deleteFile(req.URL.Path[12:], done)
	<-done
	wg.Wait()
	utils.Logger.Info("Deleted file", zap.String("file", req.URL.Path[12:]))
	fs := fmt.Sprintf(`{"deleted":{"file":"%s"}}`, s.Dir+req.URL.Path[12:])
	io.WriteString(w, fs)
}

// key, cert, port, dir string for future use
func (s *Server) Serve(key, cert, dir, port string) error {

	s.Port = port
	s.Dir = dir
	http.Handle("/", http.FileServer(http.Dir("media")))
	http.HandleFunc("/healthz", EnableCors(HealthCheckHandler))
	http.HandleFunc("/delete/file/", EnableCors(s.DeleteFile))
	http.HandleFunc("/delete/cascade", EnableCors(s.CascadeDelete))
	http.HandleFunc("/list", EnableCors(s.ListFiles))

	if key != "" && cert != "" {
		// fmt.Printf("Starting server at %s watching directory %s.\n", port, dir)
		utils.Logger.Info("Starting server", zap.String("port", port), zap.String("directory", dir))
		err := http.ListenAndServeTLS(":"+port, cert, key, nil)
		if err != nil {
			utils.Logger.Fatal("ListenAndServe: ", zap.Error(err))
			return err
		}
	} else {
		// fmt.Printf("Starting server at %s watching directory %s.\n", port, dir)
		utils.Logger.Info("Starting server", zap.String("port", port), zap.String("directory", dir))
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			utils.Logger.Fatal("ListenAndServe: ", zap.Error(err))
			return err
		}
	}
	return nil
}

func listFiles(dir string, file_chan chan []File) {
	//elements := make(map[string]interface{})
	file_array := []File{}
	defer wg.Done()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			utils.Logger.Fatal("Error reading directory", zap.Error(err))
		}
		if !info.IsDir() {
			file_array = append(file_array, File{
				Name: info.Name(),
				Date: info.ModTime().String(),
				Path: path,
			})
			// utils.Logger.Info("list File", zap.String("name", info.Name()), zap.String("date", info.ModTime().String()), zap.String("path", path))
		}
		return nil
	})
	if err != nil {
		utils.Logger.Fatal("Error reading directory", zap.Error(err))
	}

	file_chan <- file_array
}

// List files in the designated directory
// Assumes the server has access to the files
// in Kubernetes via a persistent volume with a hostpath mount
func (s *Server) ListFiles(w http.ResponseWriter, r *http.Request) {
	utils.Logger.Info("List Files")
	files := make(chan []File)
	w.Header().Set("Content-Type", "application/json")
	wg.Add(1)

	go listFiles(s.Dir, files)
	file_list := <-files
	wg.Wait()
	json.NewEncoder(w).Encode(file_list)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if os.Getenv("ENV") != "CI" {
		utils.Logger.Info("Health check", zap.String("status", "OK"))
	}

	io.WriteString(w, `{"alive": true}`)
}
