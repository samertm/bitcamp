package main

import (
	"github.com/samertm/bitcamp/foodstore"
	"github.com/samertm/bitcamp/server"
)

func main() {
	f := foodstore.NewFromDb()
        server.ListenAndServe(":8000", f)
}
