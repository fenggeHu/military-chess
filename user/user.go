package user

import (
	"military-chess/chess"
)

type User struct {
	Id string
	chess.Player
	//Server *server.Server
}
