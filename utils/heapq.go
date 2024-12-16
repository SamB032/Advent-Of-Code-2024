package util

import (
	"container/heap"
)

type Item[T any] struct {
	Value    *T  // The value of the item
	Priority int // The priority of the item (used to determine ordering)
}

type PriorityQueue[T any] []*Item[T] // Define a heap type for the queue

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	item := x.(*Item[T])
	*pq = append(*pq, item)
}

// Pop is part of the heap.Interface and removes the item with the highest priority (min-heap)
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// NewPriorityQueue creates a new priority queue with a specified compare function
func NewPriorityQueue[T any](items []*Item[T]) *PriorityQueue[T] {
	pq := PriorityQueue[T](items)
	heap.Init(&pq)
	return &pq
}
