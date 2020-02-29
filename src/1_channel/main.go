package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println("man done")
	func() {
		fmt.Println("buff val", <-c)
	}()

	fmt.Println("man done")

}
