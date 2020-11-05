package main

import "fmt"
// import "math"

func main() {
	const days int = 7

	var i int 
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	//x, y := 5, 0       this is an error
	//fmt.Println(x / y)

	const b = 5
	const a = 0

	// fmt.Println(b / a)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	const (
		min4 = -500
		min5
		min6
	)

	fmt.Println(min4, min5, min6)

	// Constant rules
	// 1. you can not change a constant
	const temp int = 100
	// temp = 50
	// 2. you can not initiate a constant at runtime
	// const power = math.Pow(2, 3)

	// 3. you can not use a variable to initialize a constant
	//t := 5
	//const tc = t

	// 4.
	const l1 = len("hello")
}