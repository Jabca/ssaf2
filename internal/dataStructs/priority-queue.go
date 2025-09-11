package data_structs

import (
	"container/heap"
)

type pqItem[T any] struct {
	value    *T
	priority int
	index    int
}

type PriorityQueue[T any] []*pqItem[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i int, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	item := x.(*pqItem[T])
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	n := len(*pq)
	old := *pq
	item := *old[n-1]
	item.index = -1
	old[n-1] = nil

	*pq = old[0 : n-1]
	return item.value
}

// check fot interface implementation
var _ heap.Interface = (*PriorityQueue[int])(nil)

// add item into PriorityQueue
func (pq *PriorityQueue[T]) Enqueue(value *T, priority int) {
	item := &pqItem[T]{
		value:    value,
		priority: priority,
	}
	heap.Push(pq, item)
}

// get last item and it's rating in queue and delete it. Returns (stuff, -1 , false) if queue was empty
func (pq *PriorityQueue[T]) Dequeue() (value *T, priority int, success bool) {
	if pq.Len() == 0 {
		var empty *T
		return empty, -1, false
	}
	item := heap.Pop(pq).(*pqItem[T])
	return item.value, item.priority, true
}

// change priority of item
func (pq *PriorityQueue[T]) Update(item *pqItem[T], newPriority int) {
	item.priority = newPriority
	heap.Fix(pq, item.index)
}
