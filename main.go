package main

import (
	"github.com/samertm/bitcamp/foodstore"
	_ "github.com/samertm/bitcamp/server"

	//DEBUG
	"fmt"
)

func main() {
	foodstore.NewFromDb() // TODO change to init
	fmt.Println(foodstore.OptimalFoods(10.0))
	fmt.Println(foodstore.Sortedfoods.Calories)
        //server.ListenAndServe(":8889", f)
}
