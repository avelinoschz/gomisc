// This snippet shows that different types can be used through the same interface.
// The calling code depends on the interface, not on a concrete notifier.
package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending email:", message)
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

func sendNotification(notifier Notifier, message string) {
	notifier.Send(message)
}

func main() {
	sendNotification(EmailNotifier{}, "Welcome to our service!")
	sendNotification(SMSNotifier{}, "Your verification code is 123456")
}
