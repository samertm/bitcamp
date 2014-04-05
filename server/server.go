package server

import (
	"net/http"
	"fmt"
)

type Snake int

func serveRoot(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Whatever yoooooo get well soon."))
}

func ListenAndServe(ip string) {
	fmt.Printf("Running server on %s.\n", ip)
	http.HandleFunc("/", serveRoot)
	http.ListenAndServe(ip, nil)
}
