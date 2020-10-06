package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// adding a http handler
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {  //add a path and a function - i.e, when a request comes in and matches this path, handle this function
		log.Println("Hello World")
		

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
	})

	// adding another http handler
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {  //add a path and a function - i.e, when a request comes in and matches this path, handle this function
		log.Println("Goodbye World")
	})

	//handleFunc is a convennience method on the go http package, and it registers a function to a path of what is called the "Default serve mux".
	// The Default serve mux is a http handler and everything related to a server in go in a http handler.


	// the ListenAndSerce function passed below is a convenience method that constructs a http server and registers a default handler to it.
	// the second parameter is a handler which when not registered uses the default serve mux

	http.ListenAndServe(":9000", nil) // the bind address - binds to all ip addresses on the laptop. second option is a handler
}