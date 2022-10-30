/*
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func QueryCheck(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Printf("%+v", query)
	fmt.Printf("%T", query)
	fmt.Println("\n\nleft: ", r.URL.Query()["left"][0])
	left, present := query["left"]
	if !present || len(left) == 0 {
		fmt.Println("left not present")
	}

	right, present := query["right"]
	if !present || len(right) == 0 {
		fmt.Println("right not present")
	}

	w.Write([]byte(left[0] + "," + right[0]))
}
func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	// http.Handle("/", http.FileServer(http.Dir(*directory)))
	http.HandleFunc("/query", QueryCheck)
	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
