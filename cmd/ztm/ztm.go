package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var options Options
	Flag(&options)
	options := runner.ParseOptions()
	pdtmRunner, err := runner.NewRunner(options)
	if err != nil {
		log.Fatal("Could not create runner: %s\n", err)
	}
	log.Fatal()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Setup close handler
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal, Exiting...")
		pdtmRunner.Close()
		os.Exit(0)
	}()

	err = pdtmRunner.Run()
	if err != nil {
		log.Fatal("Could not run pdtm: %s\n", err)
	}
}
