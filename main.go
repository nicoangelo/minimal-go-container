package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const PORT = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hello from %s", hostname)
	})

	log.Printf("Starting webserver on port %d", PORT)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
