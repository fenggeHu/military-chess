package chess

import (
	"strconv"
)

// 玩家
type Player int

// 红黑对战，四国：橙色、紫色、绿色、蓝色 - 橙绿、紫蓝分别一对
const (
	White Player = iota // 观众
	Red
	Black
	// 待扩展
	//Orange
	//Purple
	//Green
	//Blue
)

func (p Player) Name() string {
	switch p {
	case Red:
		return "Red"
	case Black:
		return "Black"
	case White:
		return "White"
	default:
		return ""
	}
}

func (p Player) DrawAllPieces() []*Piece {
	player := strconv.Itoa(int(p))
	ps := make([]*Piece, 25)
	ps[0] = &Piece{Commander: Flag, Player: p, uid: player + Flag.Id()}
	ps[1] = &Piece{Commander: Miner, Player: p, uid: player + Miner.Id() + "A"}
	ps[2] = &Piece{Commander: Miner, Player: p, uid: player + Miner.Id() + "B"}
	ps[3] = &Piece{Commander: Miner, Player: p, uid: player + Miner.Id() + "C"}
	ps[4] = &Piece{Commander: Platoon, Player: p, uid: player + Platoon.Id() + "A"}
	ps[5] = &Piece{Commander: Platoon, Player: p, uid: player + Platoon.Id() + "B"}
	ps[6] = &Piece{Commander: Platoon, Player: p, uid: player + Platoon.Id() + "C"}
	ps[7] = &Piece{Commander: Company, Player: p, uid: player + Company.Id() + "A"}
	ps[8] = &Piece{Commander: Company, Player: p, uid: player + Company.Id() + "B"}
	ps[9] = &Piece{Commander: Company, Player: p, uid: player + Company.Id() + "C"}
	ps[10] = &Piece{Commander: Battalion, Player: p, uid: player + Battalion.Id() + "A"}
	ps[11] = &Piece{Commander: Battalion, Player: p, uid: player + Battalion.Id() + "B"}
	ps[12] = &Piece{Commander: Regiment, Player: p, uid: player + Regiment.Id() + "A"}
	ps[13] = &Piece{Commander: Regiment, Player: p, uid: player + Regiment.Id() + "B"}
	ps[14] = &Piece{Commander: Brigade, Player: p, uid: player + Brigade.Id() + "A"}
	ps[15] = &Piece{Commander: Brigade, Player: p, uid: player + Brigade.Id() + "B"}
	ps[16] = &Piece{Commander: Division, Player: p, uid: player + Division.Id() + "A"}
	ps[17] = &Piece{Commander: Division, Player: p, uid: player + Division.Id() + "B"}
	ps[18] = &Piece{Commander: Army, Player: p, uid: player + Army.Id()}
	ps[19] = &Piece{Commander: Chief, Player: p, uid: player + Chief.Id()}
	ps[20] = &Piece{Commander: Bomb, Player: p, uid: player + Bomb.Id() + "A"}
	ps[21] = &Piece{Commander: Bomb, Player: p, uid: player + Bomb.Id() + "B"}
	ps[22] = &Piece{Commander: Mine, Player: p, uid: player + Mine.Id() + "A"}
	ps[23] = &Piece{Commander: Mine, Player: p, uid: player + Mine.Id() + "B"}
	ps[24] = &Piece{Commander: Mine, Player: p, uid: player + Mine.Id() + "C"}
	return ps
}
