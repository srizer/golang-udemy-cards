package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// *person: type description, means we're working with a pointer to a person
// *pointerToPerson: operator, means we want to manipulate the value the pointer references
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// *pointer: give me the value this
// memory address is pointing at
// &variable: give me the memory address of
// the value this variable is pointing at

// [0001: person{firstName:"Jim"...}]
// [address, value]
// Turn address into value with *address
// Turn value into address with &value

func (p person) print() {
	fmt.Printf("%+v", p)
}
