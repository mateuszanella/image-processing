package main

import (
	"fmt"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	red, green, blue, err := HexToRGB("#ff0000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(red, green, blue)

	// tmpl := template.Must(template.ParseFiles("internal/templates/index.html"))
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	data := TodoPageData{
	// 		PageTitle: "My TODO list",
	// 		Todos: []Todo{
	// 			{Title: "Task 1", Done: false},
	// 			{Title: "Task 2", Done: true},
	// 			{Title: "Task 3", Done: true},
	// 		},
	// 	}
	// 	tmpl.Execute(w, data)
	// })
	// http.ListenAndServe(":80", nil)

	// http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	// http.ListenAndServe(":8080", nil)
}
