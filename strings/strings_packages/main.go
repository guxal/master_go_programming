package main

import (
	"fmt"
	"strings"
)

func main() {
	// declaring a variable of type func to call the Println function easier.
	p := fmt.Println
	// it returns true whether a substr is within a string
	result := strings.Contains("I Love Go Programming!", "Love")
	p(result)
	// it returns true whether any Unicode code points are within our string, and false otherwise.
	result = strings.ContainsAny("success", "xy") // "xys" -> true
	p(result)
	// it reports whether a rune is within a string.
	result = strings.ContainsRune("golang", 'g')
	p(result)
	// it returns the number of instances of a substring in a string
	n := strings.Count("cheese", "e")
	p(n)
	// if the substr is an empty string Count() returns 1 + the number
	n = strings.Count("Five", "")
	p(n)
	// ToUpper() and ToLower() return a new string with all the letters
    // of the original string converted to uppercase or lowercase.
	p(strings.ToLower("GO PyTHON jaVA"))
	p(strings.ToUpper("GO PyTHON jaVA"))

	// comparing strings (case matters)
	p("go" == "go")
	p("GO" == "go")

	lower := strings.ToLower
	// comparing strings (case doesn't matter) -> it is not efficient
	p(lower("GO") == lower("go"))
	// EqualFold() compares strings (case doesn't matter) -> it's efficient
	p(strings.EqualFold("GO", "go"))
	// it returns a new string consisting of n copies of the string that is passed as the first argument
	myStr := strings.Repeat("ab", 10)
	p(myStr)
	// it returns a copy of a string by replacing a substring (old) by another substring (new)
	myStr = strings.Replace("192.168.0.1", ".", ":", 2)  // it replaces the first 2 occurrences
	p(myStr)
	// if the last argument is -1 it replaces all occurrences of old by new
	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // replace all
	p(myStr)
	// ReplaceAll() returns a copy of the string s with all non-overlapping instances of old replaced by new.
	myStr = strings.ReplaceAll("192.168.0.1", ".", ":") // replace all
	p(myStr)
	// it slices a string into all substrings separated by separator and returns a slice of the substrings between those separators.
	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)
	// If separator is empty Split function splits after each UTF-8 rune literal.
	s = strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", s)
	// Join() concatenates the elements of a slice of strings to create a single string.
    // The separator string is placed between elements in the resulting string.
	s = []string{"I", "Learn", "Golang"}
	myStr = strings.Join(s, "-")
	p(myStr)
	// splitting a string by whitespaces and newlines.
	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)
    // TrimSpace() removes leading and trailing whitespaces and tabs.
	s1 := strings.TrimSpace("\t Goodby Windows, Welcome Linux! \n")
	fmt.Printf("%q\n", s1)
    // To remove other leading and trailing characters, use Trim()
	s2 := strings.Trim("...Hello Gophers!!!?", ".!?")
	fmt.Printf("%q\n", s2)
}
