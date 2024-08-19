package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = ":8083"
const tmpl_page = "../../templates/home_page.html"

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles(tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Panic("Error parsing template html page...")
		return
	}
}

func main() {

	fmt.Printf("MICRO SERVICES IS STARTED!!! AT LOCALHOST PORT %s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, tmpl_page)
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
