package main

import (
	"fmt"
	"math"
)

func main(){
	var x uint8 = 255
	x++ //overflow x is 0
	fmt.Println(x)

	// a := int8(255 + 1) this is error

	var b int8 = 127

	fmt.Printf("%d\n", b+1)

	b = -128
	b--
	fmt.Println(b)
	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f = f * 1.2
	fmt.Println(f) // +Inf
}