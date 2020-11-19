package model

import "container/heap"

// QueueItem is a priority queue item
type QueueItem struct {
	value    *Node // The value of the item; arbitrary.
	priority int64 // The priority of the item in the queue.
	index    int   // The index of the item in the heap.
}

// GetNode will return value of queue item
func (i *QueueItem) GetNode() *Node {
	return i.value
}

// PriorityQueue is type of queue with priority
type PriorityQueue struct {
	queue []*QueueItem
}

// NewPriorityQueue will create new priority queue with default start node
func NewPriorityQueue(start *Node) *PriorityQueue {
	queue := []*QueueItem{}
	queue = append(queue, &QueueItem{
		value:    start,
		priority: 0,
		index:    0,
	})

	return &PriorityQueue{
		queue: queue,
	}
}

// Len is queue length
func (q PriorityQueue) Len() int {
	return len(q.queue)
}

// Less will check is first given index less than second index
func (q PriorityQueue) Less(i, j int) bool {
	return q.queue[i].priority < q.queue[j].priority
}

// Swap will swap element in given index
func (q PriorityQueue) Swap(i, j int) {
	q.queue[i], q.queue[j] = q.queue[j], q.queue[i]
	q.queue[i].index = i
	q.queue[j].index = j
}

// Push will add Edge to queue
func (q *PriorityQueue) Push(x interface{}) {
	edge := x.(Edge)
	item := &QueueItem{
		value:    edge.GetNode(),
		priority: edge.GetDistance(),
		index:    q.Len(),
	}

	q.queue = append(q.queue, item)
}

// Pop will get first element out
func (q *PriorityQueue) Pop() interface{} {
	old := q.queue
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	q.queue = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (q *PriorityQueue) update(item *QueueItem, value *Node, priority int64) {
	item.value = value
	item.priority = priority
	heap.Fix(q, item.index)
}
