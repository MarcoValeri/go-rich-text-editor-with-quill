package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home() {
	tmpl := template.Must(template.ParseFiles("./views/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Get data from the form
		getFormSubmit := r.FormValue("form-submit")
		getContent := r.FormValue("form-content")

		/**
		* Print out in the terminal
		* the form-content
		* if the form has been submitted
		 */
		if getFormSubmit == "Send your content" {
			fmt.Println(getContent)
		}

		tmpl.Execute(w, nil)
	})
}

func main() {
	// Set static files under public folder
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	home()
	http.ListenAndServe(":80", nil)
}
