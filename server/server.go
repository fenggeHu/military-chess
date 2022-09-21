package server

import (
	"errors"
	"military-chess/action"
)

type Server struct {
	world        map[Player]*Territory
	alivePieces  map[Player][]*Piece
	alivePlayers []*Player
	// 开盘了
	running bool
	//轮到谁了
	turn Player
}

func (s *Server) RandomGame() *Server {
	// map
	s.world = DrawAMap(Red, Black)
	s.alivePieces[Red] = Red.DrawAllPieces()
	s.alivePieces[Black] = Black.DrawAllPieces()
	s.turn = Red
	return s
}

// 动作
func (s *Server) action(msg *action.Message) (err error) {
	//
	if msg.Action == action.Ready {
		if s.running {
			return errors.New("server is running")
		}
		isReady := false
		for _, p := range s.alivePlayers {
			if *p == msg.Player {
				isReady = true
			}
		}
		if !isReady {
			s.alivePlayers = append(s.alivePlayers, &msg.Player)
		}
		return
	}
	//
	if s.turn != msg.Player {
		return errors.New("it's not your turn")
	}
	// todo

	return
}

func (s *Server) run() {
	select {}
}
