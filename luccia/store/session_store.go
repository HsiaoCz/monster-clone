package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"gorm.io/gorm"
)

type SessionStorer interface {
	CreateSession(context.Context, *st.Session) (*st.Session, error)
	GetSessionByToken(context.Context, string) (*st.Session, error)
}

type SessionStore struct {
	db *gorm.DB
}

func SessionStoreInit(db *gorm.DB) *SessionStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateSession(ctx context.Context, session *st.Session) (*st.Session, error) {
	tx := s.db.Debug().WithContext(ctx).Model(&st.Session{}).Create(session)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return session, nil
}

func (s *SessionStore) GetSessionByToken(ctx context.Context, token string) (*st.Session, error) {
	var session st.Session
	tx := s.db.Debug().WithContext(ctx).Model(&st.Session{}).Where("token = ?", token).First(&session)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &session, nil
}
