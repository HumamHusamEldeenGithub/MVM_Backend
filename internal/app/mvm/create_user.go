package mvm

import (
	"context"
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
	"mvm_backend/internal/pkg/utils"
)

func (s *MVMServiceServer) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.Empty, error) {
	if err := s.service.CreateUser(ctx, utils.ParseUser(req)); err != nil {
		return nil, err
	}
	return &v1.Empty{}, nil
}
