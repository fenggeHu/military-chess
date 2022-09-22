package chess

import (
	"container/list"
	"military-chess/graph"
	"strconv"
)

type Map struct {
	graph.Graph
}

func (m *Map) AddEdgeNodes(a, b *graph.Node, weight int) bool {
	if a == b {
		return false
	}
	edge := graph.NewEdge(a, b, weight)
	return m.AddEdge(edge)
}

func (m *Map) AddEdge(e *graph.Edge) bool {
	if m.HasEdge(e) {
		return false
	} else {
		m.Edges = append(m.Edges, e)
		return true
	}
}

func (m *Map) AddNode(n *graph.Node) bool {
	if m.HasNode(n) {
		return false
	} else {
		m.Nodes = append(m.Nodes, n)
		return true
	}
}
func (m *Map) HasEdge(e *graph.Edge) bool {
	for _, d := range m.Edges {
		if d == e {
			return true
		}
	}
	return false
}

func (m *Map) HasNode(n *graph.Node) bool {
	for _, d := range m.Nodes {
		if d == n {
			return true
		}
	}
	return false
}

// return
func (m *Map) Paths(a, b *graph.Node) []*list.List {
	ret := make([]*list.List, 0)
	// 直连
	path := m.directed(a, b)
	if path != nil {
		l := list.New()
		l.PushBack(path)
		ret = append(ret, l)
		return ret
	}
	// 遍历

	return ret
}

// 非直连的节点，只要edge的weight =0
// avoid 不能通行的节点
func (m *Map) undirected(a, b *graph.Node, avoid []*graph.Node, paths *graph.Path) {
	for _, e := range m.Edges {
		if e.Weight > 0 || paths.Contains(e) {
			continue
		}

		if e.Connect(a, b) {
			p := e.Path(a)
			paths.AddChild(p)
			return
		} else if e.Has(a) {
			if !contains(avoid, e.Another(a)) { // 排除节点
				p := e.Path(a)
				paths.AddChild(p)
				m.undirected(e.Another(a), b, avoid, paths)
			}
		}
	}
}

func contains(nodes []*graph.Node, e *graph.Node) bool {
	for _, v := range nodes {
		if v == e {
			return true
		}
	}
	return false
}

func (m *Map) directed(a, b *graph.Node) *graph.Path {
	for _, e := range m.Edges {
		if e.Connect(a, b) {
			return e.Path(a)
		}
	}
	return nil
}

type NodeType int

const (
	Normal   NodeType = iota
	Barrack           // 大本营
	FlagNode          //军旗点
)

type RoadType int

const (
	Railway RoadType = iota
	Pathway
)

// 一个点
type Point struct {
	*graph.Node
	Who *Piece //
}

type Territory struct {
	Points [6][5]Point
	Index  int // 表示位置/顺序
}

func pointId(p Player, x, y int) string {
	return strconv.Itoa(int(p)) + "_" + strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

// draw a player territory
func DrawATerritory(p Player) *Territory {
	var land = [6][5]Point{}
	land[0] = line(p, 0)
	land[1] = line(p, 1)
	land[2] = line(p, 2)
	land[3] = line(p, 3)
	land[4] = line(p, 4)
	land[5] = line(p, 5)

	return &Territory{Points: land}
}

// define a line
func line(p Player, x int) [5]Point {
	return [5]Point{{&graph.Node{Id: pointId(p, x, 0)}, nil},
		{&graph.Node{Id: pointId(p, x, 1)}, nil},
		{&graph.Node{Id: pointId(p, x, 2)}, nil},
		{&graph.Node{Id: pointId(p, x, 3)}, nil},
		{&graph.Node{Id: pointId(p, x, 4)}, nil},
	}
}

// draw a map for players
func DrawAMap(players ...Player) map[Player]*Territory {
	m := make(map[Player]*Territory)
	for _, p := range players {
		ter := DrawATerritory(p)
		ter.Index = int(p)
		m[p] = ter
	}
	return m
}
