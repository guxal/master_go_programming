package main

import "fmt"

func main() {
	// INT TYPE
	var i1 int8 = 100
	// var i1 int8 = -129 // overflows int8
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	// var i2 uint16 = 65539 // overflows uint16
	fmt.Printf("%T\n", i2)

	// FLOAT TYPES
	var f1, f2, f3 float64 = 1.1, -.2, 5.

	fmt.Printf("%T %T %T\n", f1, f2, f3)
	
	// RUNE TYPE
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r) // HEXADECIMAL

	// BOOL TYPE
	var b bool = true
	fmt.Printf("%T\n", b)
	
	// STRING TYPE

	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)
}