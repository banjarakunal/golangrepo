package main

import (
	"fmt"
)

func main() {
	o := make(chan int)
	e := make(chan int)
	q := make(chan int)
	//go recieve(o, e, q)
	go send(o, e, q)
	//runtime.Gosched()
	recieve(o, e, q)
	fmt.Println("About to exit")
	close(q)

}

func send(o, e, q chan<- int) {
	fmt.Println("In send")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("In even", i)
			o <- i
		} else {
			e <- i
		}

	}
	//close(o)
	//close(e)
	q <- 0
}

func recieve(o, e, q <-chan int) {
	//runtime.Gosched()
	fmt.Println("In recieve", runtime.NumGoroutine())
	for {
		select {
		case v := <-e:
			fmt.Println("Even :", v)
		case v := <-o:
			fmt.Println("Odd :", v)
		case v := <-q:
			fmt.Println("Odd :", v)
			return

		}
	}
}
