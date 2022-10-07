package server

import (
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
func (s *Server) Serve(dir, port string) error {

	s.Port = port
	s.Dir = dir

	http.HandleFunc("/healthz", HealthCheckHandler)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}
