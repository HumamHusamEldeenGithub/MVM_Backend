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

	Rooms[roomId] = append(Rooms[roomId], &model.SocketClient{UserID: userID, SocketConnection: ws})

	if err := s.JoinRoom(roomId, userID); err != nil {
		handleError(ws, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Client %s has been authorized\nAnd connected to room : %s\n", userID, roomId)

	for {

		var msg mvmPb.SimpleSocketMessage

		// Read in a new message as JSON and map it to a Message object
		err = ws.ReadJSON(&msg)
		if err != nil {
			s.LeaveRoom(roomId, userID)
			deleteUserFromRoom(roomId, userID)
			break
		}

		fmt.Println(&msg)

		Broadcaster <- model.SocketMessage{
			UserID:    userID,
			RoomID:    roomId,
			Message:   msg.Message,
			Keypoints: msg.Keypoints,
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
