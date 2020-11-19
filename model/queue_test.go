package model_test

import (
	"container/heap"
	"testing"

	model "challenge.kamontat.net/crea-model"
)

func TestCreateQueue(t *testing.T) {
	// graph := model.NewGraph()

	// nodeA := graph.AddNode("A")
	// nodeB := graph.AddNode("B")
	// nodeC := graph.AddNode("C")
	// nodeD := graph.AddNode("D")

	// nodeA.AddEdge(nodeB, int64(5))
	// nodeA.AddEdge(nodeC, int64(7))
	// nodeA.AddEdge(nodeD, int64(9))

	nodeA := model.NewNode("A")
	queue := model.NewPriorityQueue(nodeA)
	if queue.Len() != 1 {
		t.Errorf("Expected initial queue length 1, but got %d\n", queue.Len())
	}
}

func TestSupportPopFromHeap(t *testing.T) {
	nodeA := model.NewNode("A")
	queue := model.NewPriorityQueue(nodeA)
	heap.Init(queue)

	item := heap.Pop(queue).(*model.QueueItem)
	if item == nil {
		t.Errorf("Expected queue is popable from heap, but got nil")
	}
}

func TestSupportPushFromHeap(t *testing.T) {
	nodeA := model.NewNode("A")
	nodeB := model.NewNode("B")
	edge := nodeA.AddEdge(nodeB, 10)

	queue := model.NewPriorityQueue(nodeA)
	heap.Init(queue)

	heap.Push(queue, *edge)
	if queue.Len() != 2 {
		t.Errorf("Expected queue been push, but got %+v\n", queue)
	}
}
