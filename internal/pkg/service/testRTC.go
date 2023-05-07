package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type          string      `json:"type"`
	FromId        string      `json:fromId`
	ToId          string      `json:"toId"`
	Data          interface{} `json:"data"`
	IceCandidates []string    `json:"iceCandidates"`
}

type OfferMessage struct {
	SDP string `json:"sdp"`
}

type CandidateMessage struct {
	Candidate string `json:"candidate"`
}

type ResponseMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Client struct {
	ID            string
	Connection    *websocket.Conn
	ICECandidates []string
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
		handleError(conn, "client connection has been terminated ", http.StatusUnauthorized)
		return
	}

	// Register client
	client := &Client{
		ID:         userID,
		Connection: conn,
	}
	Clients[client.ID] = client
	log.Printf("Registered client %s\n", client.ID)

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
		case "offer":
			// Forward offer to other client
			message.FromId = userID
			toID := message.ToId
			err = forwardMessage(toID, &message)
			if err != nil {
				log.Println("Failed to forward offer:", err)
			}
			//break
		case "answer":
			// Forward answer to other client
			message.FromId = userID
			toID := message.ToId
			err = forwardMessage(toID, &message)
			if err != nil {
				log.Println("Failed to forward answer:", err)
			}
			//break
		case "ice":
			// Add ICE candidate to client
			fmt.Println(message.Data)
			data := message.Data.(string)
			clientID := userID
			client := Clients[clientID]
			client.ICECandidates = append(client.ICECandidates, data)
			log.Printf("Added ICE candidate to client %s\n", clientID)
			//break

		case "getIce":
			message.Data = Clients[message.ToId].ICECandidates
			err = forwardMessage(userID, &message)
			if err != nil {
				log.Println("Failed to forward answer:", err)
			}

		default:
			log.Println("Unknown message type:", message.Type)
			break
		}
	}
	// Close connection and remove client from list
	for clientID, client := range Clients {
		if client.Connection == conn {
			log.Printf("Unregistered client %s\n", clientID)
			delete(Clients, clientID)
			break
		}
	}
}

func forwardMessage(clientID string, message *Message) error {
	log.Printf("Forward %s to %s", *&message.Type, clientID)
	client := Clients[clientID]
	if client == nil {
		return fmt.Errorf("client %s not found", clientID)
	}
	err := client.Connection.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}
