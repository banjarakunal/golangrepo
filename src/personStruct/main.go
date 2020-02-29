package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	iceCreamFlavours []string
}

func main() {
	tukaram := person{
		firstName:        "Tukaram",
		lastName:         "Borse",
		iceCreamFlavours: []string{"vanila", "choclates"},
	}
	fmt.Printf("%#v", tukaram)
	fmt.Println(tukaram.iceCreamFlavours[:2])
}
