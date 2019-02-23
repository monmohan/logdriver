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
	log.Fatal(http.ListenAndServe(":9090", nil))
}
