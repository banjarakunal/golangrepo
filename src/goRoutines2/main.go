package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	for i := 0; i < 2; i++ {

		go func() {
			fmt.Println("Goroutine: ", runtime.NumGoroutine, " no : ", i)
		}()
		wg.Done()
	}
	wg.Wait()

	fmt.Println("DOne")
}
