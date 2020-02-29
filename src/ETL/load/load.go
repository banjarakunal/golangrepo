//Package load store/load converted data into output mechanism
package load

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//Load struct is for loading purpose
type Load struct {
	FileName string
}

//LoadContent Write string contents into out put file
func (l Load) LoadContent(s interface{}) {
	if os.MkdirAll("output", 0666) != nil {
		panic("Unable to create directory for tagfile!")
	}

	dst, err := os.OpenFile(strings.Join([]string{"output", l.FileName}, "/"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer dst.Close()

	if err != nil {
		log.Fatal("Error while writing file: ", err)
	}

	switch v := s.(type) {
	case string:
		_, err = dst.Write([]byte(v))
		if err != nil {
			log.Fatal("Error while writing1 file: ", err)
		}
	case map[string]int:
		var sw string
		for word, count := range v {
			sw += fmt.Sprintf("\n%v -> %v \n", word, count)
		}
		_, err = dst.Write([]byte(sw))
		if err != nil {
			log.Fatal("Error while writing2 file: ", err)
		}
	default:
		log.Fatal("Not supporting given type :", v)
	}
}
