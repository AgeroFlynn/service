package main

import (
	"fmt"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/automaxprocs/maxprocs"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var build = "develop"

func main() {
	// Set the correct number of threads for the services
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		fmt.Printf("maxprocs: %v", err)
		os.Exit(1)
	}

	g := runtime.GOMAXPROCS(0)

	log.Printf("starting sales-api-service build[%s] CPU[%d]", build, g)
	defer log.Println("services ended")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	log.Println("stopping services ")

}
