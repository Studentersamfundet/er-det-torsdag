package main

import (
	"html/template"
	"net/http"
	"time"
)

func main() {
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", mainHandler(templ))
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("./imgs"))))

	err = http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		panic(err)
	}
}

func mainHandler(templ *template.Template) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		now := time.Now()

		err := templ.Execute(w, struct {
			Torsdag bool
			Patrick bool
		}{
			Torsdag: now.Weekday() == time.Thursday,
			Patrick: now.Month() == time.March && now.Day() == 17,
		})
		if err != nil {
			panic(err)
		}
	}
}
