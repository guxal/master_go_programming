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

func (c circle) volume() float64 {
	return 4/3 * math.Pi * math.Pow(c.radius, 3)
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
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s) //interface dynamic type is circle
 
    // no direct access to interface's dynamic values
    // s.volume() -> error
 
    // there is access only to the methods that are defined inside the interface
    fmt.Printf("Circle Area:%v\n", s.area())
 
    // an interface value hides its dynamic value.
    // use type assertion to extract and return the dynamic value of the interface value.
    fmt.Printf("Sphere Volume:%v\n", s.(circle).volume())

	// s.volume()
	// checking if the assertion succeded or not
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume())
	}

	//** TYPE SWITCHES **//
 
    // it permits several type assertions in series
	s = rectangle { width: 3.4, height: 2.2 }
	switch value := s.(type) {
		case circle:
			fmt.Printf("%#v has has circle type\n", value)
		case rectangle:
			fmt.Printf("%#v has has rectangle type\n", value)		
	}
}
