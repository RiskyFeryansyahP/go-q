package main

type Queue[T any] interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value T)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() T

	// Length get the size of the queue
	// Length() int

	// IsEmpty check if the queue is empty
	// IsEmpty() bool

	// HeadValue get the front / head value of the queue
	// HeadValue() T

	// TailValue get the rear / tail value of the queue
	// TailValue() T

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	// FindIndex(value T) (int, bool)
}

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

type QueueInt32 interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value int32)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() int32

	// Length get the size of the queue
	Length() int

	// IsEmpty check if the queue is empty
	IsEmpty() bool

	// HeadValue get the front / head value of the queue
	HeadValue() int32

	// TailValue get the rear / tail value of the queue
	TailValue() int32

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	FindIndex(value int32) (int, bool)
}

type QueueInt64 interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value int64)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() int64

	// Length get the size of the queue
	Length() int

	// IsEmpty check if the queue is empty
	IsEmpty() bool

	// HeadValue get the front / head value of the queue
	HeadValue() int64

	// TailValue get the rear / tail value of the queue
	TailValue() int64

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	FindIndex(value int64) (int, bool)
}

type QueueFloat32 interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value float32)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() float32

	// Length get the size of the queue
	Length() int

	// IsEmpty check if the queue is empty
	IsEmpty() bool

	// HeadValue get the front / head value of the queue
	HeadValue() float32

	// TailValue get the rear / tail value of the queue
	TailValue() float32

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	FindIndex(value float32) (int, bool)
}

type QueueFloat64 interface {
	// Enqueue add new element add store it in rear / end
	// position of the queue
	Enqueue(value float64)

	// Dequeue get element at front of the queue
	// and remove the element at front of the queue
	Dequeue() float64

	// Length get the size of the queue
	Length() int

	// IsEmpty check if the queue is empty
	IsEmpty() bool

	// HeadValue get the front / head value of the queue
	HeadValue() float64

	// TailValue get the rear / tail value of the queue
	TailValue() float64

	// FindIndex find the index of given value as param
	// and return the index if value matches with value
	// in the queue
	FindIndex(value float64) (int, bool)
}
