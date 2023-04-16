package service

import (
	"fmt"
	"log"
	"mvm_backend/internal/pkg/model"
)

func (s *mvmService) HandleMessages() {
	for {
		var msg model.SocketMessage
		msg = <-Broadcaster
		fmt.Println(msg)
		for _, client := range Rooms[msg.RoomID] {
			if client.UserID == msg.UserID {
				continue
			}
			if err := client.SocketConnection.WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				client.SocketConnection.Close()
				s.LeaveRoom(msg.RoomID, client.UserID)
				deleteUserFromRoom(msg.RoomID, client.UserID)
			}
		}
	}
}
