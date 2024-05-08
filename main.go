package main

import (
	"fmt"
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
	fmt.Println("Go app...")

	getFilms := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
				{Title: "Schindler's List", Director: "Steven Spielberg"},
				{Title: "Pulp Fiction", Director: "Quentin Tarantino"},
			},
		}
		tmpl.Execute(w, films)
	}

	addNewFilm := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	http.HandleFunc("/", getFilms)
	http.HandleFunc("/add-film/", addNewFilm)

	log.Fatal(http.ListenAndServe(":8000", nil))

}