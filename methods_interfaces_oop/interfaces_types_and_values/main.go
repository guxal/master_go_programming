package main

import (
	"fmt"
	"math"
)

// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface {
	area() float64
	perimeter() float64
}
// declaring 2 struct types that represent geometrical shapes: rectangle and circle
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}
// method that calculates circle's area
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
// method that calculates circle's perimeter
func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}
// method that calculates rectangle's area
func (r rectangle) area() float64{
	return r.height * r.width
}
// method that calculates rectangle's perimeter
func (r rectangle) perimeter() float64{
	return 2 * (r.width + r.height)
}

 
// any type that implements the interface is also of type of the interface
// rectangle and circle values are also of type shape
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape
	fmt.Printf("%T\n", s)

	ball := circle {radius: 2.5}
	s = ball

	print(s)
	fmt.Printf("Type of s: %T\n", s)

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s)

}
