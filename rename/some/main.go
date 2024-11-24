package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx1, cancel := context.WithCancel(ctx)

	// ch := make(chan struct{})
	// cancel()
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	ch := make(chan string, 1)
	for {
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1 is done")
			ch <- "hello"
		default:
			fmt.Println("context dont cancel...")
		}
	}
}
