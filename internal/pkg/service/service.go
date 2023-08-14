package service

import (
	"fmt"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/utils"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var Rooms = make(map[string][]*model.SocketClient)

var NotificationsBroadcaster = make(chan *mvmPb.Notification)

var Upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {

		return true

	},
}

type ClientsMutex struct {
	clients map[string]*model.SocketClient
	mu      sync.Mutex
}

var Clients = ClientsMutex{
	clients: make(map[string]*model.SocketClient),
	mu:      sync.Mutex{},
}

type IMVMStore interface {
	CreateUser(user *model.User) (string, error)

	GetProfile(id string, withPassword bool) (*model.User, error)
	GetProfiles(ids []string) ([]*model.User, error)
	GetUserByUsername(username string, withPassword bool) (*model.User, error)
	SearchForUsers(searchInput string) ([]*model.User, error)
	UpsertAvatarSettings(id string, settings *mvmPb.AvatarSettings) error
	GetAvatarSettings(id string) (*mvmPb.AvatarSettings, error)

	CreateRoom(room *mvmPb.Room) (*mvmPb.Room, error)
	GetRooms(searchQuery string) ([]*mvmPb.Room, error)
	GetUserRooms(userId string) ([]*mvmPb.Room, error)
	GetRoom(id string) (*mvmPb.Room, error)
	DeleteRoom(roomId string) error
	JoinRoom(roomId, userId string) error
	LeaveRoom(roomId, userId string) error
	CreateRoomInvitation(roomId, recipientId string) error
	DeleteRoomInvitation(roomId, recipientId string) error

	GetFriends(userID string) (*model.Friends, error)
	GetPendingFriends(userID string) ([]string, error)
	CreateFriendRequest(userID, friendID string) error
	DeleteFriendRequest(userID, friendID string) error
	AddFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error

	CreateNotification(notification *mvmPb.Notification) (*mvmPb.Notification, error)
	GetNotifications(userID string) ([]*mvmPb.Notification, error)
	DeleteNotifications(userID string) error
}

type IMVMAuth interface {
	GenerateToken(user *model.User, refreshToken bool) (*jwt_manager.JWTToken, error)
	VerifyToken(token string, isRefreshToken bool) (string, error)
}

type mvmService struct {
	store IMVMStore
	auth  IMVMAuth
}

func NewMVMService(store IMVMStore, auth IMVMAuth) *mvmService {
	return &mvmService{
		store: store,
		auth:  auth,
	}
}

func (s *mvmService) CreateFriendRequestNotification(fromUser, toUser string) (*mvmPb.Notification, error) {
	profile, err := s.GetProfile(fromUser)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("%s has sent you a friend request", profile.Username)

	return &mvmPb.Notification{
		Id:       utils.GenerateID(),
		UserId:   toUser,
		Type:     int32(mvmPb.NotificationType_FRIEND_REQUEST),
		FromUser: fromUser,
		Message:  &msg,
	}, nil
}

func (s *mvmService) CreateRoomInvitationNotification(fromUser, toUser, roomId string) (*mvmPb.Notification, error) {
	profile, err := s.GetProfile(fromUser)
	if err != nil {
		return nil, err
	}

	room, err := s.GetRoom(roomId)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("%s has sent you a room invitation to %s room", profile.Username, room.Title)

	return &mvmPb.Notification{
		Id:       utils.GenerateID(),
		UserId:   toUser,
		Type:     int32(mvmPb.NotificationType_ROOM_INVITATION),
		FromUser: fromUser,
		Message:  &msg,
	}, nil
}

func (s *mvmService) CreateAcceptFriendRequestNotification(fromUser, toUser string) (*mvmPb.Notification, error) {
	profile, err := s.GetProfile(fromUser)
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("%s has accepted your friend request", profile.Username)

	return &mvmPb.Notification{
		Id:       utils.GenerateID(),
		UserId:   toUser,
		Type:     int32(mvmPb.NotificationType_ACCEPT_REQUEST),
		FromUser: fromUser,
		Message:  &msg,
	}, nil
}
