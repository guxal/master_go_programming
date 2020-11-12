package main

import "fmt"

func main() {
	a, b := 4, 2

	r := (a + b) / (a - b) * 2

	fmt.Println(r) // -> 6

	r = 9 % a

	fmt.Println(r)

	// ASSIGNMENT OPERATORS
	c, d := 2, 3

	// INCREMENT
	c += d
	fmt.Println(c)

	// DECREMENT
	d -= 2
	fmt.Println(d)

	// MULTIPLICATION
	d *= 10
	fmt.Println(d)

	// DIVISION
	d /= 5
	fmt.Println(d)
	
	// MODULUS
	c %= 3
	fmt.Println(c)
	
	// INCREMENT
	x := 1
	x += 1
	x++
	fmt.Println(x)
	// DECREMENT
	x--
	fmt.Println(x)

	//fmt.Println(5 + x--) // Error

}