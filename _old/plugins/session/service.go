package session

import (
	"context"

	v1 "github.com/titpetric/platform/proto/v1"
)

type Service struct {
	v1.UnimplementedSessionServiceServer
}

func NewService() *Service {
	return &Service{}
}

func (*Service) Get(ctx context.Context, r *v1.GetSessionRequest) (*v1.GetSessionResponse, error) {
	return nil, nil
}

var _ v1.SessionServiceServer = &Service{}
