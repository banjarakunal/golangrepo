package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello struct represents hello handler
type Hello struct {
	l *log.Logger
}

//NewHello creates Hello struct with log.logger injected
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//ServeHTTP retruns hello handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	log.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Hello %s \n", d)
}
