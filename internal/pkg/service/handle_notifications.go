package service

import (
	"log"
	"mvm_backend/internal/pkg/generated/mvmPb"

	"google.golang.org/protobuf/encoding/protojson"
)

func (s *mvmService) HandleNotifications() {
	log.Print("Start listening for notifications")
	for {
		var msg mvmPb.Notification
		msg = <-NotificationsBroadcaster

		Clients.mu.Lock()
		for _, peer := range Clients.clients {
			if peer.ID == msg.UserId {
				jsonString := protojson.Format(&msg)
				message := &Message{
					Type:   "notification",
					ToId:   peer.ID,
					FromId: msg.FromUser,
					Data:   jsonString,
				}
				forwardMessage(peer.ID, message)
			}
		}
		Clients.mu.Unlock()
	}
}
