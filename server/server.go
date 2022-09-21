package server

import (
	"errors"
	"military-chess/chess"
	"military-chess/user"
)

type Server struct {
	world        map[chess.Player]*chess.Territory
	alivePieces  map[chess.Player][]*chess.Piece
	alivePlayers []*chess.Player
	// 开盘了
	running bool
	//轮到谁了
	turn chess.Player
}

func (s *Server) RandomGame() *Server {
	// map
	s.world = chess.DrawAMap(chess.Red, chess.Black)
	s.alivePieces = map[chess.Player][]*chess.Piece{}
	s.alivePieces[chess.Red] = chess.Red.DrawAllPieces()
	s.alivePieces[chess.Black] = chess.Black.DrawAllPieces()

	s.alivePlayers = []*chess.Player{}
	s.turn = chess.Red
	return s
}

// 动作
func (s *Server) action(msg *user.Message) (err error) {
	switch msg.Action {
	case user.Ready:
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
	case user.Move:
		if s.turn != msg.Player {
			return errors.New("it's not your turn")
		}

	case user.Join:
		return
	}

	return
}

func (s *Server) run() {
	select {}
}
