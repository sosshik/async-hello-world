package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Hello world")
	start := time.Now()
	timer := time.After(10 * time.Second)
	interruptionChannel := make(chan os.Signal, 1)
	signal.Notify(interruptionChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-interruptionChannel:
		fmt.Printf("Stopped by the user after %v seconds\n", time.Since(start).Seconds())
		return
	case <-timer:
		fmt.Println("Goodbye world")
		return
	}
}
