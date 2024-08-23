package models

import (
	"github.com/gorilla/websocket"
)

type Room struct {
	ID      string
	Clients map[*websocket.Conn]bool
}

type Message struct {
	Type     string `json:"type"`
	SenderID string `json:"senderId"`
	Content  string `json:"content"`
}
