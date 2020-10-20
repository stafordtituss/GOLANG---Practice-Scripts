package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Type-1
	// john := person{"John", "Doe"}
	// Type-2
	// john := person{firstName: "John", lastName: "Doe"}
	// fmt.Println(john)
	// Type-3
	// var john person
	// john.firstName = "John"
	// john.lastName = "Doe"
	// fmt.Printf("%+v", john)

	john := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "john.doe@yolo.com",
			zipCode: 94000,
		},
	}

	// johnPointer := &john
	// johnPointer.updateName("Jane")

	john.updateName("Jane")
	john.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
