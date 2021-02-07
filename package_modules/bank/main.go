package main

import (
	"fmt"
	"my_packages/numbers"
)

func main() {
	var x int = 40
	fmt.Printf("%d is even: %t \n", x,numbers.Even(x))
	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100))
	fmt.Println(numbers.OddAndPrime(25))
	fmt.Println(numbers.OddAndPrime(13))
}