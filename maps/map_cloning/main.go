package main

import "fmt"

func main() {
	//** CLONING A MAP **//
 
    // When creating a map variable Go creates a pointer to a map header value in memory.
    // The key: value pairs of the map are not stored directly into the map.
    // They are stored in memory at the address referenced by the map header.
 
	friends := map[string]int{"Dan": 40, "Maria": 25}

	// neighbors is not a copy of friends.
    // both maps reference the same data structure in memory
	neighbors := friends

	// modifying friends AND neighbors
	friends["Dan"] = 50
	fmt.Println(neighbors) // -> map[Dan:50 Maria:25]

	// How to clone a map?
    // 1. initialize a new map
	people := make(map[string]int) // clone
	// colleagues = friends // -> ERROR, illegal with maps!

	// 2. use a for loop to copy each element into the new map
	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18
	fmt.Printf("%#v --------------- %#v\n", people, friends)

	delete(friends, "Dan")
	fmt.Println(friends)
	// colleagues and friends are different maps in memory!
}