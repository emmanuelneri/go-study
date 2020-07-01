package main

import (
	"fmt"
	"time"

	"./structs"
)

// Item struct
type Item structs.Item

func main() {
	const helloWorld = "Hello World"

	fmt.Println(helloWorld, time.Now().Format("02-01-2006"))
	fmt.Println("------------")

	var total float32

	for i := 0; i < 10; i++ {
		item := Item{"Product", 10}
		total = add(item.Value, total)
	}

	fmt.Println("O total é: ", total)
}

func add(valor, valorAnterior float32) float32 {
	return valorAnterior + valor
}

func addNamed(valor, valorAnterior float32) (total float32) {
	total = valorAnterior + valor
	return
}
