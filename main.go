package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my homepage!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<p>To keep in touch, please send an email to pinta@pinta.com</p>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Page not found</h1>")
	}
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
