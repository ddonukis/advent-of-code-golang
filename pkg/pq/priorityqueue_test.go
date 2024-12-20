package pq_test

import (
	"testing"

	"github.com/ddonukis/advent-of-code-golang/pkg/pq"
)

func TestIntPriorityQueue(t *testing.T) {
	pq := pq.NewIntPriorityQueue(99, 31, 121, 9, 57)

	t.Logf("before init: %s\n", pq)

	expected := [5]int{9, 31, 57, 99, 121}
	for i, expectedItem := range expected {
		item := pq.Pop()
		t.Logf("got %v after %dth pop: %v\n", item, i+1, pq)
		if item != expectedItem {
			t.Fatalf("got: %v, expected: %v\n", item, expectedItem)
		}
	}

	t.Log("Pushing new items into the priority queue...")
	pq.Push(10)
	pq.Push(5)
	pq.Push(15)
	pq.Push(6)
	pq.Push(5)
	pq.Push(7)

	for i, expectedItem := range []int{5, 5, 6, 7, 10, 15} {
		item := pq.Pop()
		t.Logf("got %v after %dth pop: %v\n", item, i+1, pq)
		if item != expectedItem {
			t.Fatalf("got: %v, expected: %v\n", item, expectedItem)
		}
	}

}
