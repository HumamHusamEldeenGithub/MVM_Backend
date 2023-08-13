package mvm

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"net/http"
)

type MVMServiceServer struct {
	service IMVMService
}

type IMVMService interface {
	LoginUser(username, password string) (string, *jwt_manager.JWTToken, error)
	LoginByRefreshToken(refreshToken string) (*jwt_manager.JWTToken, error)
	CreateUser(user *model.User) (string, *jwt_manager.JWTToken, error)

	GetUserByUsername(username string) (*model.User, error)
	GetProfile(id string) (*model.User, error)
	GetProfiles(ids []string) ([]*model.User, error)
	SearchForUsers(searchInput string) ([]*model.User, error)
	UpsertAvatarSettings(id string, settings *mvmPb.AvatarSettings) error
	GetAvatarSettings(id string) (*mvmPb.AvatarSettings, error)

	CreateRoom(room *mvmPb.Room) (*mvmPb.Room, error)
	GetRooms(searchQuery string) ([]*mvmPb.Room, error)
	GetUserRooms(userId string) ([]*mvmPb.Room, error)
	DeleteRoom(userId, roomId string) error
	CreateRoomInvitation(userId, roomId, recipientId string) error
	DeleteRoomInvitation(userId, roomId, recipientId string) error

	CreateFriendRequest(userID, friendID string) error
	DeleteFriendRequest(userID, friendID string) error
	GetFriends(userID string) ([]string, error)
	GetPendingFriends(userID string) ([]string, error)
	AddFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error

	HandleNotifications()
	GetNotifications(userID string) ([]*mvmPb.Notification, error)
	DeleteNotification(userID, notificationID string) error

	HandleWebSocketRTC(w http.ResponseWriter, r *http.Request)
}

func NewIMVMServiceServer(service IMVMService) *MVMServiceServer {
	return &MVMServiceServer{
		service: service,
	}
}

func encodeAvatarSettings(settings model.AvatarSettings) *mvmPb.AvatarSettings {
	return &mvmPb.AvatarSettings{
		HeadStyle:        settings.Settings.HeadStyle,
		HairStyle:        settings.Settings.HairStyle,
		EyebrowsStyle:    settings.Settings.EyebrowsStyle,
		EyeStyle:         settings.Settings.EyeStyle,
		NoseStyle:        settings.Settings.NoseStyle,
		MouthStyle:       settings.Settings.MouthStyle,
		SkinImperfection: settings.Settings.SkinImperfection,
		Tattoo:           settings.Settings.Tattoo,
		HairColor:        settings.Settings.HairColor,
		BrowsColor:       settings.Settings.BrowsColor,
		SkinColor:        settings.Settings.SkinColor,
		EyeColor:         settings.Settings.EyeColor,
		Gender:           settings.Settings.Gender,
	}
}
