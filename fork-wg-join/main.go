package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("DOING SOME WORK!!!")
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		work()
	}()
	wg.Wait()
	fmt.Println("Elapsed:", time.Since(now))
	fmt.Println("Done Waiting; Main Exits!!!")
}
