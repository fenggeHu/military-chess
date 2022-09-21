package server

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	s := &Server{}
	s.RandomGame()
	fmt.Println(s)
}
