package main

import "net/http"

func main() {
	http.HandleFunc("/", FindCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func FindCEPHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
