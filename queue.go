package main

type queue[T any] struct {
	// hold of the length of queue
	Size int

	// hold value of head / front element in the queue
	Head T

	// hold value of rear / tail element in the queue
	Tail T

	// store or hold all element value in the queue
	Data []T
}

// NewQueue initialize new queue
func NewQueue[T any]() Queue[T] {
	queue := &queue[T]{
		Size: 0,
		Data: make([]T, 0),
	}

	return queue
}

// Enqueue add new element add store it in rear / end
// position of the queue
func (q *queue[T]) Enqueue(value T) {
	q.Data = append(q.Data, value)
	q.Size = q.Size + 1
	q.Head = q.Data[0]
	q.Tail = q.Data[q.Size-1]
}

// Dequeue get element at front of the queue
// and remove the element at front of the queue
func (q *queue[T]) Dequeue() T {
	var t T

	// no need to enqueu, because
	// length of the data is 0 / empty
	if q.Size == 0 {
		return t
	}

	enqueueVal := q.Data[0]

	q.Data = q.Data[1:]
	q.Size = q.Size - 1

	// emtpy head and tail value
	// because the data was empty
	if q.Size == 0 {
		q.Head = t
		q.Tail = t
	}

	// move head value to next data
	if q.Size > 0 {
		q.Head = q.Data[0]
	}

	return enqueueVal
}
