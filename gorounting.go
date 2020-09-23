package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//println(sum(4))
	//demo1()
	//demo2()
	demo3()
}

func demo1() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	go pump1(ch1)
	go pump2(ch2)
	sink(ch1, ch2)
	time.Sleep(1E6)
}

func sum(value int) (res int) {
	res = 0
	for i := 0; i < value; i++ {
		res += i
		fmt.Printf("sum [0~%d] is %d \n", i, res)
	}
	return
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 5
	}
}

func sink(ch1, ch2 chan int) {
	for true {
		select {
		case v := <-ch1:
			fmt.Printf("in case1 : %d \n", v)
		case v := <-ch2:
			fmt.Printf("in 2 : %d \n", v)
		}
	}
}

func demo2() {
	evalFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	evaluator := BuildLazyIntEvaluator(evalFunc, 0)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d th even: %d\n", i, evaluator())
	}
	fmt.Print("--------\n")
	fibonacciEvalFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		first := os[0]
		second := os[1]
		ns := []uint64{second, first + second}
		return first, ns
	}
	fibonacciEvaluator := BuildLazyFibonacciEvaluator(fibonacciEvalFunc, []uint64{0, 1})
	for i := 0; i < 10; i++ {
		fmt.Printf("%d th fibonacci is %d\n", i, fibonacciEvaluator())
	}

}

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChain := make(chan Any)
	loopFunc := func() {
		var initValue = initState
		var retVal Any
		for {
			retVal, initValue = evalFunc(initValue)
			retValChain <- retVal
		}
	}
	go loopFunc()
	retFunc := func() Any {
		return <-retValChain
	}
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	evaluator := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return evaluator().(int)
	}
}

func BuildLazyFibonacciEvaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	evaluator := BuildLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return evaluator().(uint64)
	}
}

func demo3() {
	flag.Parse()
	leftmost := make(chan int)
	left := make(chan int)
	var right = leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	fmt.Println(x)  // 100000, ongeveer 1,5 s
}

var ngoroutine = flag.Int("n", 90000, "how many goroutines")

func f(left, right chan int) {
	i := 1 + <-right
	fmt.Printf("in func f():%d\n", i)
	left <- i
}
