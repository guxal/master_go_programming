package main

import "fmt"

func main() {
	// Slicing a string is efficient because it reuses the same backing array
	// Slicing returns bytes not runes
	
	s1 := "I love Golang!"
	fmt.Println(s1[2:5])


	s2 := "中文维基是世界上"

	fmt.Println(s2[0:2]) // infect // -> � - the unicode representation of bytes from index 0 and 1.

	// returning a slice of runes
    // 1st step: converting string to rune slice
	rs := []rune(s2)
	fmt.Printf("%T\n", rs)
	// 2st step: slicing the rune slice
	fmt.Println(string(rs[0:3])) // => 中文维
}
