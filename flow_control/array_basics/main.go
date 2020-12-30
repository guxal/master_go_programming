package main

import "fmt"

func main() {
	var numbers [4]int
	
	fmt.Printf("%v\n", numbers) // [0 0 0 0]
	fmt.Printf("%#v\n", numbers) // [4]int{0, 0, 0, 0}

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1) // [4]float64{0, 0, 0, 0}

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2) // [3]int{-10, 1, 100}

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3) // [4]string{"Dan", "Diana", "Paul", "John"}

	// only initialize two elements
	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4) // [4]string{"x", "y", "", ""}

	// ellipsis operator
	a5 := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("%#v\n", a5) // [6]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("the length of a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3,
		4,
		5, // this comma is mandatory
	}
	fmt.Println(a6) // [1 2 3 4 5] 
}
