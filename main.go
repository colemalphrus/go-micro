package main

import (
	"fmt"
	"net/http"

	"github.com/colemalphrus/go-micro/controllers"
	"github.com/gorilla/mux"
)


func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Homepage)
	fmt.Println("running server")
	
	fs := http.FileServer(http.Dir("public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	http.ListenAndServe(":8080", r)
	fmt.Println("closing")
}