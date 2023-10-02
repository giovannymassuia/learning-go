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
	t := template.Must(template.New("content.html").ParseFiles("content.html"))
	err := t.Execute(os.Stdout, Courses{
		{"Go", 20},
		{"Python", 30},
		{"JavaScript", 40},
		{"Java", 50},
	})
	if err != nil {
		panic(err)
	}
}
