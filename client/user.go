package client

import "military-chess/server"

type User struct {
	Id string
	server.Player
	Server *server.Server
}
