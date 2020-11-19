package model

import (
	"errors"
	"fmt"
)

// Node is graph node
type Node struct {
	name  string
	edges map[string]*Edge
}

// NewNode will create new node struct
func NewNode(name string) *Node {
	return &Node{
		name:  name,
		edges: make(map[string]*Edge),
	}
}

// GetName will return name of node
func (n *Node) GetName() string {
	return n.name
}

// AddEdge will create new edge and add to current node
func (n *Node) AddEdge(node *Node, distance int64) *Edge {
	edge := NewEdge(node, distance)

	// n.edges = append(n.edges, edge) // when edge is array
	n.edges[node.name] = edge
	return edge
}

// GetEdges will return all edge from this node
func (n *Node) GetEdges() []*Edge {
	values := []*Edge{}
	for _, value := range n.edges {
		values = append(values, value)
	}
	return values
}

// GetEdge will return edge base on node name
func (n *Node) GetEdge(name string) (*Edge, error) {
	edge := n.edges[name]
	if edge == nil {
		return nil, errors.New("cannot found edge name " + name)
	}

	return edge, nil
}

// ToString will return string represent current node
func (n *Node) ToString() string {
	return fmt.Sprintf("Node(%s)", n.name)
}

// ToStringWithDistance will return string with distance to given name
func (n *Node) ToStringWithDistance(name string) string {
	edge, err := n.GetEdge(name)
	if err != nil {
		return fmt.Sprintf("%s: %s", n.ToString(), err.Error())
	}

	return fmt.Sprintf("%s: %s", n.ToString(), edge.ToString())
}

// Edge is graph edge with distance between node
type Edge struct {
	distance    int64
	destination *Node
}

// NewEdge will create new edge struct
func NewEdge(node *Node, distance int64) *Edge {
	return &Edge{
		destination: node,
		distance:    distance,
	}
}

// GetNode will return edge to node
func (e *Edge) GetNode() *Node {
	return e.destination
}

// GetDistance return how long this edge last
func (e *Edge) GetDistance() int64 {
	return e.distance
}

// ToString will return string represent current node
func (e *Edge) ToString() string {
	return fmt.Sprintf("%d to %s", e.distance, e.destination.ToString())
}

// Graph is collection of nodes
type Graph struct {
	nodes map[string]*Node
}

// NewGraph will create empty graph
func NewGraph() *Graph {
	return &Graph{
		nodes: map[string]*Node{},
	}
}

// AddNode will add new node
func (g *Graph) AddNode(name string) *Node {
	node := NewNode(name)
	g.nodes[name] = node

	return node
}

// GetNode will return named node or create new one if not exist
func (g *Graph) GetNode(name string) *Node {
	node := g.nodes[name]
	if node == nil || node.name == "" {
		return g.AddNode(name)
	}

	return node
}
