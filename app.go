package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", app)
	http.ListenAndServe(":3000", nil)
}

func app(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
