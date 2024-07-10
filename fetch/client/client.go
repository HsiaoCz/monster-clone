package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/monster-clone/fetch/types"
)

type Client struct {
	endpoint string
}

func NewClient(endponit string) *Client {
	return &Client{
		endpoint: endponit,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("ticker", ticker)
	req.URL.RawQuery = query.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	priceResp := &types.PriceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil
}
