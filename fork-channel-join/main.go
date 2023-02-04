package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("DOING SOME WORK!!!")
}

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Elapsed:", time.Since(now))
	fmt.Println("Done Waiting; Main Exits!!!")
}
