package main

// There are 3 Scopes:
//  - File Scope
//  - Package Scope
//  - Block (local)  Scope

import "fmt"

// import "fmt" //error

// importing as another name (alias) is permitted
import f "fmt" //permitted in go

const done = false //package scoped

func main() {
	var task = "Running" //local (block) scoped
	fmt.Println(task, done)

	const done = true //local scoped
	f.Printf("done in main() is %v\n", done)

	f.Println("Bye Bye!")
}

func f1() {
	fmt.Printf("done in f1(): %v\n", done) //this is done from package scope
}