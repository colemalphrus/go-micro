package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request){
	data := map[string]string{
		"title": "go micro",
		"header": "Welcome to gomicro",
	}

	templateFile, _ := filepath.Abs("src/templates/index.html.mustache")
	view, err := mustache.RenderFile(templateFile, data)
	fmt.Println(err)
	fmt.Fprintf(w, view)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", homepage)
	fmt.Println("running server")
	http.ListenAndServe(":8080", r)
	fmt.Println("closing")
}