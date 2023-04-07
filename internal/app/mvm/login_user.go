package mvm

import (
	"context"
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
)

func (s *MVMServiceServer) LoginUser(ctx context.Context, req *v1.LoginUserRequest) (*v1.LoginUserResponse, error) {
	token, err := s.service.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginUserResponse{
		Token: token,
	}, nil
}
