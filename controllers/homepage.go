package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
)


func Homepage(w http.ResponseWriter, r *http.Request){
	data := map[string]string{
		"title": "go micro",
		"header": "Welcome to gomicro",
	}

	templateFile, _ := filepath.Abs("src/templates/index.html.mustache")
	view, err := mustache.RenderFile(templateFile, data)
	fmt.Println(err)
	fmt.Fprintf(w, view)
}