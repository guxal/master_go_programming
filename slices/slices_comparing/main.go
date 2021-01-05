package main

import "fmt"

func main() {
	// declaring a string slice, by default is initialized with nil
	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	// to compare 2 slices use a for loop to iterate over the slices 
	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	// fmt.Println(a == b) // invalid operation a == b (slice can only be compared to nil)

	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}
	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}