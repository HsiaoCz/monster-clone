package data

import (
	"context"

	"github.com/HsiaoCz/monster-clone/monster/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Post) (*biz.Post, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Post, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Post, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Post, error) {
	return nil, nil
}
