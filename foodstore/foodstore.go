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

var Store FoodStore

func New() {
	Store = make(FoodStore)
}

// Add f2 to f1. Modify f1.
func add(f1 *Food, f2 Food) {
	f1.Price += f2.Price
	f1.Fiber += f2.Fiber
	f1.Calories += f2.Calories
	f1.Sugar += f2.Sugar
	f1.Fat += f2.Fat
	f1.Sodium += f2.Sodium
	f1.Carbohydrates += f2.Carbohydrates
	f1.Protein += f2.Protein
	f1.Cholesterol += f2.Cholesterol
}

func multiply(name string, times int) Food {
	f := Store[name]
	return Food{
		f.Name,
		f.Price * float64(times),
		f.Fiber * times,
		f.Calories * times,
		f.Sugar * times,
		f.Fat * times,
		f.Sodium * times,
		f.Carbohydrates * times,
		times,
		f.Cholesterol * times,
		f.Protein * times,
	}
}


func NewFromDb() {
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
	New()
	for rows.Next() {
		var name string
		var price float64
		var fiber, calories, sugar, fat, sodium,
			carbohydrates, serving, cholesterol, protein int
		rows.Scan(&name, &price, &fiber, &calories, &sugar, &fat,
			&sodium, &carbohydrates, &serving, &cholesterol,
			&protein)
		Store[name] = Food{
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
