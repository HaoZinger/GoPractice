package main

import (
	"fmt"
	"time"
)

type Car struct {
	buildYear int
	color     string
}

type Cars []Car

func (cars Cars) process(f func(c Car)) {
	for _, each := range cars {
		f(each)
	}
}

func printCar(car Car) {
	fmt.Println(car, time.Now())
}

func (cars Cars) findAll(f func(c Car) bool) Cars {
	filters := make([]Car, 0)
	for _, car := range cars {
		if f(car) == true {
			filters = append(filters, car)
		}
	}
	return filters
}

func main() {
	cars := Cars{{1, "red"}, {2, "white"}, {5, "blue"}}
	cars.process(printCar)
	all := cars.findAll(func(c Car) bool {
		return c.buildYear < 3
	})
	fmt.Println(all)
}
