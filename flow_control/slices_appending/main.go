package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	numbers = append(numbers, 10)
	fmt.Println(numbers) // [2 3 10]

	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) // [2 3 10 20 30 40]

	n := []int{100, 200}
	numbers = append(numbers, n...)
	fmt.Println(numbers) // [2 3 10 20 30 40 100 200]

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)

	fmt.Println(src, dst, nn) // [10 20 30] [10 20 30] 3

	dst2 := make([]int, 2)
	nn2 := copy(dst2, src)

	fmt.Println(src, dst2, nn2) // [10 20 30] [10 20] 2
}
