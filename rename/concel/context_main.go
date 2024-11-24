package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	DoSomethin(ctx)
}

func DoSomethin(ctx context.Context) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("context err : ", err)
		fmt.Println("ctx is cancel")
	case <-ch:
		fmt.Println("do something done")
	}
}
