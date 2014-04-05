package server

import (
	"net/http"
	"html/template"
	"fmt"
)

func serveRoot(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func serveFood(w http.ResponseWriter, req *http.Request) {
	
}

func ListenAndServe(ip string) {
	fmt.Printf("Running server on %s.\n", ip)
	http.HandleFunc("/", serveRoot)
	http.ListenAndServe(ip, nil)
}
