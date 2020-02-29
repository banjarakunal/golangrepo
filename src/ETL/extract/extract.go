//Package extract provide capablity of retrieving contents from  given input
package extract

import (
	"io"
	"log"
	"os"
)

//Extracter interface is used for polymorphic purpose
//we can add other reader or extracter later
type Extracter interface {
	read() (string, error)
}

//FileReader struct reads file content of Fname
type FileReader struct {
	Fname string
}

//Reader strcuts reads contents from  file
type Reader struct{}

//read extract the content from given input
func (r FileReader) read() (content string, err error) {
	file, err := os.Open(r.Fname)
	if err != nil {
		return "", err
	}
	b := make([]byte, 2)

	var s string
	for {
		n, err := file.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
		if n > 0 {
			//fmt.Println("abc")
			s += string(b[:n])
		}
	}

	if err != nil {
		return "", err
	}
	//fmt.Println("File read...", s)
	return s, nil
}

//Extract extracts the content from given input
func Extract(e Extracter) (op string, err error) {
	return e.read()
}
