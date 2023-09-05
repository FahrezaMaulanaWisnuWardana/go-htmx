package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	homePage := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Film A", Director: "Director A"},
				{Title: "Film B", Director: "Director B"},
				{Title: "Film C", Director: "Director C"},
			},
		}
		tmpl.Execute(w, films)
	}
	addFilm := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/add-film/", addFilm)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
