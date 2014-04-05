package foodstore

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Food struct {
	Name string
	// All of the following are per serving. Sizes are in grams.
	Price         float64
	Fiber         int
	Calories      int
	Sugar         int
	Fat           int
	Sodium        int
	Carbohydrates int
	Serving       int
	Cholesterol   int
	Protein       int
}

// Map from food names to food. The key and Name field are
// always equal.
type FoodStore map[string]Food

func New() FoodStore {
	f := make(FoodStore)
	return f
}

func NewFromDb() FoodStore {
	dbName := "./apifetcher/foods.db"
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(`
select name, price, fiber, calories, sugar, fat, sodium, carbohydrates, serving, cholesterol, protein from foods
`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var store FoodStore = New()
	for rows.Next() {
		var name string
		var price float64
		var fiber, calories, sugar, fat, sodium,
			carbohydrates, serving, cholesterol, protein int
		rows.Scan(&name, &price, &fiber, &calories, &sugar, &fat,
			&sodium, &carbohydrates, &serving, &cholesterol,
			&protein)
		store[name] = Food{
			name,
			price,
			fiber,
			calories,
			sugar,
			fat,
			sodium,
			carbohydrates,
			serving,
			cholesterol,
			protein,
		}
	}
	return store
}
