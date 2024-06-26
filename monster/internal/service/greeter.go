package service

import (
	"context"

	v1 "github.com/HsiaoCz/monster-clone/monster/api/helloworld/v1"
	"github.com/HsiaoCz/monster-clone/monster/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) CreatePost(ctx context.Context, in *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	return &v1.CreatePostResponse{}, nil
}

func (s *GreeterService) DeletePost(ctx context.Context, in *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	return &v1.DeletePostResponse{}, nil
}

func (s *GreeterService) GetPostByIDRequest(ctx context.Context, in *v1.GetPostByIDRequest) (*v1.GetPostByIDResponse, error) {
	return &v1.GetPostByIDResponse{}, nil
}

func (s *GreeterService) ListPost(ctx context.Context, in *v1.ListPostRequest) (*v1.ListPostResponse, error) {
	return &v1.ListPostResponse{}, nil
}
