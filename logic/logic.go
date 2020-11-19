package logic

import (
	"container/heap"
	"errors"

	model "challenge.kamontat.net/crea-model"
)

// ShortestPath is dijkstra implementation to get shortest path from graph
func ShortestPath(graph model.Graph, start string, end string) ([]*model.Node, error) {
	visited := make(map[*model.Node]bool)
	travel := make(map[*model.Node]int64)
	prev := make(map[*model.Node]*model.Node)

	startNode := graph.GetNode(start)
	endNode := graph.GetNode(end)

	travel[startNode] = 0
	queue := model.NewPriorityQueue(startNode)
	heap.Init(queue)

	for queue.Len() > 0 {
		// Done.
		if value := visited[endNode]; value == true {
			break
		}

		item := heap.Pop(queue).(*model.QueueItem)

		node := item.GetNode()
		edge := node.GetEdges()
		for _, edge := range edge {
			dest := edge.GetNode()
			distance := travel[node] + edge.GetDistance()

			if maybeDistance, ok := travel[dest]; !ok || distance < maybeDistance {
				travel[dest] = distance
				prev[dest] = node
				heap.Push(queue, *edge)
			}
		}
		visited[node] = true
	}

	if value := visited[endNode]; value != true {
		return nil, errors.New("no shortest path exists")
	}

	path := []*model.Node{endNode}
	for next := prev[endNode]; next != nil; next = prev[next] {
		path = append(path, next)
	}

	// Reverse path.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, nil
}
