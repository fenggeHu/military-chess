package client

import (
	"github.com/gorilla/websocket"
	"military-chess/user"
)

type Client struct {
	ID      string
	Socket  *websocket.Conn
	Message chan *user.Message
}
