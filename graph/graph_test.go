package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	a := NewNode("1_0_0", nil)
	b := NewNode("1_0_1", nil)
	//e := NewEdge(a, b, 0)
	graph := NewGraph(nil, nil)
	graph.AddEdgeNodes(a, b, 0)

	fmt.Println(graph)
}
