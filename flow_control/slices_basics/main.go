package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	// cities[0] = "London" // error index out of range

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers) // [2 3 4 5]	
	fmt.Printf("%#v\n", numbers) // []int{2, 3, 4, 5} 


	nums := make([]int, 2) 
	fmt.Printf("%#v\n", nums) // []int{0, 0}

	type names []string
	friends := names{"Dan", "Maria"} 
	fmt.Println(friends) // [Dan Maria]

	myFriend := friends[0]
	fmt.Println("My best friend is ", myFriend)
	
	friends[0] = "Gabriel"
	fmt.Println("My best friend is ", friends[0])

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	var n[]int
	n = numbers 
	fmt.Println(n) // [2 3 4 5]
}