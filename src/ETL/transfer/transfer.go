//Package transfer converts input content into required format
package transfer

import "strings"

//Transfer struct is for transfering extract content one form to another
type Transfer struct{}

//ConvertUpper string  to upper case
func ConvertUpper(content string) string {
	return strings.ToUpper(content)
}

//WordCount counts uniques words from given string
func WordCount(content string) map[string]int {
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
	return m
}
