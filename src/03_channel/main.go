package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://godoc.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, site := range sites {
		go statusChecker(site, c)
		//fmt.Println("Status checking")
		//fmt.Println(<-c)
		//fmt.Println(<-c)
	}
	//close(c)
	//fmt.Println(<-c)
	for l := range c {
		fmt.Println("In loop")
		go func(link string) {
			time.Sleep(5 * time.Second)
			statusChecker(link, c)
		}(l)
	}

}

func statusChecker(link string, c chan string) {
	//time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	fmt.Println(runtime.NumGoroutine())
	//time.Sleep(5 * time.Second)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- link
		return
	}

	c <- fmt.Sprintf("%s\t is up", link)
}
