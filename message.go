package main

import "military-chess/client"

type Action int

const (
	Login Action = iota
	Join
	Moving
	Speak
)

// 通讯
type Message struct {
	client.User
	Action
}
