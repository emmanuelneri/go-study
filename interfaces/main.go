package main

import "fmt"

// Printer interface to use print method
type Printer interface {
	print()
}

// TextMessage struct represent a message
type TextMessage struct {
	message string
}

func (textMessage TextMessage) print() {
	fmt.Println("Message: " + textMessage.message)
}

// Email rstruct epresent a message
type Email struct {
	subject, body string
}

func (email Email) print() {
	fmt.Println("Message: " + email.subject + " - " + email.body)
}

func print(printer Printer) {
	printer.print()
}

func main() {
	textMessage := TextMessage{message: "Text message"}
	print(textMessage)

	email := Email{subject: "Email Subject", body: "Hello"}
	print(email)

}
