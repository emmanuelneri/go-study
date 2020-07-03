package main

import (
	"fmt"
	"reflect"
)

func main() {
	const helloworld = "Hello World"
	fmt.Println(helloworld)

	var text string = "text"
	fmt.Println(text)

	var number = 10.5
	fmt.Println(number)
	fmt.Println("number type:", reflect.TypeOf(number))

}
