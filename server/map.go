package server

type RoadType int

const (
	Pathway RoadType = iota
	Railway
	Barrack
)

// 一个点
type Point struct {
	Road RoadType
	Who  *Piece
}

type Territory struct {
	Points [6][5]Point
	Index  int // 表示位置/顺序
}

// draw a player territory
func DrawATerritory() *Territory {
	var land = [6][5]Point{}
	land[0] = [5]Point{{Road: Railway}, {Road: Railway}, {Road: Railway}, {Road: Railway}, {Road: Railway}}
	land[1] = [5]Point{{Road: Railway}, {Road: Barrack}, {Road: Pathway}, {Road: Barrack}, {Road: Railway}}
	land[2] = [5]Point{{Road: Railway}, {Road: Pathway}, {Road: Barrack}, {Road: Pathway}, {Road: Railway}}
	land[3] = [5]Point{{Road: Railway}, {Road: Barrack}, {Road: Pathway}, {Road: Barrack}, {Road: Railway}}
	land[4] = [5]Point{{Road: Railway}, {Road: Railway}, {Road: Railway}, {Road: Railway}, {Road: Railway}}
	land[5] = [5]Point{{Road: Pathway}, {Road: Pathway}, {Road: Pathway}, {Road: Pathway}, {Road: Pathway}}

	return &Territory{Points: land}
}

// draw a map for players
func DrawAMap(players ...Player) map[Player]*Territory {
	m := make(map[Player]*Territory)
	for _, p := range players {
		ter := DrawATerritory()
		ter.Index = int(p)
		m[p] = ter
	}
	return m
}
