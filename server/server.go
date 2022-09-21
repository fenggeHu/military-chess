package server

type Server struct {
	Map          map[Player]*Territory
	AlivePieces  map[Player][]*Piece
	AlivePlayers []*Player
	// 轮到谁了
	Index int
}

func (s *Server) RandomGame() {
	// map
	s.Map = DrawAMap(Red, Black)
	allPieces := make(map[Player][]*Piece)
	allPieces[Red] = Red.DrawAllPieces()
	allPieces[Black] = Black.DrawAllPieces()
}

func (s *Server) run() {

}
