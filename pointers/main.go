package main

import "fmt"

func main() {

	array := ArrayList{}
	add(0, "teste", &array)
	add(1, "teste2", &array)
	array.add(2, "teste3")
	addReferenceAndMethod(3, "teste4", &array)

	fmt.Println(array)
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
