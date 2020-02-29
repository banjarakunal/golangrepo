package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type hindiBot struct{}

func main() {
	eb := englishBot{}
	hb := hindiBot{}

	printGreeting(eb)
	printGreeting(hb)

}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (hb hindiBot) getGreeting() string {
	return "Namaste"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
