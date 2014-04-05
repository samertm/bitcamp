package foodstore

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"log"
)

type Food struct {
	Name string
	// All of the following are per serving. Sizes are in grams.
	Price         float
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

func PopulateFromDb() {
	dbName := "./foods.db"
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
	store := New()
	for rows.Next() {
		var name string
		var price float
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
}

func NewDebug() FoodStore {
	f := make(FoodStore)
	f["rice"] = Food{
		"rice",
		3.0,
		202,
		211,
		0,
		0,
		209,
		5500,
		2000,
		100,
		5000,
	}
	f["beans"] = Food{
		"beans",
		2.0,
		502,
		511,
		0,
		0,
		509,
		3500,
		1000,
		500,
		6000,
	}
	return f
}
