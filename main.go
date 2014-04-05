package main

import (
	"bitcamp/foodstore"
	"bitcamp/server"
)

func main() {
	f := foodstore.NewDebug()
	server.ListenAndServe(":8000")
}
