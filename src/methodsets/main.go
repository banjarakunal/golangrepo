package main

import "fmt"

type person struct{}

func (p *person) speak() {
	fmt.Println("helloo")
}

func saySomething(h human) {
	h.speak()
}

type human interface {
	speak()
}

func main() {

	p := person{}
	saySomething(&p)
}
