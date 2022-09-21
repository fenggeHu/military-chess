package client

import (
	"github.com/gorilla/websocket"
	"military-chess/action"
)

type Client struct {
	ID      string
	Socket  *websocket.Conn
	Message chan *action.Message
}
