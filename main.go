package main

import (
	"fmt"
	"pointerstest"
)
func main() {
	p1 := pointerstest.MakePerson("John", "Doe", 30)
	p2 := pointerstest.MakePersonPointer("Jane", "Smith", 25)

	fmt.Printf("Stack allocation: %+v\n", p1)
	fmt.Printf("Heap allocation: %+v\n", p2)
}