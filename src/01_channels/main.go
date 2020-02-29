package main

import "fmt"

func main() {
	c := make(chan int)
	//close(c)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
