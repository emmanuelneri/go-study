package main

import (
	"fmt"
	"reflect"
)

// Product representation
type Product struct {
	Name        string
	Description string
	Value       float32
}

func main() {

	cellphone := Product{"Galaxy", "Galaxy s10", 3000}
	fmt.Println(cellphone)
	fmt.Println("Cellphone type:", reflect.TypeOf(cellphone))

	var notebook Product
	notebook.Name = "Macbook Pro"
	notebook.Description = "Macbook Pro 15 pol"
	notebook.Value = 10000

	fmt.Println(notebook)

	var productPrice ProductPrice
	productPrice.Name = "Galaxy s20"
	productPrice.Description = "Galaxy s20 Ultra"
	productPrice.CostValue = 4999
	salesValue := productPrice.CalculateSalesPrice(20)

	fmt.Printf("Product %s cost %.2f", productPrice.Name, salesValue)

	order := NewOrder()
	order.addProduct(notebook, 1)
	order.addProduct(cellphone, 2)
	fmt.Print(order)

}
