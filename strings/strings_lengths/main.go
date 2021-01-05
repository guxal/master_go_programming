package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Codruţa"

	fmt.Println(len(name))

	fmt.Println(len("ஆ"))

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("ஆ")
	fmt.Println(n, m)
}
