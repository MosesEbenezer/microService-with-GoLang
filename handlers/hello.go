package handlers

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)

// to create a http handler, we first create a struct which implements go http handler

//Hello ...
type Hello struct { // this says "type Hello", "type struct"
	l *log.Logger
}

//NewHello ...
func NewHello(l *log.Logger) *Hello { // dependency injection here
	return &Hello{l}
}

// adding a http handler
// we then add the method which satisfies the http handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		
	h.l.Println("Hello World") // using the dependency injection here as "l" is log.Logger and log.Logger has Println method
	// read from the body - like req body
	d, err := ioutil.ReadAll(r.Body) // body implements the interface ioreadcloser - this reads everything from the body and reads it into the variable d

	// handle err
	if err != nil {
		// rw.WriteHeader(http.StatusBadRequest) // rw = response writer. write header specifies the status response we're going to send back to the user
		// rw.Write([]byte("Ooops an error occured"))
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
		// as an alternative to the above, we can use the Go's httpError convenience method to handle errors
	}

	// log.Printf("Data %s\n", d)

	//write bcak to the user
	fmt.Fprintf(rw, "Hello there %s", d)
}