package main

import (
	"etl.com/etl/extract"
	"etl.com/etl/load"
	"etl.com/etl/transfer"
	"fmt"
	"log"
	"os"
)

func main() {
	f := os.Args[1]

	e := extract.FileReader{
		Fname: f,
	}
	fmt.Println("FileReader : ", e)
	c, err := extract.Extract(e)
	if err != nil {
		log.Fatal(e)
	}
	fmt.Println("file content :", c)

	upperContent := transfer.ConvertUpper(c)
	count := transfer.WordCount(c)

	l := load.Load{
		FileName: f,
	}
	l.LoadContent(upperContent)
	l.LoadContent(count)

	fmt.Println("Transformed content :", transfer.WordCount(c))
}
