package server

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"html/template"
	"log"
	"net/http"

	"github.com/samertm/bitcamp/foodstore"
)

func serveRoot(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

type FoodWrapper struct {
	Foods []foodstore.Food
}

func serveFood(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, _ := ioutil.ReadAll(req.Body)
		num, err := strconv.ParseInt(string(data), 10, 0)
		if err != nil {
			fmt.Fprintf(w, "error: not a number.")
			return
		}
		foods := foodstore.OptimalFoods(int(num))
		wrapper := FoodWrapper{foods}
		t, _ := template.ParseFiles("templates/food.html")
		t.Execute(w, wrapper)
	} else {
		serveRoot(w, req)
	}
}

func serveAbout(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/about.html")
	t.Execute(w, nil)
}

func serveContact(w http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/contact.html")
	t.Execute(w, nil)
}

func ListenAndServe(ip string) {
	fmt.Printf("Running server on %s.\n", ip)
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/food/", serveFood)
	http.HandleFunc("/about/", serveAbout)
	http.HandleFunc("/contact/", serveContact)
	log.Fatal(http.ListenAndServe(ip, nil))
}
