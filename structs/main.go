package main

import (
	"fmt"
	"reflect"
)

func main() {

	cellphone := Product{"Galaxy", "Galaxy s10", 2999}
	fmt.Println(cellphone)
	fmt.Println("Cellphone type:", reflect.TypeOf(cellphone))

	var notebook Product
	notebook.Name = "Macbook Pro"
	notebook.Description = "Macbook Pro 15 pol"
	notebook.Value = 10000

	fmt.Println(notebook)
}

// Product representantion
type Product struct {
	Name        string
	Description string
	Value       float32
}
