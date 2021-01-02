package main

import "fmt"

func main() {
	a := [5]int {1, 2, 3, 4, 5}
	// a[start:stop]
	b := a[1:3]
	fmt.Printf("%v, %T", b, b) // [2 3], []int

	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]
	fmt.Println(s2) // [2 3]

	//for convenience, any of the indexis may be omitted.
    // a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
    fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
    fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
    fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
    fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
    // fmt.Println(s1[:100])   //panic: runtime error: slice bounds out of range
 
    s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
    fmt.Println(s1)          // -> [1 2 3 4 100]
 
    s1 = append(s1[:4], 200) // overwrites the last element
    fmt.Println(s1)          // -> [1 2 3 4 200]

}