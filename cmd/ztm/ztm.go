package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zer0-hex/ztm/internal/runner"
)

func main() {
	// 初始化参数
	ztm, err := runner.NewRunner(&runner.Options{})
	if err != nil {
		panic(err)
	}
	ztm.Flag()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// Setup close handler
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal, Exiting...")
		//pdtmRunner.Close()
		os.Exit(0)
	}()

	err = ztm.Run()
	if err != nil {
		log.Fatalf("Could not run ztm: %s\n", err)
	}
}
