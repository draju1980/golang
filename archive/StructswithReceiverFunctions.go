package main

import "fmt"

// created struct person and included contactInfo struct on it
type  person struct {
    firstName string
    lastName string
    contactInfo  
}

type contactInfo struct {
  email string
  zipCode int
}


func main() {
  tim := person {
    firstName: "Tim",
    lastName: "Party",
    contactInfo: contactInfo{
      email: "hello@world.com",
      zipCode: 94600,
    },
  }
  tim.print()
}

// function to print the struct
func (p person) print(){
  fmt.Printf("%+v\n", p)
  fmt.Println("\n",p)
}