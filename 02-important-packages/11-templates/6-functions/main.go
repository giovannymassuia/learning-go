package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name       string
	TotalHours int
}

type Courses []Course

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"MyToUpper": toUpper,
	})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		Course{"Go", 20},
		Course{"Python", 30},
		Course{"JavaScript", 40},
	})
	if err != nil {
		panic(err)
	}
}
