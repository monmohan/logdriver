package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stdout, "Receieved request: ", r.URL.Path)

		for hk, hv := range r.Header {
			fmt.Fprintln(os.Stdout, "Header", hk, hv)
		}
		fmt.Fprintf(w, "echo %v", r.URL.Path)

	})
	log.Printf("Starting server on host %v  and port %v", "localhost", 9000)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
