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
	// alex := person{"Alex", "Anderson"}
	// alex2 := person{firstName: "Alex", lastName: "Anderson"}
	// var alex3 person

	// alex3.firstName = "Alex"
	// alex3.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Println(alex2)
	// fmt.Println(alex3)
	// fmt.Printf("%+v", alex3)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // note the commma
		},
	}

	fmt.Printf("%+v", jim)

	// &variable: Give me the memory address of the value this variable is pointing at
	// jimPointer := &jim
	// jimPointer.updateName("jimmy") // change the firstname of Jim
	// jim.print()

	jim.updateName("jimmy") // jim is person, = "pointer to that person"
	jim.print()
}

// * before a type is different from the * before a pointer
func (pointerToPerson *person) updateName(newFirstName string) { // *person: a type of pointer that points to the person
	// *pointer: Give me the value this memory adress is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
