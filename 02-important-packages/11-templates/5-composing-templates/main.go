package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		Course{"Go", 20},
		Course{"Python", 30},
		Course{"JavaScript", 40},
	})
	if err != nil {
		panic(err)
	}
}
