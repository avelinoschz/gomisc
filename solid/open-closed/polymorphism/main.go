// This snippet shows Open/Closed through interfaces and polymorphism.
// The calculator depends on the Shape interface, so new shapes do not change it.
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

// Without the interface, this function would need a type switch for every new shape.
func areaSum(shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

func main() {
	circle := Circle{Radius: 5}
	square := Square{Length: 7}

	fmt.Println(areaSum(circle, square))
}
