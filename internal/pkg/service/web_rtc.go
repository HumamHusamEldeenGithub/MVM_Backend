package service

import (
	"encoding/json"
	"fmt"
	"log"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/encoding/protojson"
)

type Message struct {
	Type          string      `json:"type"`
	FromId        string      `json:fromId`
	ToId          string      `json:"toId"`
	Data          interface{} `json:"data"`
	IceCandidates []string    `json:"iceCandidates"`
}

type OnlineStatus struct {
	ID       string `json:"id"`
	IsOnline bool   `json:"isOnline"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *mvmService) HandleWebSocketRTC(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	log.Printf("Client has been connected ")

	authHeader := r.Header.Get("Authorization")
	tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
	userID, err := s.auth.VerifyToken(tokenString, false)
	if err != nil {
		handleError(conn, "client connection has been terminated , invalid token.", http.StatusUnauthorized)
		return
	}

	friends, err := s.GetFriends(userID)
	if err != nil {
		// TODO : handle err
		fmt.Println(err)
	}

	// Register client
	client := &model.SocketClient{
		ID:         userID,
		Connection: conn,
		Friends:    friends,
	}

	Clients.mu.Lock()
	Clients.clients[client.ID] = client
	log.Printf("Registered client %s\n", client.ID)
	Clients.mu.Unlock()

	profile, err := s.GetProfile(client.ID)
	if err != nil {
		handleError(conn, "user profile not found", http.StatusNotFound)
		return
	}
	Clients.clients[client.ID].Profile = profile

	// get online friends
	s.GetOnlineFriendStatus(userID)

	// push online event
	PushOnlineStatusToFriends(client, true)

	// Receive and handle messages from client
	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// Parse message
		var message Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Failed to parse message:", err)
			break
		}

		switch message.Type {

		case "join_room":
			// Joining room
			roomId := message.Data.(string)
			log.Printf("user:%s trying to connecting to room: %s  ", client.ID, roomId)

			if len(roomId) == 0 {
				handleError(conn, "room id not found", http.StatusNotFound)
				return
			}

			if err := s.CheckRoomAvailability(roomId, userID); err != nil {
				handleError(conn, "Not authorized to enter this room ", http.StatusUnauthorized)
				return
			}

			Rooms[roomId] = append(Rooms[roomId], client)

			if err := s.JoinRoom(roomId, userID); err != nil {
				handleError(conn, err.Error(), http.StatusInternalServerError)
				return
			}

			client.RoomID = roomId
			// Push user_enter event to all room members
			var message Message = Message{
				Type:   "user_enter",
				FromId: userID,
			}
			forwardMessageToRoom(userID, roomId, &message)

		case "leave_room":
			if len(client.RoomID) != 0 {
				s.LeaveRoom(client.RoomID, userID)
				deleteUserFromRoom(client.RoomID, userID)
				client.RoomID = ""
			}

		case "offer":
			// Forward offer to other client
			message.FromId = userID
			toID := message.ToId
			err = forwardMessage(toID, &message)
			if err != nil {
				log.Println("Failed to forward offer:", err)
			}

		case "answer":
			// Forward answer to other client
			message.FromId = userID
			toID := message.ToId
			err = forwardMessage(toID, &message)
			if err != nil {
				log.Println("Failed to forward answer:", err)
			}
		case "ice":
			// Add ICE candidate to client
			message.FromId = userID
			data := message.Data.(string)
			clientID := userID
			client := Clients.clients[clientID]
			client.ICECandidates = append(client.ICECandidates, data)
			log.Printf("Added ICE candidate to client %s\n", clientID)

			forwardMessageToRoom(userID, client.RoomID, &message)

		case "getIce":
			Clients.mu.Lock()
			message.Data = Clients.clients[message.ToId].ICECandidates
			Clients.mu.Unlock()
			err = forwardMessage(userID, &message)
			if err != nil {
				log.Println("Failed to forward answer:", err)
			}

		default:
			log.Println("Unknown message type:", message.Type)
		}
	}

	// push offline status
	PushOnlineStatusToFriends(client, false)

	// Close connection and remove client from list
	Clients.mu.Lock()
	for clientID, client := range Clients.clients {
		if client.Connection == conn {
			log.Printf("Unregistered client %s\n", clientID)
			delete(Clients.clients, clientID)
			break
		}
	}
	Clients.mu.Unlock()
	// Leave room
	if len(client.RoomID) != 0 {
		s.LeaveRoom(client.RoomID, userID)
		deleteUserFromRoom(client.RoomID, userID)
		client.RoomID = ""
	}

}

func forwardMessage(clientID string, message *Message) error {
	log.Printf("Forward %s to %s", *&message.Type, clientID)
	Clients.mu.Lock()
	client := Clients.clients[clientID]
	Clients.mu.Unlock()
	if client == nil {
		return fmt.Errorf("client %s not found", clientID)
	}
	err := client.Connection.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

func forwardMessageToRoom(clientID, roomID string, message *Message) {
	for _, client := range Rooms[roomID] {
		if clientID == client.ID {
			continue
		}
		if client == nil {
			log.Printf("client %s not found", clientID)
			continue
		}
		log.Printf("Forward %s to %s", *&message.Type, client.ID)
		err := client.Connection.WriteJSON(message)
		if err != nil {
			log.Printf("Error in forwarding  %v ", err)
			continue
		}
	}
}

func deleteUserFromRoom(roomId, userId string) {
	for i, client := range Rooms[roomId] {
		if client.ID == userId {
			Rooms[roomId][i] = Rooms[roomId][len(Rooms[roomId])-1]
			Rooms[roomId] = Rooms[roomId][:len(Rooms[roomId])-1]
		}
	}
}

func PushOnlineStatusToFriends(client *model.SocketClient, isOnline bool) {
	friendMap := make(map[string]bool)
	for _, friend := range client.Friends {
		friendMap[friend] = true
	}

	Clients.mu.Lock()
	for _, peer := range Clients.clients {
		if friendMap[peer.ID] {
			jsonString := protojson.Format(&mvmPb.OnlineStatus{
				Id:       client.ID,
				Username: client.Profile.Username,
				IsOnline: isOnline,
			})
			message := &Message{
				Type:   "user_status_changed",
				ToId:   peer.ID,
				FromId: client.ID,
				Data:   jsonString,
			}
			forwardMessage(peer.ID, message)
		}
	}
	Clients.mu.Unlock()
}

func (s *mvmService) GetOnlineFriendStatus(userID string) {
	friends, err := s.GetFriends(userID)
	if err != nil {
		log.Printf("error: %v", err)
	}

	Clients.mu.Lock()
	Clients.clients[userID].Friends = friends
	Clients.mu.Unlock()

	friendMap := make(map[string]bool)
	for _, friend := range friends {
		friendMap[friend] = true
	}

	onlineStatusList := &mvmPb.OnlineStatuses{}

	Clients.mu.Lock()
	for _, peer := range Clients.clients {
		if friendMap[peer.ID] {
			profile, err := s.GetProfile(peer.ID)
			if err != nil {
				continue
			}
			onlineStatusList.Users = append(onlineStatusList.Users, &mvmPb.OnlineStatus{
				Id:       peer.ID,
				Username: profile.Username,
				IsOnline: true,
			})
		}
	}
	onlineStatusList.Users = append(onlineStatusList.Users, &mvmPb.OnlineStatus{
		Id:       "112233",
		Username: "testingUser",
		IsOnline: true,
	})
	Clients.mu.Unlock()

	jsonString := protojson.Format(onlineStatusList)

	message := &Message{
		Type: "get_users_online_status_list",
		ToId: userID,
		Data: jsonString,
	}
	forwardMessage(userID, message)

}

func handleError(ws *websocket.Conn, message string, code int64) {
	errMsg := &mvmPb.SocketMessage{
		Type: mvmPb.SocketMessageType_ERROR,
		Data: errors.NewSocketError(message, code),
	}
	log.Printf("error: %v", errMsg)
	if err := ws.WriteJSON(errMsg); err != nil {
		log.Printf("error: %v", err)
	}
	ws.Close()
	return
}
