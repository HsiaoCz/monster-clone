package biz

import (
	"context"
	"time"

	v1 "github.com/HsiaoCz/monster-clone/monster/api/helloworld/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Post struct {
	ID          string    `bson:"_id,omitempty"`
	UserID      string    `bson:"userID"`
	Title       string    `bson:"title"`
	Content     string    `bson:"content"`
	CreatedTime time.Time `bson:"createdTime"`
	Location    string    `bson:"location"`
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Post) (*Post, error)
	Update(context.Context, *Post) (*Post, error)
	FindByID(context.Context, string) (*Post, error)
	ListByUserID(context.Context, string) ([]*Post, error)
	ListAll(context.Context) ([]*Post, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Post) (*Post, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.CreatedTime)
	return uc.repo.Save(ctx, g)
}
