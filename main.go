package main

import (
	"github.com/samertm/bitcamp/foodstore"
	"github.com/samertm/bitcamp/server"
)

func main() {
	foodstore.NewFromDb() // TODO change to init
        server.ListenAndServe(":8889")
}
