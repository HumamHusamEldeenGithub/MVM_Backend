package model

import "mvm_backend/internal/pkg/generated/mvmPb"

type SocketMessage struct {
	RoomID    string
	UserID    string
	Message   *string
	Keypoints []*mvmPb.Keypoint
}
