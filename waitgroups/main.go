package main

import (
	"fmt"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Task", id, "is done.")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	now := time.Now()
	for i := 1; i <= 10; i++ {
		go work(&wg, i)
	}
	wg.Wait()
	fmt.Println("Elapsed:", time.Since(now))
	fmt.Println("DONE WAITING... EXITING MAIN!!!")
}
