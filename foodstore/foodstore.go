package foodstore

type Food struct {
	Name string
	// Price per serving
	Price float
	// Serving in grams
	ServingSize float
	// Calories per serving
	Calories int
	// All the rest of these are per serving
	// in mg (milligrams)
	Fat          int
	SaturatedFat int
	TransFat     int
	Cholesterol  int
	DietaryFiber int
	Sugar        int
	Protein      int
}

type FoodStore map[string]Food

func New() FoodStore {
	f := make(FoodStore)
	return f
}

func NewDebug () FoodStore {
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
