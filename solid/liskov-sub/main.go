package main

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(message string) error {
	fmt.Println("Sending email:", message)
	return nil // simplified for example purposes
}

type SMSNotifier struct{}

func (s SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS:", message)
	return nil
}

type PushNotifier struct{}

func (p PushNotifier) Send(message string) error {
	fmt.Println("Sending push notification:", message)
	return nil
}

type AnotherNotifier struct{}

func (a AnotherNotifier) Send(message string) error {
	fmt.Println("Sending via Another Notifier: ", message)
	return nil
}

// constructor returns an interface, instead of a concrete struct
func NewAnotherNotifier() Notifier {
	return &AnotherNotifier{}
}

func SendNotification(notifier Notifier, message string) {
	err := notifier.Send(message)
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func main() {
	emailNotifier := EmailNotifier{}
	smsNotifier := SMSNotifier{}
	pushNotifier := PushNotifier{}

	// this is a example that we can substitute the interface
	// with concrete structs. Both can be sent as an arg
	anotherNotifier := NewAnotherNotifier()
	SendNotification(anotherNotifier, "Yes I am (an Interface)!")

	SendNotification(emailNotifier, "Welcome to our service!")
	SendNotification(smsNotifier, "Your verification code is 123456")
	SendNotification(pushNotifier, "You have a new message")
}
