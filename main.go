package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// handling exit signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		signalType := <-ch

		log.Println("Exit command received. Exiting...")
		log.Println("Signal type : ", signalType)
		os.Exit(0)
	}()

	// registering handlers
	http.HandleFunc("/health", getHealth)

	// starting server
	err := http.ListenAndServe("localhost:8090", nil)
	fmt.Printf("error: %+v", err)
	println("exit")
}

func getHealth(rw http.ResponseWriter, _ *http.Request) {
	_, _ = rw.Write([]byte("OK"))
}
