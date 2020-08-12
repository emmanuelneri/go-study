package main

import "fmt"

// Printer interface to use printMessage method
type Printer interface {
	print()
}

// TextMessage struct to represent a message
type TextMessage struct {
	message string
}

func (textMessage TextMessage) print() {
	fmt.Println("Message: " + textMessage.message)
}

// Email struct to represent a message
type Email struct {
	subject, body string
}

func (email Email) print() {
	fmt.Println("Message: " + email.subject + " - " + email.body)
}

func printMessage(printer Printer) {
	printer.print()
}

func main() {
	textMessage := TextMessage{message: "Text message"}
	printMessage(textMessage)

	email := Email{subject: "Email Subject", body: "Hello"}
	printMessage(email)

}
