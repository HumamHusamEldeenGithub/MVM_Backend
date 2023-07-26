package model

import "mvm_backend/internal/pkg/generated/mvmPb"

type AvatarSettings struct {
	ID       string
	Settings *mvmPb.AvatarSettings
}
