package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// var wg sync.WaitGroup

type Server struct {
	Port string
	Dir  string
}

// key, cert, port, dir string for future use
func (s *Server) Serve(key, cert, dir, port string) error {

	s.Port = port
	s.Dir = dir

	http.HandleFunc("/healthz", HealthCheckHandler)

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

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
