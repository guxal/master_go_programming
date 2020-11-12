package main

import "fmt"

func main() {
	// COMPARISON
	a, b := 5, 10
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	
	fmt.Println(a > b) // false
	fmt.Println(a > 5, a >= 5) // false, true
	fmt.Println(b < a, 10 <= b) // false, true

	// LOGICAL

	//&& logical and
	c, d := 5, 10
	fmt.Println(c > 1 && d == 10)

	fmt.Println(c == 5 || d == 100)

	// NEGATION
	fmt.Println(!(c > 0))

	fmt.Println(!(a == 1) || (b == 100))

}