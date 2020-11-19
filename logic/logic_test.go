package logic_test

import (
	"testing"

	logic "challenge.kamontat.net/crea-logic"
	model "challenge.kamontat.net/crea-model"
)

func TestShortestPathWithDirectShorter(t *testing.T) {
	mockGraph := model.NewGraph()
	nodeA := mockGraph.AddNode("A")
	nodeB := mockGraph.AddNode("B")
	nodeC := mockGraph.AddNode("C")
	nodeD := mockGraph.AddNode("D")

	nodeA.AddEdge(nodeB, int64(5))
	nodeB.AddEdge(nodeC, int64(5))
	nodeC.AddEdge(nodeD, int64(2))
	nodeA.AddEdge(nodeD, int64(10))

	nodes, err := logic.ShortestPath(*mockGraph, "A", "D")
	if err != nil {
		t.Errorf("Expected pass, but got %+v\n", err)
	}

	if len(nodes) != 2 {
		t.Errorf("Expected 0 stops, but got %+v\n", nodes)
	}
}

func TestShortestPathWithDirectLonger(t *testing.T) {
	mockGraph := model.NewGraph()
	nodeA := mockGraph.AddNode("A")
	nodeB := mockGraph.AddNode("B")
	nodeC := mockGraph.AddNode("C")
	nodeD := mockGraph.AddNode("D")

	nodeA.AddEdge(nodeB, int64(5))
	nodeB.AddEdge(nodeC, int64(5))
	nodeC.AddEdge(nodeD, int64(2))
	nodeA.AddEdge(nodeD, int64(13))

	nodes, err := logic.ShortestPath(*mockGraph, "A", "D")
	if err != nil {
		t.Errorf("Expected pass, but got %+v\n", err)
	}

	if len(nodes) != 4 {
		t.Errorf("Expected 2 stops, but got %+v\n", nodes)
	}
}

func TestErrorWhenNoShortestPath(t *testing.T) {
	mockGraph := model.NewGraph()
	nodeA := mockGraph.AddNode("A")
	nodeB := mockGraph.AddNode("B")
	nodeC := mockGraph.AddNode("C")
	nodeD := mockGraph.AddNode("D")

	nodeA.AddEdge(nodeB, int64(5))
	nodeA.AddEdge(nodeC, int64(7))
	nodeC.AddEdge(nodeD, int64(9))

	nodes, err := logic.ShortestPath(*mockGraph, "B", "D")
	if err == nil {
		t.Errorf("Expected error, but got %+v\n", err)
	}

	if nodes != nil {
		t.Errorf("Expected return node list is nil, but got %+v\n", nodes)
	}
}
