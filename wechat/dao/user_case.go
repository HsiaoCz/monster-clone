package dao

import (
	"context"

	"github.com/HsiaoCz/monster-clone/wechat/types"
)

type UserCaser interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}
