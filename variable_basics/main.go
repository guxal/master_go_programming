package main

import "fmt"

func main(){
	// int age = 10; // c++ way
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	// fmt.Println("Your name is:", name)
	_ = name // not error program

	s := "learning golang!" // short declaration operator
	fmt.Println(s)

	// name := "Andrei" // is an error
	name = "Andrei"
	name1 := "John"
	_ = name1
}