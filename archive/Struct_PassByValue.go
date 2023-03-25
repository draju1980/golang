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
	tim := person{
		firstName: "Tim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "hello@world.com",
			zipCode: 94600,
		},
	}
	tim.updateName("Jimmy")
	tim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
	fmt.Println("\n", p)
}
