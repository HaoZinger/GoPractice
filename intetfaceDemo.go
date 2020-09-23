package main

import (
	"fmt"
)

func main() {
	TypeSwitch("string")
	TypeSwitch(float32(12.0))

	bird := new(Bird)
	bird.walk()
	bird.name= "2"
	fmt.Print(bird)

}

type myStr string

var str myStr = "hello"

func TypeSwitch(any interface{}) {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case string:
			println("String")
		case bool:
			println("bool")
		case int, float32:
			fmt.Print(v)
		case myStr:
			println("myStr")
		default:
			println("not match")
		}
	}
	testFunc(any)
}


type IDuck interface {
	walk()
	swing()
}

type Bird struct {
	name  string
	color string
}

func (*Bird) walk() {
	println("bird walk")
}
