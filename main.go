package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Hello world")
	start := time.Now()
	timer := time.NewTimer(10 * time.Second)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		end := time.Now()
		timer.Stop()
		fmt.Printf("Stopped by the user after %v seconds\n", end.Sub(start).Seconds())
		os.Exit(1)
	}()
	<-timer.C
	fmt.Println("Goodbye world")
}
