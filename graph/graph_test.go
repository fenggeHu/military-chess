package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	a := NewNode("1_0_0", nil)
	b := NewNode("1_0_1", nil)
	graph := NewGraph(nil, nil)
	graph.AddEdgeNodes(a, b, 0)
	c := NewNode("1_0_2", nil)
	graph.AddEdgeNodes(b, c, 0)
	d := NewNode("1_0_3", nil)
	graph.AddEdgeNodes(c, d, 0)
	e := NewNode("1_0_4", nil)
	graph.AddEdgeNodes(d, e, 0)

	fmt.Println(graph)
}
