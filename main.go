package main

import (
	"fmt"
	"net/http"
	"text/template"

	"pinta.com/controllers"
	"pinta.com/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

var notFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h3>Page not found.</h3>")
})

func main() {
	homeView = views.NewView("views/home.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.NotFoundHandler = notFoundHandler
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	http.ListenAndServe(":8080", r)
}
