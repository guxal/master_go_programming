package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3 // int type
	var y = 3.1 // float64 type

	// x = x * y // this is an error 
	x = x * int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)

	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5 // int type
	fmt.Printf("%T\n", a)
	var b int64 = 2
	a = int(b)
	// preventing unused variable error
	_ = a

	//** CONVERTING NUMBERS TO STRINGS AND STRINGS TO NUMBERS **//

	s := string(99)            // int to rune (Unicode code point)
	fmt.Println(s)             // => 99, the ascii code for symbol c
	fmt.Println(string(34234)) // => 34234 is the unicode code point for 薺

	// we cannot convert a float to a string similar to an int to a string
	// s1 := string(65.1) // error

	// converting float to string
	var myStr = fmt.Sprintf("%f", 5.12)
	fmt.Println(myStr) // => 5.120000

	// converting int to string
	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1) // => 34234

	// converting string to float
	var result, err = strconv.ParseFloat("3.142", 64)
	if err == nil {
	 fmt.Printf("Type: %T, Value: %v\n", result, result) // => Type: float64, Value: 3.142
	} else {
	 fmt.Println("Cannot convert to float64!")
	}

	// Atoi(string to int) and Itoa(int to string).
	i, err := strconv.Atoi("-50")
	s = strconv.Itoa(20)
	fmt.Printf("i Type is %T, i value is %v\n", i, i) // => i Type is int, i value is -50
	fmt.Printf("s Type is %T, s value is %q\n", s, s) // => s Type is string, s value is "20"
}