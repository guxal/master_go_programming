package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string
		author string
		year  int
	}

	type book1 struct {
		title, author string
		year  int
	}

	myBook := book{"The Divine Comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook)

	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945}
	_ = bestBook

	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook)
}