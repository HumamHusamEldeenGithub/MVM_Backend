package model

import "mvm_backend/internal/pkg/generated/mvmPb"

type User struct {
	ID             string                `json:"id" bson:"id,omitempty"`
	Username       string                `json:"username" bson:"username"`
	Email          string                `json:"email" bson:"email"`
	Password       string                `json:"password" bson:"password"`
	Friends        []string              `json:"friends" bson:"friends"`
	AvatarSettings *mvmPb.AvatarSettings `json:"avatar_settings" bson:"avatar_settings"`
}
