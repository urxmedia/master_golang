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
	go work()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Done Waiting; Main Exits!!!")
	// join point
}
