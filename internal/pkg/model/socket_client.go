package model

import "github.com/gorilla/websocket"

type SocketClient struct {
	UserID           string
	SocketConnection *websocket.Conn
}
