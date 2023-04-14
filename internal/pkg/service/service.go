package service

import (
	"mvm_backend/internal/pkg/jwt_manager"
	"mvm_backend/internal/pkg/model"
	"mvm_backend/internal/pkg/payloads"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]model.SocketClient)
var Clients2 = make(map[string]*websocket.Conn)

var Broadcaster = make(chan payloads.SokcetMessage)

var Upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {

		return true

	},
}

type IMVMStore interface {
	CreateUser(user *model.User) (string, error)

	GetProfile(id string, withPassword bool) (*model.User, error)
	GetUserByUsername(username string, withPassword bool) (*model.User, error)
	SearchForUsers(searchInput string) ([]*model.User, error)

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
