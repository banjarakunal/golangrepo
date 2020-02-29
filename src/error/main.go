package main

import "fmt"

type customError struct {
	info string
}

func main() {
	c1 := customError{
		info: "Need tea",
	}
	foo(c1)

}

func foo(e error) {
	fmt.Println(e)
}

func (c customError) Error() string {
	return fmt.Sprintf("Error string : %v ", c.info)
}
