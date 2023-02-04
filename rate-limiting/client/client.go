package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup
	for i := 0; i < total; i = i + max {
		limit := max
		if i+max > total {
			limit = total - i
		}

		wg.Add(1)
		for j := 0; j < limit; j++ {
			go func(j int) {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("Could not dial: %v", err)
				}

				dataBytes, err := ioutil.ReadAll(conn)
				if err != nil {
					log.Fatalf("Could not read data: %v", err)
				}

				if string(dataBytes) != "SUCCESS" {
					log.Fatalf("request error; request: %d", i+1+j)
				}
				fmt.Printf("request %d: %s\n", i+1+j, string(dataBytes))
			}(j)
		}
		wg.Wait()
	}
}
