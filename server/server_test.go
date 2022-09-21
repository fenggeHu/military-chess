package server

import (
	"fmt"
	"military-chess/chess"
	"military-chess/user"
	"testing"
)

func TestServer(t *testing.T) {
	s := &Server{}
	s.RandomGame()

	msg := &user.Message{
		Action: user.Ready,
		User:   &user.User{Player: chess.Red},
	}
	s.action(msg)
	fmt.Println(s)
}
