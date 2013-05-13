package controllers

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

func renderTemplate(w http.ResponseWriter, name string, temp string) {
	tmpl, err := template.New(name).Parse(temp)
	if err != nil {

	}
	tmpl.ExecuteTemplate(w, name, temp)
}

func Index(rw http.ResponseWriter, req *http.Request) {
	/*
		data, count, err := ReadHtml("views/index.html")
		if err != nil {
			rw.Write([]byte("Error!!"))
		}
		rw.Write(data[:count])
	*/
	t, err := template.ParseFiles("views/index.html")
	if err != nil {

	}
	t.Execute(rw, &User{"admin"})
}

func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("user login page"))
}
