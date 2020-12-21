package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	total := getStartValue()

	for i := 0; i < 10; i++ {
		total = addFunc(10, total)
	}

	total = addFuncNamed(10, total)

	fmt.Println("Total: ", total)

	code, name := getProduct()
	fmt.Println(code)
	fmt.Println(name)

	_, onlyName := getProduct()
	fmt.Println(onlyName)

	value, err := addValue(10, 20)
	if err != nil {
		log.Printf("ERROR - %s", err.Error())
	}

	fmt.Println(value)

	anonymousFunctions()
}

func anonymousFunctions() {
	defer func() {
		fmt.Println("defer anonymousFunctions")
	}()

	func() {
		fmt.Println("anonymousFunctions")
	}()
}

func init() {
	fmt.Println("App started", time.Now().Format("02-01-2006"))
	fmt.Println("------------")
}

func addFunc(newValue, currentValue float32) float32 {
	return newValue + currentValue
}

func addFuncNamed(newValue, currentValue float32) (total float32) {
	total = newValue + currentValue
	return
}

func getStartValue() float32 {
	return 0
}

func getProduct() (string, string) {
	return "123", "Product"
}

func addValue(newValue, currentValue float32) (float32, error) {
	if newValue == 0 {
		return 0, errors.New("invalid new value")
	}
	return newValue + currentValue, nil
}
