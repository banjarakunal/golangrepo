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

	c, err := extract.Extract(e)
	if err != nil {
		log.Fatal(e)
	}
	uc := transfer.ConvertUpper(c)
	wc := transfer.WordCount(c)
	l := load.Load{
		FileName: f,
	}

	l.LoadContent(uc, wc)

	fmt.Println("Done...")
}
