package main

import (
	"./parse"
	"fmt"
)

func main() {
	strs := []string{"1 2 3", "2", "tom", " 1 2 a", "2 3 4",""}
	for index, ele := range strs {
		numbers, err := parse.Parse(ele)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("element %d : %v\n", index, numbers)
	}
}
