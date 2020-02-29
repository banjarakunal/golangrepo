package transfer

import (
	"strings"
)

//Transfer struct is for transfering extract content one form to another
type Transfer struct{}

//ConvertUpper string  to upper case
func ConvertUpper(content string) <-chan string {
	uc := make(chan string)
	go func() {
		uc <- strings.ToUpper(content)
		close(uc)
	}()
	return uc
}

//WordCount counts uniques words from given string
func WordCount(content string) <-chan map[string]int {
	wc := make(chan map[string]int)
	go func() {
		s := strings.Fields(content)
		m := make(map[string]int)
		for _, word := range s {
			_, ok := m[word]
			if ok {
				m[word]++
			} else {
				m[word] = 1
			}
		}

		wc <- m
		close(wc)
	}()
	return wc
}
