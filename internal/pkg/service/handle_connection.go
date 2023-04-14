package service

import (
	"fmt"
	"log"
	"mvm_backend/internal/pkg/payloads"
	"net/http"
)

func (s *mvmService) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()

	fmt.Println("Client has been connected ")

	isAuthorized := false

	for {
		// if Clients[ws].UserID == "" {
		// 	var tokenMessage payloads.InitSocketMessage
		// 	err := ws.ReadJSON(&tokenMessage)
		// 	if err != nil || tokenMessage.Token == "" {
		// 		delete(Clients, ws)
		// 		fmt.Println("Client connection has been terminated ")
		// 		break
		// 	}
		// 	userID, err := s.auth.VerifyToken(tokenMessage.Token, false)
		// 	if err != nil {
		// 		delete(Clients, ws)
		// 		fmt.Println("Client connection has been terminated ")
		// 		break
		// 	}
		// 	client := Clients[ws]
		// 	client.UserID = userID
		// }

		if !isAuthorized {
			var tokenMessage payloads.InitSocketMessage
			err := ws.ReadJSON(&tokenMessage)
			if err != nil || tokenMessage.Token == "" {
				delete(Clients, ws)
				fmt.Println("Client connection has been terminated ")
				break
			}
			userID, err := s.auth.VerifyToken(tokenMessage.Token, false)
			if err != nil {
				delete(Clients, ws)
				fmt.Println("Client connection has been terminated ")
				break
			}
			Clients2[userID] = ws
			isAuthorized = true
		}

		var msg payloads.SokcetMessage

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(Clients, ws)
			break
		}
		// send new message to the channel
		Broadcaster <- msg
	}
}
