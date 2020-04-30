package main

import (
	"fmt"
)

type Any interface{}

type Car struct {
	Model       string
	Manufacture string
}

type Cars []*Car

func (cs Cars) Process(f func(car *Car)) {
	for _, car := range cs {
		f(car)
	}
}

func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0)
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufacture []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)

	for _, m := range manufacture {
		sortedCars[m] = make([]*Car, 0)
	}

	sortedCars["Default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacture]; ok {
			sortedCars[c.Manufacture] = append(sortedCars[c.Manufacture], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}

	return appender, sortedCars
}

func main() {
	ford := &Car{"Fiesta", "Ford"}
	bmw := &Car{"XL 450", "BMW"}
	merc := &Car{"D600", "Mercedes"}
	bmw2 := &Car{"X 800", "BMW"}

	allCars := Cars([]*Car{ford, bmw, merc, bmw2})

	brand := "Ford"

	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacture == brand)
	})

	fmt.Println("deleted Cars: ", allCars[])

	fmt.Println("All Cars: ", allCars)
	fmt.Println("New BMWs: ", allNewBMWs)

	manufactures := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaugar"}

	sortedAppender, sortedCars := MakeSortedAppender(manufactures)
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, "BMWs")
}

// func search(brand string, allCars []*Car) {
// 	brandFound := allCars.FindAll(func(car *Car) bool {
// 		return (car.Manufacture == brand)
// 	})
// 	fmt.Println("Found Cars: ", brandFound)

// }
