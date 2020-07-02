package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("App started", time.Now().Format("02-01-2006"))
	fmt.Println("------------")

	var total float32

	for i := 0; i < 10; i++ {
		total = addFunc(10, total)
	}

	total = addFuncNamed(10, total)

	fmt.Println("O total é: ", total)
}

func addFunc(valor, valorAnterior float32) float32 {
	return valorAnterior + valor
}

func addFuncNamed(valor, valorAnterior float32) (total float32) {
	total = valorAnterior + valor
	return
}
