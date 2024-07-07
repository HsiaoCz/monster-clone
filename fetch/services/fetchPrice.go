package services

import (
	"context"

	"github.com/HsiaoCz/monster-clone/fetch/types"
)

type FetchPricer interface {
	FetchPrice(context.Context, string) (*types.Price, error)
}
