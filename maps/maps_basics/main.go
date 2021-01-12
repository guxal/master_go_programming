package main

import "fmt"

func main() {
	// declaring a map with keys of type string and values of type stri
	var employees map[string]string
	//the zero value of a map is nil

	// the zero value of a map is nil
	fmt.Printf("%#v\n", employees) // -> map[string]string(nil).

	fmt.Printf("No of pairs: %d\n", len(employees)) // => No. of elements: 0

	// getting the corresponding value of a key
    // if the key doesn't exist or the map is not initialized it returns the zero value for the value type
	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	// key must be of comparable types
    // var m1 map[[]int]string // error -> invalid map key type []int (slices are not comparable)
	var m1 map[[5]int]string
	_ = m1

	// inserting a key:value pair in a nil map
    // employees["Dan"] = "Programmer" // error -> panic: assignment to entry in nil map

	// declaring and initializing a map using a map literal
	people := map[string]float64{} // empty map

	// inserting key:value pairs in a map
	people["John"] = 21.4
	people["Marry"] = 10
	fmt.Println(people) // => map[John:30.5 Marry:22]

	// declaring and initializing a map using the make() function:
	map1 := make(map[int]int)
	map1[4] = 8
	fmt.Printf("%#v\n",map1) // -> map1: map[int]int{}

	// declaring and initializing a map using a map literal
	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.13,
		//50: "ABC"  //illegal, all keys have the same type which is string and all values have the same type which is float64
		"CHF": 3243.1, //this last comma (,) is mandatory when declaring the map on multiple lines
	}
	fmt.Printf("%#v\n",balances) // => map[CHF:600 EUR:555.11 USD:233.11]

	//If we declare and initialize a map on a single line, it's not mandatory to use a comma after the last element
	m := map[int]int{ 1: 10, 2: 20, 3: 30}
	_ = m

	// initializing a map with duplicate keys
    // n := map[int]int{1: 3, 4: 5, 6: 8, 1: 4} // error -> duplicate key 1 in map literal

	// if the key exists it updates its value and if the key doesn't exist it adds the key: value pair
	balances["USD"] = 500.1
	balances["GBP"] = 100
	fmt.Println(balances) // => map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]

	fmt.Println(balances["RON"])
	// "comma ok" idiom is used to distinguish between a missing key:value pair and an existing key with value zero
	v, ok := balances["RON"]

	
    //v is the key's corresponding value
    // ok is bool type value which is true if the key exists and false otherwise
 
	if ok {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON Key doesn't exist in the map!")
	}

	// iterating over a map
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	// deleting a key:value pair from the map
	delete(balances, "USD")
	fmt.Println(balances)
}