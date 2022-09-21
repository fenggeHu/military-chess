package user

import (
	"military-chess/chess"
)

type Type int

const (
	Login Type = iota
	Join
	Ready
	Move
)

type MovingData struct {
	Piece *chess.Piece
	From  *chess.Point
	To    *chess.Point
}

// 通讯
type Message struct {
	*User
	Action Type
	Data   interface{}
}
