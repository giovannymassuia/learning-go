package main

import "net/http"

func main() {
	// multiplexer
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
