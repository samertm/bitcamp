package foodstore

import (
	"sort"
)

var constraints = Food{
	Name:          "constraint",
	Fiber:         266000,
	Calories:      14000,
	Sugar:         0,
	Fat:           420000,
	Sodium:        16100,
	Carbohydrates: 1750000,
	Protein:       455000,
}

type ByCalories []string

func (b ByCalories) Len() int      { return len(b) }
func (b ByCalories) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByCalories) Less(i, j int) bool {
	return ((float64(Store[b[i]].Calories) / Store[b[i]].Price) >
		(float64(Store[b[j]].Calories) / Store[b[j]].Price))
}

type ByProtein []string

func (b ByProtein) Len() int      { return len(b) }
func (b ByProtein) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByProtein) Less(i, j int) bool {
	return ((float64(Store[b[i]].Protein) / Store[b[i]].Price) >
		(float64(Store[b[j]].Protein) / Store[b[j]].Price))
}

type ByCarbohydrates []string

func (b ByCarbohydrates) Len() int      { return len(b) }
func (b ByCarbohydrates) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByCarbohydrates) Less(i, j int) bool {
	return ((float64(Store[b[i]].Carbohydrates) / Store[b[i]].Price) >
		(float64(Store[b[j]].Carbohydrates) / Store[b[j]].Price))
}

type ByFat []string

func (b ByFat) Len() int      { return len(b) }
func (b ByFat) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByFat) Less(i, j int) bool {
	return ((float64(Store[b[i]].Fat) / Store[b[i]].Price) >
		(float64(Store[b[j]].Fat) / Store[b[j]].Price))
}

type ByFiber []string

func (b ByFiber) Len() int      { return len(b) }
func (b ByFiber) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByFiber) Less(i, j int) bool {
	return ((float64(Store[b[i]].Fiber) / Store[b[i]].Price) >
		(float64(Store[b[j]].Fiber) / Store[b[j]].Price))
}

type SortedFoods struct {
	Calories      []string
	Protein       []string
	Carbohydrates []string
	Fat           []string
	Fiber         []string
}

var Sortedfoods SortedFoods

func InitSortedFoods() {
	for key, v := range Store {
		if v.Calories != 0 {
			Sortedfoods.Calories = append(Sortedfoods.Calories, key)
		}
		if v.Protein != 0 {
			Sortedfoods.Protein = append(Sortedfoods.Protein, key)
		}
		if v.Carbohydrates != 0 {
			Sortedfoods.Carbohydrates = append(Sortedfoods.Carbohydrates, key)
		}
		if v.Fat != 0 {
			Sortedfoods.Fat = append(Sortedfoods.Fat, key)
		}
		if v.Fiber != 0 {
			Sortedfoods.Fiber = append(Sortedfoods.Fiber, key)
		}
	}
	sort.Sort(ByCalories(Sortedfoods.Calories))
	sort.Sort(ByProtein(Sortedfoods.Protein))
	sort.Sort(ByCarbohydrates(Sortedfoods.Carbohydrates))
	sort.Sort(ByFat(Sortedfoods.Fat))
	sort.Sort(ByFiber(Sortedfoods.Fiber))
}

func OptimalFoods(money float64) ([]Food, *Food) { // append makeup?
	// number of each food item. Each item is zeroed (yay go :D).
	numfoods := make(map[string]int)
	// the person's nutritional/price makeup
	InitSortedFoods()
	makeup := &Food{}
loop:
	for makeup.Price < money {
		switch {
		case makeup.Calories < constraints.Calories:
			for _, f := range Sortedfoods.Calories {
				if makeup.Calories+Store[f].Calories < constraints.Calories &&
					numfoods[f] < 5 &&
					makeup.Price+Store[f].Price < money {
					add(makeup, Store[f])
					numfoods[f] += 1
					continue loop
				}
			}
			fallthrough
		case makeup.Protein < constraints.Protein:
			for _, f := range Sortedfoods.Protein {
				if makeup.Protein+Store[f].Protein < constraints.Protein &&
					numfoods[f] < 5 &&
					makeup.Price+Store[f].Price < money {
					add(makeup, Store[f])
					numfoods[f] += 1
					continue loop
				}
			}
			fallthrough
		case makeup.Carbohydrates < constraints.Carbohydrates:
			for _, f := range Sortedfoods.Carbohydrates {
				if makeup.Carbohydrates+Store[f].Carbohydrates < constraints.Carbohydrates &&
					numfoods[f] < 5 &&
					makeup.Price+Store[f].Price < money {
					add(makeup, Store[f])
					numfoods[f] += 1
					continue loop
				}
			}
			fallthrough
		case makeup.Fat < constraints.Fat:
			for _, f := range Sortedfoods.Fat {
				if makeup.Fat+Store[f].Fat < constraints.Fat &&
					numfoods[f] < 5 &&
					makeup.Price+Store[f].Price < money {
					add(makeup, Store[f])
					numfoods[f] += 1
					continue loop
				}
			}
			fallthrough
		case makeup.Fiber < constraints.Fiber:
			for _, f := range Sortedfoods.Fiber {
				if makeup.Fiber+Store[f].Fiber < constraints.Fiber &&
					numfoods[f] < 5 &&
					makeup.Price+Store[f].Price < money {
					add(makeup, Store[f])
					numfoods[f] += 1
					continue loop
				}
			}
			fallthrough
		default:
			break loop
		}
	}
	foods := make([]Food, 0, 10)
	for key, value := range numfoods {
		foods = append(foods, multiply(key, value))
	}
	return foods, makeup
}
