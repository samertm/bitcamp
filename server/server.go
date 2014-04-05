package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/samertm/bitcamp/foodstore"
)

func serveRoot(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func serveFood(w http.ResponseWriter, req *http.Request) {

}

func serveAbout(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/about.html")
	t.Execute(w, nil)
}

func serveContact(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/contact.html")
	t.Execute(w, nil)
}

func ListenAndServe(ip string, f foodstore.FoodStore) {
	fmt.Printf("Running server on %s.\n", ip)
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/about/", serveAbout)
	http.HandleFunc("/contact/", serveContact)
	log.Fatal(http.ListenAndServe(ip, nil))
}
