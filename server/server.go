package server

import (
	"errors"
	"military-chess/chess"
	"military-chess/user"
	"sync"
)

type Status int

const (
	Opening Status = iota
	Ready
	Fighting
	Finished
)

func (s Status) String() string {
	switch s {
	case Opening:
		return "Opening"
	case Fighting:
		return "Fighting"
	case Finished:
		return "Finished"
	default:
		return ""
	}
}

type Server struct {
	world        map[chess.Player]*chess.Territory
	alivePieces  map[chess.Player][]*chess.Piece
	allPlayers   map[chess.Player]*user.User
	playerStatus map[chess.Player]Status
	mutex        sync.RWMutex
	//
	status Status
	//轮到谁了
	turn chess.Player
}

// 红黑对弈
func (s *Server) TwoPlayerGame() *Server {
	// map
	s.world = chess.DrawAMap(chess.Red, chess.Black)
	s.alivePieces = map[chess.Player][]*chess.Piece{}
	s.alivePieces[chess.Red] = chess.Red.DrawAllPieces()
	s.alivePieces[chess.Black] = chess.Black.DrawAllPieces()
	// 初始化棋子的位置 TODO

	s.allPlayers = map[chess.Player]*user.User{chess.Red: nil, chess.Black: nil}
	s.turn = chess.Red
	s.status = Opening
	return s
}

// 动作
func (s *Server) action(msg *user.Message) (err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	switch msg.Action {
	case user.Ready:
		if s.status != Opening {
			return errors.New("server is " + s.status.String())
		}

		_, ok := s.allPlayers[msg.Player]
		if ok {
			s.playerStatus[msg.Player] = Ready
			ready := true
			for k, _ := range s.allPlayers {
				v, _ := s.playerStatus[k]
				if v != Ready {
					ready = false
				}
			}
			if ready { // 如果所有玩家都准备好了，就开启战斗
				s.status = Fighting
			}
		} else {
			return errors.New("you're not in this server")
		}

	case user.Move:
		if s.turn != msg.Player {
			return errors.New("it's not your turn")
		}

	case user.Join:
		if s.status != Opening {
			return errors.New("server is " + s.status.String())
		}
		us := s.allPlayers[msg.Player]
		if us == nil {
			s.allPlayers[msg.Player] = msg.User
		} else {
			return errors.New(msg.Player.Name() + " player is " + us.Id)
		}
	}

	return
}

func (s *Server) run() {
	select {}
}
