package services

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/HsiaoCz/monster-clone/fetch/types"
)

type FetchPricer interface {
	FetchPrice(context.Context, string) (*types.PriceResponse, error)
}

type fetcher struct{}

func NewFetcher() FetchPricer {
	return &fetcher{}
}

func (f *fetcher) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": rand.Float64(),
	"ETH": rand.Float64(),
	"GG":  rand.Float64(),
}

func MockPriceFetcher(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return nil, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return &types.PriceResponse{Ticker: ticker, Price: price}, nil
}
