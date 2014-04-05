package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/samertm/bitcamp/foodstore"
)

var store foodstore.FoodStore

func serveRoot(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func serveFood(w http.ResponseWriter, req *http.Request) {
	foods := runAlgorithm(15)
	t, _ := template.ParseFiles("templates/food.html")
	t.Execute(w, foods)
}

func serveAbout(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/about.html")
	t.Execute(w, nil)
}

func serveContact(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/contact.html")
	t.Execute(w, nil)
}

func runAlgorithm(money int) []foodstore.Food {
	return []foodstore.Food{
		store["white rice"],
		store["brown rice"],
	}
}
	
func ListenAndServe(ip string, f foodstore.FoodStore) {
	store = f
	fmt.Printf("Running server on %s.\n", ip)
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/food/", serveFood)
	http.HandleFunc("/about/", serveAbout)
	http.HandleFunc("/contact/", serveContact)
	log.Fatal(http.ListenAndServe(ip, nil))
}
