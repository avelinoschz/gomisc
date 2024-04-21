package main

import "fmt"

type Transporter struct {
	name string
}

func (t *Transporter) Transport() string {
	return "generic transporter"
}

func (t *Transporter) Move() bool {
	return true
}

// here we are using the composition pattern
type Truck struct {
	Transporter
}

// with the composition pattern, we can in a way
// apply overriding the original functionality
// from the base struct `Transporter`, extending functionality
// without modifying the original
func (t *Truck) Transport() string {
	return "transported by truck"
}

type Ship struct {
	Transporter
}

func (s *Ship) Transport() string {
	return "transported by ship"
}

type Plane struct {
	Transporter
}

func main() {
	truck := &Truck{}
	ship := &Ship{}
	plane := &Plane{
		Transporter: Transporter{
			name: "just a transport",
		},
	}

	fmt.Println(truck.Transport())
	fmt.Println(ship.Transport())
	fmt.Println(plane.Transport())

	fmt.Println(plane.name)
	fmt.Println(plane.Move())
	manageTransport(&plane.Transporter)
	// manageTransport(&plane)
	// this does not work because it has embedded a Transporter
	// it is not considered a Transporter like in inheritance
}

func manageTransport(t *Transporter) {
	fmt.Println(t.Transport())
}
