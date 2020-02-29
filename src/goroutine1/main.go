package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("os", runtime.GOOS)
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("CPUS", runtime.NumCPU())
	fmt.Println("Go routine", runtime.NumGoroutine())
}
