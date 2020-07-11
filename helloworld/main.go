package main

import (
	"fmt"
	"reflect"
)

const (
	text = "GO const"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println(text)

	var text string = "text"
	fmt.Println(text)

	var number = 10.5
	fmt.Println(number)
	fmt.Println("number type:", reflect.TypeOf(number))

}
