package main

import (
	"fmt"
	"log"
)

func main() {
	var p *int //<nil>
	fmt.Println(p)

	i := 10
	p = &i
	i = 20
	fmt.Println(*p)

	*p = 30
	fmt.Println(*p)

	recoverInvalidMemoryAddress()

	array := ArrayList{}
	add(0, "teste", &array)
	add(1, "teste2", &array)
	array.add(2, "teste3")
	addReferenceAndMethod(3, "teste4", &array)

	fmt.Println(array)
}

func recoverInvalidMemoryAddress() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic recover:", err)
		}
	}()

	var p *int
	var i int

	i = *p
	fmt.Println(i)
}

type ArrayList struct {
	values [10]string
}

func (list *ArrayList) add(index int, value string) {
	list.values[index] = value
}

func add(index int, value string, list *ArrayList) {
	list.values[index] = value
}

func addReferenceAndMethod(index int, value string, list *ArrayList) {
	list.add(index, value)
}
