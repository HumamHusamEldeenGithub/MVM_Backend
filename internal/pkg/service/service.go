package service

import (
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"net/http"

	"github.com/gorilla/websocket"
)

var Rooms = make(map[string][]*model.SocketClient)

var Broadcaster = make(chan *mvmPb.SocketMessage2)

var Upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {

		return true

	},
}
var Clients map[string]*model.SocketClient = make(map[string]*model.SocketClient)

type IMVMStore interface {
	CreateUser(user *model.User) (string, error)

	GetProfile(id string, withPassword bool) (*model.User, error)
	GetUserByUsername(username string, withPassword bool) (*model.User, error)
	SearchForUsers(searchInput string) ([]*model.User, error)

	CreateRoom(room *mvmPb.Room) (*mvmPb.Room, error)
	GetRooms(searchQuery string) ([]*mvmPb.Room, error)
	GetUserRooms(userId string) ([]*mvmPb.Room, error)
	GetRoom(id string) (*mvmPb.Room, error)
	DeleteRoom(roomId string) error
	JoinRoom(roomId, userId string) error
	LeaveRoom(roomId, userId string) error
	CreateRoomInvitation(roomId, recipientId string) error
	DeleteRoomInvitation(roomId, recipientId string) error

	GetFriends(userID string) ([]string, error)
	CreateFriendRequest(userID, friendID string) error
	DeleteFriendRequest(userID, friendID string) error
	AddFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
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
