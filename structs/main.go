package main

import "fmt"

type Contact struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   Contact
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		contact: Contact{
			email:   "jimparty@mail.com",
			zipCode: 94000,
		},
	}

	jim.print()
	jim.updateName("Jimmy")
	jim.print()
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) updateName(n string) {
	p.firstName = n
}
