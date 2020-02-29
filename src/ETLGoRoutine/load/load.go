package load

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

//Load struct is for loading purpose
type Load struct {
	FileName string
}

//LoadContent Write string contents into out put file
func (l Load) LoadContent(s <-chan string, w <-chan map[string]int) {
	fmt.Println("In LoadContent", runtime.NumGoroutine())
	if os.MkdirAll("output", 0666) != nil {
		panic("Unable to create directory for tagfile!")
	}

	dst, err := os.OpenFile(strings.Join([]string{"output", l.FileName}, "/"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer dst.Close()

	if err != nil {
		log.Fatal("Error while writing file: ", err)
	}
	//go func() {
		for {
			select {
			case upperContent := <-s:
				if upperContent == ""{
					return
				}
				_, err = dst.Write([]byte(upperContent))
				if err != nil {
					log.Fatal("Error while writing1 file: ", err)
				}
			case wordCount := <-w:
				if wordCount == nil{
					return
				}
				var sw string
				for word, count := range wordCount {
					sw += fmt.Sprintf("\n%v -> %v \n", word, count)
				}
				_, err = dst.Write([]byte(sw))
				if err != nil {
					log.Fatal("Error while writing2 file: ", err)
				}
			default:
				fmt.Println("No message recieved")
			}
		}
	//}()
	// dst.Close()
}
