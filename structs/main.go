package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // note the commma
		},
	}

	fmt.Printf("%+v", jim)
}
