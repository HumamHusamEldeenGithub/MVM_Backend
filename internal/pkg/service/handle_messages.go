package service

import (
	"fmt"
	"log"
	"mvm_backend/internal/pkg/generated/mvmPb"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func (s *mvmService) HandleMessages() {
	var msg mvmPb.SocketMessage2
	for {
		msg = <-Broadcaster
		// Serialize the Protobuf message into a byte slice
		serializedMessage, err := proto.Marshal(&msg)
		if err != nil {
			fmt.Println("Error reading message from WebSocket:", err)
			continue
		}
		for _, client := range Rooms[msg.RoomId] {
			if client.UserID != msg.UserId {
				continue
			}
			fmt.Println("SEND")
			if err := client.SocketConnection.WriteMessage(websocket.BinaryMessage, serializedMessage); err != nil {
				log.Printf("error: %v", err)
				client.SocketConnection.Close()
				s.LeaveRoom(msg.RoomId, client.UserID)
				deleteUserFromRoom(msg.RoomId, client.UserID)
			}
		}
	}
}
