package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	fmt.Println(fileName)
	readFile(fileName)
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)
	defer file.Close()

	/*for {
		b := make([]byte, 5)
		n, err := file.Read(b)

		if n > 0 {
			fmt.Println(string(b[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error :", err)
			break
		}

	}*/

}
