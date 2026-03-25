// This snippet explores Open/Closed through composition instead of inheritance.
// Embedded behavior can be reused, while specific types override only what they need.
package main

import "fmt"

type Transporter struct{}

func (t Transporter) Transport() string {
	return "generic transporter"
}

func (t Transporter) Move() string {
	return "moving cargo"
}

type Truck struct {
	Transporter
}

func (t Truck) Transport() string {
	return "transported by truck"
}

type Plane struct {
	Transporter
}

func main() {
	truck := Truck{}
	plane := Plane{}

	fmt.Println(truck.Transport())
	fmt.Println(truck.Move())
	fmt.Println(plane.Transport())
	fmt.Println(plane.Move())
}
