package main

import "fmt"

const boilingPoint = 100

func sayHello(name string){
	fmt.Println("Hello %s!\n", name)
}

func toFahrenheit(t float64) float64 {
	return t*1.8 + 32
}
