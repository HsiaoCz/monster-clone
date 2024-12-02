package services

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/peek/types"
	"gorm.io/gorm"
)

type SessionCaser interface {
	CreateSession(context.Context, *types.Session) (*types.Session, error)
	GetSessionByToken(context.Context, string) (*types.Session, error)
}

type SessionCase struct {
	db *gorm.DB
}

func SessionCaseInit(db *gorm.DB) *SessionCase {
	return &SessionCase{
		db: db,
	}
}

func (s *SessionCase) CreateSession(ctx context.Context, session *types.Session) (*types.Session, error) {
	if s.db.WithContext(ctx).Debug().Model(&types.Session{}).Create(session).Error != nil {
		return nil, errors.New("create session failed")
	}
	return session, nil
}

func (s *SessionCase) GetSessionByToken(ctx context.Context, token string) (*types.Session, error) {
	var session types.Session
	tx := s.db.Debug().WithContext(ctx).Model(&types.Session{}).Where("token = ?", token).First(&session)
	if tx.Error!=nil{
		return nil,tx.Error
	}
    return &session,nil
}

