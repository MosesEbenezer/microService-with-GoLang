package main

import (
	"os/signal"
	"context"
	"time"
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


	// manually creating a http server with some configurations - to better handle server downtimes
	s := &http.Server{
		Addr: ":9000",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	//  we then register the hellohandler "hh" with our server
	// http.ListenAndServe(":9000", sm) // the bind address - binds to all ip addresses on the laptop. second option is a handler. if set to "nil", it will use the default handler which is the http server mux
	go func() {
			err := s.ListenAndServe()
			if err != nil {
				l.Fatal(err)
			}
	}()

	// attempt a gracefull shutdown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, greaceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second) // when called, wait for 30 seconds for running processes to finish and if not done after 30 seconds, forcefully shutdown
	s.Shutdown(tc)
}