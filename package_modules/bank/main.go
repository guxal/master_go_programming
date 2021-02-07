package main

import (
	"fmt"
	"my_packages/numbers"
)

func main() {
	var x uint = 40
	fmt.Printf("%d is even: %t \n", x,numbers.Even(x))
}