package main

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles("content.html"))
		err := t.Execute(w, Courses{
			{"Go", 20},
			{"Python", 30},
			{"JavaScript", 40},
			{"Java", 50},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
