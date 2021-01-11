package main

import "fmt"

func main() {
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}

	// fmt.Println(a == b) // invalid operation a == b (map can only be compared to nil)

	// string representation
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)
	
	fmt.Println(s1, s2) // map[A:X] map[B:Y]

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	c := map[string]string{"A": "X", "B": "Y"}
	d := map[string]string{"B": "Y", "A": "X"}

	s3 := fmt.Sprintf("%s", c)
	s4 := fmt.Sprintf("%s", d)

	fmt.Println(s3, s4) // map[A:X B:Y] map[A:X B:Y]

	if s3 == s4 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
