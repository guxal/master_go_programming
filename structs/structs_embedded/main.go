package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone int
	}

	type Employee struct {
		name string
		salary int
		contactInfo Contact
	}

	john := Employee {
		name: "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email: "jkeller@company.com",
			address: "Stree 20, London",
			phone: 003323221121212,
		},
	}
	fmt.Printf("%+v\n", john)
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email)
	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's new email:%s\n", john.contactInfo.email)
}