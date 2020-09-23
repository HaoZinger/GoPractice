package main

import (
	"fmt"
)

const num = 50

var fibs [num]int

func main() {

	//for i := 0; i < 13; i++ {
	//	Season(i)
	//}
	//fmt.Println(mult_returnval_named(2, 3))

	//fmt.Println(mult_returnval_unnamed(2, 3))

	// fibonacci
	//start := time.Now()
	//fmt.Println(fibonacci(num - 1))
	//fmt.Println(time.Now().Sub(start))
	//
	//start2 := time.Now()
	//fmt.Println(fibonacci_pretty(num - 1))
	//fmt.Println(time.Now().Sub(start2))
	//fmt.Println(fibs)

	b := []int{1, 2, 3, 4, 5, 6}
	s1 := b[:2]
	s2 := b[0:]
	s3 := b[2:4]
	s4 := b[:]
	fmt.Println(s1, s2, s3, s4)

}

func Season(month int) {
	switch month {
	case 3, 4, 5:
		fmt.Println("Spring!")
	case 6, 7, 8:
		fmt.Println("Summer!")
	case 9, 10, 11:
		fmt.Println("Autumn!")
	case 12, 1, 2:
		fmt.Println("Winter!")
	default:
		fmt.Println("Wrong month!")
	}

}

func mult_returnval_named(a int, b int) (sum int, product int, reduce int) {
	//sum = a + b
	//product = a * b
	//reduce = a - b
	sum, product, reduce = a+b, a*b, a-b
	return
}

func mult_returnval_unnamed(a int, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func fibonacci(n int) (s int) {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci_pretty(n int) (s int) {
	if fibs[n] != 0 {
		return fibs[n]
	}

	if n <= 1 {

		fibs[n] = 1
		return 1
	}
	s = fibonacci_pretty(n-1) + fibonacci_pretty(n-2)
	fibs[n] = s
	return
}
