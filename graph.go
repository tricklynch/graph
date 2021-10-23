package graph

import (
	"fmt"
)

type Node struct {
	key  int
	data interface{}
}

type Edge struct {
	start int
	end   int
}

type Graph struct {
	nodes []Node
	edges []Edge
}

func MakeGraph() Graph {
	return Graph{
		make([]Node, 0),
		make([]Edge, 0),
	}
}

func (g *Graph) AddNode(data interface{}) int {
	g.nodes = append(g.nodes, Node{len(g.nodes), data})
	return len(g.nodes) - 1
}

func (g *Graph) AddEdge(start, end int) error {
	if start >= len(g.nodes) {
		return fmt.Errorf("Graph does not contain start node %d", start)
	}
	if end >= len(g.nodes) {
		return fmt.Errorf("Graph does not contain end node %d", end)
	}
	g.edges = append(g.edges, Edge{start, end})
	return nil
}
