
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}
type Person struct {
	FirstName    string
	LastName     string
	Age          int
	Address      string
	PhoneNumbers []string
	Emails       []string
	Notes        map[string]string
	Friends      []*Person
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
