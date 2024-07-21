package service

import "context"

type PriceFetcher interface {
	FetchPrice(context.Context, string)
}

type DefaultPriceFetcher struct{}

func (d *DefaultPriceFetcher) FetchPrice(ctx context.Context, ticker string) {}
