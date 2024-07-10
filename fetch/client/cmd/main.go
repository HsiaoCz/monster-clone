package main

import (
	"context"

	"github.com/HsiaoCz/monster-clone/fetch/client"
)

func main() {
	client := client.NewClient("http://localhsot:3001/price")
	client.FetchPrice(context.Background(), "gg")
}
