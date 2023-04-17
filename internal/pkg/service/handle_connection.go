package service

import (
	"fmt"
	"log"
	"mvm_backend/internal/pkg/errors"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func (s *mvmService) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()

	fmt.Println("Client has been connected ")

	authHeader := r.Header.Get("Authorization")
	tokenString := strings.ReplaceAll(authHeader, "Bearer ", "")
	userID, err := s.auth.VerifyToken(tokenString, false)
	if err != nil {
		handleError(ws, "client connection has been terminated ", http.StatusUnauthorized)
		return
	}

	roomId := r.URL.Query().Get("room")

	if len(roomId) == 0 {
		handleError(ws, "room id not found", http.StatusNotFound)
		return
	}

	if err := s.CheckRoomAvailability(roomId, userID); err != nil {
		handleError(ws, "Not authorized to enter this room ", http.StatusUnauthorized)
		return
	}

	Rooms[roomId] = append(Rooms[roomId], &model.SocketClient{UserID: userID, SocketConnection: ws})

	if err := s.JoinRoom(roomId, userID); err != nil {
		handleError(ws, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Client %s has been authorized\nAnd connected to room : %s\n", userID, roomId)

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message from WebSocket:", err)
			s.LeaveRoom(roomId, userID)
			deleteUserFromRoom(roomId, userID)
			break
		}

		// Decode the message if it's a binary message
		if messageType == websocket.BinaryMessage {
			// Create a Protobuf message instance and unmarshal the binary message into it
			var messageObj mvmPb.SimpleSocketMessage
			err := proto.Unmarshal(message, &messageObj)
			if err != nil {
				fmt.Println("Error decoding Protobuf message:", err)
				s.LeaveRoom(roomId, userID)
				deleteUserFromRoom(roomId, userID)
				break
			}

			// Process the message as needed
			// fmt.Printf("Received message: Property1=%s, Property2=%v\n", *messageObj.Message, messageObj.Keypoints)
			// fmt.Println(&messageObj)

			protoMsg := mvmPb.SocketMessage2{
				UserId:    userID,
				RoomId:    roomId,
				Message:   messageObj.Message,
				Keypoints: messageObj.Keypoints,
			}

			Broadcaster <- protoMsg
		}
	}
}

func deleteUserFromRoom(roomId, userId string) {
	for i, client := range Rooms[roomId] {
		if client.UserID == userId {
			Rooms[roomId][i] = Rooms[roomId][len(Rooms[roomId])-1]
			Rooms[roomId] = Rooms[roomId][:len(Rooms[roomId])-1]
		}
	}
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
