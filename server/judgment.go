package server

import "errors"

// 判定结果
type Verdict int

const (
	Equal Verdict = iota
	Success
	Failure
)

// 移动

// 相遇判定
func (p Piece) Meet(e *Piece) (ver Verdict, err error) {
	if p.Commander == Flag || p.Commander == Mine {
		err = errors.New("non movable")
		return
	}
	if p.Commander == Bomb || p.Commander == e.Commander {
		ver = Equal
		return
	}

	if p.Commander > e.Commander {
		ver = Success
	} else {
		ver = Failure
	}
	return
}
