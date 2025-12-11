package session

import (
	v1 "github.com/titpetric/platform/proto/v1/v1connect"
)

type Service struct {
	v1.UnimplementedClientServiceHandler
}

func NewService() *Service {
	return &Service{}
}

var _ v1.ClientServiceHandler = &Service{}
