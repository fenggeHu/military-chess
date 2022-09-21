package action

import (
	"military-chess/client"
	"military-chess/server"
)

type Type int

const (
	Login Type = iota
	Join
	Ready
	Move
)

type MovingData struct {
	Piece *server.Piece
	From  *server.Point
	To    *server.Point
}

// 通讯
type Message struct {
	*client.User
	Action Type
	Data   interface{}
}
