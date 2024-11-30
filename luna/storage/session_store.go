package storage

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luna/types"
	"gorm.io/gorm"
)

type SessionStoreInter interface {
	CreateSession(context.Context, *types.Sessions) (*types.Sessions, error)
	GetSessionByID(context.Context, string) (*types.Sessions, error)
}

type SessionStore struct {
	db *gorm.DB
}

func SessionStoreInit(db *gorm.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateSession(ctx context.Context, session *types.Sessions) (*types.Sessions, error) {
	return nil, nil
}

func (s *SessionStore) GetSessionByID(ctx context.Context, session_id string) (*types.Sessions, error) {
	return nil, nil
}
