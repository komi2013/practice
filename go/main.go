package main

import (
	// "html/template"
	// "fmt"
	"log"
	"net/http"

	"./app"
	// "./common"
)

func main() {
	http.HandleFunc("/test/", app.Test)
	log.Fatal(http.ListenAndServe(":9002", nil))
}
