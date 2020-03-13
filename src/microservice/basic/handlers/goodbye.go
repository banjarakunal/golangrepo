package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

//NewGoodBye creates goodbye struct with logger injected
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (gb *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}
