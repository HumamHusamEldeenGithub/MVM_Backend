package model

import "github.com/gorilla/websocket"

type SocketClient struct {
	ID            string
	Connection    *websocket.Conn
	ICECandidates []string
	RoomID        string
	Friends       []string
}
