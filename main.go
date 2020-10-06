package main

import (
	"log"
	"firstMicroService/handlers"
	"net/http"
	"os"
)

func main() {

	// we create a reference to our handler that has been abstracted to another function in another file
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l) // needs a logger which will now be added up - we then inject the logger here
	gh := handlers.NewGoodbye(l)
	// create new serveMux
	sm := http.NewServeMux() // then on the serveMux we're registering a handler. so we pass in a path
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	//  we then register the hellohandler "hh" with our server
	http.ListenAndServe(":9000", sm) // the bind address - binds to all ip addresses on the laptop. second option is a handler. if set to "nil", it will use the default handler which is the http server mux
}