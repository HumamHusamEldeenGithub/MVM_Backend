package mvm

import (
	"context"
	"mvm_backend/internal/pkg/errors"
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
	"mvm_backend/internal/pkg/utils"
)

func (s *MVMServiceServer) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	if len(utils.GetUserID(ctx)) == 0 {
		return nil, errors.Errorf(errors.Unauthenticated)
	}

	user, err := s.service.GetUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserResponse{
		Username: user.Username,
	}, nil
}
