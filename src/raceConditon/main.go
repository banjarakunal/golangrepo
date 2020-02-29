package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	var counter int64
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter incrmentor", atomic.LoadInt64(&counter))
			fmt.Println("Goroutines: ", runtime.NumGoroutine())
		}()
		wg.Done()
	}

	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

}
