package main

import (
	"fmt"
	"pointerstest"
)

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


func main() {
	p1 := pointerstest.MakePerson("John", "Doe", 30)
	p2 := pointerstest.MakePersonPointer("Jane", "Smith", 25)

	fmt.Printf("Stack allocation: %+v\n", p1)
	fmt.Printf("Heap allocation: %+v\n", p2)
}