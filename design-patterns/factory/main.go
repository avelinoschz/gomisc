package main

import (
	"errors"
	"fmt"
)

type Transporter interface {
	Transport() string
}

type Truck struct{}

func (t *Truck) Transport() string {
	return "transported by truck"
}

type Ship struct{}

func (s *Ship) Transport() string {
	return "transported by ship"
}

type TransportType int32

const (
	Land TransportType = iota
	Sea
)

type TransportFactory struct{}

func (tf *TransportFactory) NewTransport(transportType TransportType) (Transporter, error) {
	switch transportType {
	case Land:
		return &Truck{}, nil
	case Sea:
		return &Ship{}, nil
	default:
		return nil, errors.New("unsupported transport type")
	}
}

func main() {
	fmt.Println("land:", Land)
	fmt.Println("sea:", Sea)

	factory := &TransportFactory{}
	t1, err := factory.NewTransport(Land)
	if err != nil {
		panic(err)
	}
	fmt.Println(t1.Transport())

	t2, err := factory.NewTransport(Sea)
	if err != nil {
		panic(err)
	}
	fmt.Println(t2.Transport())
}
