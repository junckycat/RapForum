package main

import (
	"html/template"
	"net/http"
	"os"
)

type User struct {
	ID int64

	Name string

	Lastname string

	Email string

	Born string

	Country string

	City string

	Postal string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("front/index.html"))
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("front/login.html"))
		tmpl.Execute(w, nil)
	})

	root, _ := os.Getwd()
	sf := http.FileServer(http.Dir(root + "/front"))
	http.Handle("/static/", http.StripPrefix("/static/", sf))

	http.ListenAndServe(":8000", nil)
}
