package dao

import (
	"context"

	"github.com/HsiaoCz/monster-clone/wechat/types"
)

type SessionCaser interface {
	CreateSession(context.Context, *types.Sessions) (*types.Sessions, error)
}
