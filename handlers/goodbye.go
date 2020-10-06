package handlers

import (
	"log"
	"net/http"
	// "fmt"
	// "io/ioutil"
)

//Goodbye ...
type Goodbye struct { // this says "type Goodbye", "type struct"
	l *log.Logger
}

//NewGoodbye ...
func NewGoodbye(l *log.Logger) *Goodbye { // dependency injection here
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}