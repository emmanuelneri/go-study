package main

import (
	"errors"
	"fmt"
)

func main() {
	err := validate("")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR - %s", err.Error()))
	}

	printMessage(&message{text: "Text"})
	printMessage(nil)
}

func validate(value string) error {
	if len(value) == 0 {
		return errors.New("value can not be empty")
	}

	return nil
}

func printMessage(m *message) {
	defer panicHandler()
	fmt.Printf("printing %s", m.text)
}

func panicHandler() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
}

type message struct {
	text string
}
