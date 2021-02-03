package main

import "fmt"

// declaring a new struct type
type Car struct {
	brand string
	price int
}

// declaring a method for car type
// it changes the value it works on
func changeCar(c Car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
	// the changes are not seen to the outside world (pass by value)
}

func (c Car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// declaring a method with a pointer receiver
func (c *Car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
	// the changes are seen the outside world
}

// method declarations are not permitted on named types that are themselves pointer types
type distance *int

func (d *distance) m1() {
	fmt.Println("just a message")
}

func main() {
  	// declaring a car value
	myCar := Car{brand: "Audi", price: 40000}
	// calling the method with a value receiver
	changeCar(myCar, "Renault", 20000)

	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar)

	// same myCar.changeCar2
	(&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar)
	// declaring a pointer to car
	var yourCar *Car
	yourCar = &myCar
	fmt.Println(*yourCar)
	// valid ways
    // calling the method on pointer type
    // valid ways to call the method:
	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar)
    // in the above example both myCar and yourCar variables have been modified.
	fmt.Println(myCar)
}
