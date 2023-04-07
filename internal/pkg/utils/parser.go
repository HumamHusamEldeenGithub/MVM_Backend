package utils

import (
	v1 "mvm_backend/internal/pkg/generated/mvm-api/v1"
	"mvm_backend/internal/pkg/model"
)

func ParseUser(req *v1.CreateUserRequest) *model.User {
	return &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}
