package pq

import (
	"container/heap"
	"fmt"
	"slices"
)

type IntPriorityQueue struct {
	storage *heapStorage
}

type heapStorage []int

func (pq heapStorage) Len() int {
	return len(pq)
}

func (pq heapStorage) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq heapStorage) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *heapStorage) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *heapStorage) Pop() any {
	n := len(*pq) - 1
	item := (*pq)[n]
	*pq = slices.Delete(*pq, n, n+1)
	return item
}

func NewIntPriorityQueue(item ...int) IntPriorityQueue {
	s := make(heapStorage, len(item))
	copy(s, item)
	heap.Init(&s)
	return IntPriorityQueue{storage: &s}
}

func (pq IntPriorityQueue) Push(item int) {
	heap.Push(pq.storage, item)
}

func (pq IntPriorityQueue) Pop() int {
	return heap.Pop(pq.storage).(int)
}

func (pq IntPriorityQueue) String() string {
	return fmt.Sprintf("%v", *pq.storage)
}

func (pq IntPriorityQueue) Len() int {
	return pq.storage.Len()
}
