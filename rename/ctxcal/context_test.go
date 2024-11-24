package ctxcal

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx1, cancel := context.WithCancel(ctx)

	// ch := make(chan struct{})
	// cancel()
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	for {
		ch := make(chan string, 1)
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1 is done")
			ch <- "hello"
		default:
			fmt.Println("context dont cancel...")
		}
		c, ok := <-ch
		if ok {
			log.Println("<-ch :", c)
			break
		}
	}

}
