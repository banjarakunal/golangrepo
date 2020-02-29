package main

import "fmt"

func main() {
	nos := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for no := range nos {
		if no%2 == 0 {
			fmt.Println(no, "is even")
		} else {
			fmt.Println(no, "is odd")
		}
	}
}
