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

	fmt.Println("Total: ", total)
}

func addFunc(newValue, currentValue float32) float32 {
	return newValue + currentValue
}

func addFuncNamed(newValue, currentValue float32) (total float32) {
	total = newValue + currentValue
	return
}
