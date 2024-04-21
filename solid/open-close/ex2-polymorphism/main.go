package main

import (
	"fmt"
	"math"
)

// this example is also the implementation of another SOLID principle
// the dependency inversion principle. rely on abstractions rather than impls

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

type calculator struct{}

// // this is the problem to solve. in this case we would need to modify this
// // to extend the functionality of areaSum() to support other shapes
// func (c *calculator) areaSum(shapes ...interface{}) float64 {
// 	var sum float64
// 	for _, shape := range shapes {
// 		switch shape.(type) {
// 		case circle:
// 			r := shape.(circle).radius
// 			sum += math.Pi * r * r
// 		case square:
// 			l := shape.(square).length
// 			sum += l * l
// 		}
// 		fmt.Println(sum)
// 	}

// 	return sum
// }

// relying on the usage of interfaces, we can extend the functionality
// of the concrete struct without changing this code, and support
// other shapes
func (c *calculator) areaSum(shapes ...shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.area()
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	calc := calculator{}
	sum := calc.areaSum(c, s)
	fmt.Println(sum)
}
