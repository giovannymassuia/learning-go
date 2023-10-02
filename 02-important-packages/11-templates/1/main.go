package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name       string
	TotalHours int
}

func main() {
	curso := Course{Name: "Go", TotalHours: 20}
	t := template.New("CourseTemplate")
	t, _ = t.Parse("Course: {{.Name}} | Hours: {{.TotalHours}}")
	err := t.ExecuteTemplate(os.Stdout, "CourseTemplate", curso)
	if err != nil {
		panic(err)
	}
}
