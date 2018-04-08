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

	jim.updateName("jimmy") // did not change the jim's firstname to "jimmy"
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
