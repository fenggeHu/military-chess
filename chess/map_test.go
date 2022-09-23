package chess

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	players := []Player{Red, Black}
	m := DrawAMap(players...)
	fmt.Println(m)
	red := m[Red]
	red.Points[0][0].Who = &Piece{
		Commander: Bomb,
		Player:    Black,
	}
	fmt.Println(red)

	a := m

}
