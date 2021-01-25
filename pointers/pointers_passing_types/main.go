package main

import "fmt"

// declaring a function that takes an int, a float, a string and a bool value.
// the function works on copy so the changes are not seen outside (pass by value)
func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

// declaring a function that takes in a pointer to int, a pointer to float, a pointer to string and a pointer to bool.
// the function makes a copy of each pointer but they point to the same address as the originals
func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	//changing the values the pointers point to is seen outside the function
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

// declaring struct type
type Product struct {
	price float64
	productName string
}

// declaring a function that takes in a struct value and modifies it
func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
	// the changes are not seen to the outside world
}

// declaring a function that takes in a pointer to struct value and modifies the value
func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivalent to p.price = 300
	p.productName = "Bicycle"
	// the changes are seen to the outside world
}

// declaring a function that takes in a slice
func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
	// the changes are seen to the outside world
}

// declaring a function that takes in a map
func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["a"] = 10
	// the changes are seen to the outside world
}

func main() {
	// declaring some variables
	quantity, price, name, sold := 5, 300.4, "Laptop", true

	fmt.Println("BEFORE calling changevalues():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("AFTER calling changevalues():", quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("AFTER calling changevaluesByPointer():", quantity, price, name, sold)

	gift := Product {
		price: 100,
		productName: "Watch",
	}

	changeProduct(gift)
	fmt.Println(gift)

	fmt.Println("BEFORE calling changeProductByPointer():", gift)
	changeProductByPointer(&gift)
	fmt.Println("AFTER calling changeProductByPointer():", gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println("prices slices after calling changeSlice():", prices)

	myMap := map[string]int{"a":1, "b":2}
	changeMap(myMap)

	fmt.Println("myMap after calling changeMap():", myMap)

}
