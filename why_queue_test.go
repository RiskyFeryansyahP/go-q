package main

import (
	"testing"
)

func TestQueueSerobot(t *testing.T) {
	q := NewQueue[int]()
	data := []int{1, 2, 3, 4, 5}
	for _, d := range data {
		q.Enqueue(d)
	}
	q.Serobot(100)
	if q.TailValue() == 100 {
		t.Errorf("want tail not 100")
	}
}

func TestQueueOrDal(t *testing.T) {
	q := NewQueue[int]()
	data := []int{1, 2, 3, 4, 5}
	for _, d := range data {
		q.Enqueue(d)
	}
	q.OrDal(100)
	if q.HeadValue() != 100 {
		t.Errorf("want head to be 100, got %v", q.HeadValue())
	}
}
