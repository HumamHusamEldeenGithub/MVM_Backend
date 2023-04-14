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
				// TODO : Delete client
				//delete(Clients, Clients2[msg.UserID])
			}
		}
		// if Clients2[msg.UserID] != nil {
		// 	if err := Clients2[msg.UserID].WriteJSON(msg); err != nil {
		// 		log.Printf("error: %v", err)
		// 		Clients2[msg.UserID].Close()
		// 		delete(Clients, Clients2[msg.UserID])
		// 	}
		// }

	}
}
