package main

import "fmt"

func reverse(s []string)[]string{
	startP := 0
	endP := len(s)-1
	midP := len(s)/2

	for {
		if startP == midP {
			fmt.Println(s)
			return s
		}
		s[startP], s[endP] = s[endP], s[startP]
		startP++
		endP--
	}
}