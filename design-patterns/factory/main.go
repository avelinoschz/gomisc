// This snippet implements a small version of the Factory pattern.
// The factory hides which concrete transport is created for each type.
package main

import (
	"errors"
	"fmt"
)

type Transporter interface {
	Transport() string
}

type Truck struct{}

func (t Truck) Transport() string {
	return "transported by truck"
}

type Ship struct{}

func (s Ship) Transport() string {
	return "transported by ship"
}

type TransportType int

const (
	Land TransportType = iota
	Sea
)

func NewTransport(transportType TransportType) (Transporter, error) {
	switch transportType {
	case Land:
		return Truck{}, nil
	case Sea:
		return Ship{}, nil
	default:
		return nil, errors.New("unsupported transport type")
	}
}

func main() {
	landTransport, err := NewTransport(Land)
	if err != nil {
		panic(err)
	}

	seaTransport, err := NewTransport(Sea)
	if err != nil {
		panic(err)
	}

	fmt.Println(landTransport.Transport())
	fmt.Println(seaTransport.Transport())
}
