package main

import "fmt"

func main() {
	// int age = 10; // c++ way
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	// fmt.Println("Your name is:", name)
	_ = name // not error program

	s := "learning golang!" // short declaration operator
	fmt.Println(s)

	// name := "Andrei" // this is an error
	name = "Andrei"
	name1 := "John"
	_ = name1

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	// car, cost := "BMW", 60000 // this is an error
	car, year := "BMW", 60000
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j // swapping variables

	// _, _ = i, j
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)
}
