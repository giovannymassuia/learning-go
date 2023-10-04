package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("server: hello handler started")
	defer log.Println("server: hello handler ended")

	select {
	case <-time.After(5 * time.Second):
		log.Println("server: hello handler: request processed")
		w.Write([]byte("Request processed"))
	case <-ctx.Done():
		log.Println("server: hello handler: request cancelled")
		http.Error(w, "Request cancelled by the client", http.StatusRequestTimeout)
	}
}
