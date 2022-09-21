package user

import (
	"military-chess/chess"
)

type User struct {
	Id     string
	Player chess.Player
	//Server *server.Server
}
