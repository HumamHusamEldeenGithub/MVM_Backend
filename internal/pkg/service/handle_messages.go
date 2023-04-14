package service

import (
	"fmt"
	"log"
)

func (s *mvmService) HandleMessages() {
	for {
		msg := <-Broadcaster
		fmt.Println(msg)
		if Clients2[msg.UserID] != nil {
			if err := Clients2[msg.UserID].WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				Clients2[msg.UserID].Close()
				delete(Clients, Clients2[msg.UserID])
			}
		}

	}
}
