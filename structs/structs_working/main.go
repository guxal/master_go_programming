package main

import "fmt"

func main() {
	// creating a struct type
	type book struct {
		title  string
		author string
		year   int
	}
	//if we create a new struct value by omitting some fields they will be zero-valued according to their type
	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	// page := lastBook.page // error undefine (type book has no field or method pages)
    // retrieving the value of a struct field
	fmt.Printf("%#v\n", lastBook)
	// updating a field
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)
	// + modifier with %v  printed out both the field names and their values
		
	// comparing struct values
    // two struct values are equal if their corresponding fields are equal.
	aBook := book{title: "Anna Karenina", author: "", year: 0}
	fmt.Println(aBook == lastBook)
	// creates a copy of a struct
	myBook := aBook
	myBook.year = 2020   // modifying only myBook
	Println(myBook, aBook)
}