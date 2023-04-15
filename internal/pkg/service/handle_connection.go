package service

import (
	"fmt"
	"log"
	"mvm_backend/internal/pkg/generated/mvmPb"
	"mvm_backend/internal/pkg/model"
	"net/http"
	"strings"
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
		fmt.Println("Client connection has been terminated ")
		return
	}

	roomId := r.URL.Query().Get("room")

	if len(roomId) == 0 {
		fmt.Println("RoomId is null ,  connection has been terminated ")
		return
	}

	Rooms[roomId] = append(Rooms[roomId], &model.SocketClient{UserID: userID, SocketConnection: ws})

	fmt.Printf("Client %s has been authorized\nAnd connected to room : %s\n", userID, roomId)

	for {

		var msg mvmPb.SocketMessage

		// Read in a new message as JSON and map it to a Message object
		err = ws.ReadJSON(&msg)
		if err != nil {
			//TODO : Remove client from room
			//delete(Rooms[roomID], &model.SocketClient{UserID: userID, SocketConnection: ws})
			break
		}
		socketMessage := model.SocketMessage{
			UserID:    userID,
			RoomID:    roomId,
			Message:   msg.Message,
			Keypoints: msg.Keypoints,
		}

		// send new message to the channel
		Broadcaster <- socketMessage
	}
}
