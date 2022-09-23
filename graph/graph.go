package graph

import (
	"container/list"
	"strings"
)

// node value type
// Node is a graph node
type Node struct {
	Id    string
	value interface{}
}

func NewNode(id string, v interface{}) *Node {
	return &Node{id, v}
}

// Edge - undirected edge
type Edge struct {
	// 小id + 大id
	Id string
	// the double nodes of the edge.
	nodes [2]*Node
	// type
	Weight int
}

func NewEdge(a, b *Node, weight int) *Edge {
	id := undirectedEdgeId(a, b)
	return &Edge{id, [2]*Node{a, b}, weight}
}

// create undirected edge id
func undirectedEdgeId(a, b *Node) string {
	if strings.Compare(a.Id, b.Id) > 0 {
		return b.Id + "_" + a.Id
	} else {
		return a.Id + "_" + b.Id
	}
}

// include
func (e *Edge) Has(n *Node) bool {
	return e.nodes[0] == n || e.nodes[1] == n
}

// 取另一个节点
func (e *Edge) Another(n *Node) *Node {
	if e.nodes[0] == n {
		return e.nodes[1]
	} else if e.nodes[1] == n {
		return e.nodes[0]
	}
	return nil
}

// Connectivity
func (e *Edge) Connect(a, b *Node) bool {
	return (e.nodes[0] == a && e.nodes[1] == b) || (e.nodes[0] == b && e.nodes[1] == a)
}

func (e *Edge) Path(from *Node) *Path {
	var to *Node
	if e.nodes[0] == from {
		to = e.nodes[1]
	} else if e.nodes[1] == from {
		to = e.nodes[0]
	} else {
		return nil
	}
	return &Path{from, to, e, nil}
}

// 有向的
type Path struct {
	// from
	F *Node
	// to
	T *Node
	// edge
	edge *Edge
	// sub
	children []*Path
}

func (p *Path) Contains(e *Edge) bool {
	if p.edge == e {
		return true
	} else if p.children != nil {
		for _, c := range p.children {
			if c.Contains(e) {
				return true
			}
		}
	}

	return false
}

// 到达终点的所有有效路线转成行
//func (p *Path) ToRow(end *Node, current []*Path) [][]*Path {
//	if p.T == end {
//		current = append(current, p)
//		return nil
//	}
//	if p.children == nil {
//		return nil
//	}
//
//	ret := make([][]*Path, 0)
//	for _, c := range p.children {
//		newPath := make([]*Path, len(current))
//		copy(newPath, current)
//		newPath = append(newPath, c)
//		fork := c.ToRow(end, newPath)
//		if fork == nil {
//			if len(newPath) > 0 {
//				pe := newPath[len(newPath)-1]
//				if pe.T == end {
//					ret = append(ret, newPath)
//				}
//			}
//		} else {
//			// TODO
//		}
//	}
//	return ret
//}

func (p *Path) AddChild(e *Path) {
	if p.children == nil {
		p.children = make([]*Path, 0)
	}
	p.children = append(p.children, e)
}

// Graph is a generalized graph.
type Graph struct {
	// all Nodes in graph
	Nodes []*Node
	// all edges in graph
	Edges []*Edge
}

type IGraph interface {
	// add edge & nodes
	AddEdgeNodes(a, b *Node, weight int) bool

	AddEdge(e *Edge) bool
	//
	AddNode(n *Node) bool

	HasEdge(e *Edge) bool
	//
	HasNode(n *Node) bool
	//
	GetNode(id string) *Node
	// 合并
	Combine(g2 *Graph) *Graph
	// all paths of two nodes
	Paths(a, b *Node) []*list.List
}

func (g *Graph) AddEdgeNodes(a, b *Node, weight int) bool {
	if a == b {
		return false
	}
	edge := NewEdge(a, b, weight)
	return g.AddEdge(edge)
}

func (g *Graph) AddEdge(e *Edge) bool {
	if g.HasEdge(e) {
		return false
	} else {
		g.Edges = append(g.Edges, e)
		return true
	}
}

func (g *Graph) AddNode(n *Node) bool {
	if g.HasNode(n) {
		return false
	} else {
		g.Nodes = append(g.Nodes, n)
		return true
	}
}
func (g *Graph) HasEdge(e *Edge) bool {
	for _, d := range g.Edges {
		if d == e {
			return true
		}
	}
	return false
}

func (g *Graph) HasNode(n *Node) bool {
	for _, d := range g.Nodes {
		if d == n {
			return true
		}
	}
	return false
}

func (g *Graph) GetNode(id string) *Node {
	for _, v := range g.Nodes {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (g *Graph) Combine(g2 *Graph) *Graph {
	nodes := make([]*Node, 0)
	for _, v := range g.Nodes {
		nodes = append(nodes, v)
	}
	for _, v := range g2.Nodes {
		nodes = append(nodes, v)
	}
	edges := make([]*Edge, 0)
	for _, v := range g.Edges {
		edges = append(edges, v)
	}
	for _, v := range g2.Edges {
		edges = append(edges, v)
	}

	// all edges in graph
	return &Graph{nodes, edges}
}
func (g *Graph) Paths(a, b *Node) []*list.List {

	return nil
}
