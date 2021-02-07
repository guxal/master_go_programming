package main

import (
	"fmt"
	f "fmt"
)

func main() {
	fmt.Println("Scope means name visibility")
	sayHello("Andrei")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water Boiling Point in Degrees Fahrenheit:", tf)
	f.Println("something")
}
