package server

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type File struct {
	Name string
	Date string
}
type Server struct {
	Port string
	Dir  string
}

// TODO
// func (s *Server) DeleteFile() error {

// }

// key, cert, port, dir string for future use
func (s *Server) Serve(key, cert, dir, port string) error {

	s.Port = port
	s.Dir = dir

	http.HandleFunc("/healthz", HealthCheckHandler)
	http.HandleFunc("/list", s.ListFiles)

	if key != "" && cert != "" {
		fmt.Printf("Starting server at %s watching directory %s.\n", port, dir)
		err := http.ListenAndServeTLS(":"+port, cert, key, nil)
		if err != nil {

			log.Fatal("ListenAndServe: ", err)
			return err
		}
	} else {
		fmt.Printf("Starting server at %s watching directory %s.\n", port, dir)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)

			return err
		}
	}
	return nil
}

func listFiles(dir string, file_chan chan []File) {
	//elements := make(map[string]interface{})
	file_array := []File{}
	defer wg.Done()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		file_array = append(file_array, File{
			Name: file.Name(),
		})
	}
	file_chan <- file_array
}

// List files in the designated directory
// Assumes the server has access to the files
// in Kubernetes via a persistent volume with a hostpath mount
func (s *Server) ListFiles(w http.ResponseWriter, r *http.Request) {
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

	io.WriteString(w, `{"alive": true}`)
}
