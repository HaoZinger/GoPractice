package main

import (
	"fmt"
	"sort"
)

type IntArray []int

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("{Person: name=%s;age=%d }", p.name, p.age)
}

type PersonArr []Person

func (arr PersonArr) Len() int {
	return len(arr)
}

func (arr PersonArr) Less(i, j int) bool {
	return arr[i].age < arr[j].age
}

func (arr PersonArr) Swap(i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	arr := IntArray{1, 2, 3, 1, 8, 456, 8, 34, 67, 34}
	Sort(arr)
	for _, i := range arr {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	arrPerson := PersonArr{{"tom", 123}, {"jerry", 22}, {"mike", 2}}
	Sort(arrPerson)
	for _, person := range arrPerson {
		fmt.Printf("%v ", person)
	}
	sort.Sort(arr)

}
