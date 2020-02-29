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
	contact := contactInfo{"abc@mail.com", 411041}
	alex := person{"kunal", "Banjara", contact}

	alex.updateName("kunnu")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
