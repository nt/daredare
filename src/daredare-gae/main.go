package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", redirHandler)
	http.HandleFunc("/updates", indexHandler)
	appengine.Main()
}

func redirHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/posts/paccup.html", http.StatusFound)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "work in progress, should be ready before we leave ^^")
}
