package server

import "strconv"

// 棋子
type Commander int

const (
	Flag      Commander = iota // *1
	Miner                      // *3
	Platoon                    // Platoon Commander *3
	Company                    // Company Commander *3
	Battalion                  // Battalion Commander *2
	Regiment                   // Regiment Commander *2
	Brigade                    // Brigade Commander *2
	Division                   // Division Commander *2
	Army                       // Army Commander *1
	Chief                      // Commander In Chief *1
	Bomb                       // *2
	Mine                       // *3
)

func (c Commander) Id() string {
	return strconv.Itoa(int(c))
}

func (c Commander) Value() int {
	return int(c)
}

type Movable int

const (
	Non  Movable = iota // unmovable
	Step                // one step
	Run                 // straight railway
	Fly                 // fly on railway
)
