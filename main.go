package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// * Finish HTML inputs with tailwind and then use htmx.

type Hero struct {
	Name, Superpower string
}

func main() {

	heroes := map[string][]Hero{
		"Heroes": {
			{
				Name:       "Superman",
				Superpower: "Super strength",
			},
			{
				Name:       "Flash",
				Superpower: "Super speed",
			},
		},
	}
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, heroes)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		name := r.PostFormValue("name")
		sp := r.PostFormValue("superpower")

		htmlstr := fmt.Sprintf("<p>%s - %s</p>", name, sp)
		tmpl, err := template.New("t").Parse(htmlstr)

		if err != nil {
			log.Println("error during parse the values in html str")
		}

		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", handler1)
	http.HandleFunc("/add-hero", handler2)

	log.Default().Println("Running server on PORT http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
