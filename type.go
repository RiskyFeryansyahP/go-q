package main

type QueueStr interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value string)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() string

	// Length get the size of the queue
	Length() int

	// IsEmpty check if the queue is empty
	IsEmpty() bool

	// HeadValue get the front / head value of the queue
	HeadValue() string

	// TailValue get the rear / tail value of the queue
	TailValue() string

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	FindIndex(value string) (int, bool)
}

type QueueInt interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value int)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() int

	// Length get the size of the queue
	Length() int

	// IsEmpty check if the queue is empty
	IsEmpty() bool

	// HeadValue get the front / head value of the queue
	HeadValue() int

	// TailValue get the rear / tail value of the queue
	TailValue() int

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	FindIndex(value int) (int, bool)
}
