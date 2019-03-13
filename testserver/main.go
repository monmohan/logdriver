package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stdout, "Received request: ", r.URL.Path)

		for hk, hv := range r.Header {
			fmt.Fprintf(os.Stdout, "%s=%s\n", hk, hv)
		}
		fmt.Fprintf(w, "echo %v", r.URL.Path)

	})
	var port string
	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "9000"
	}
	log.Printf("Starting server on host %v  and port %v", "localhost", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
