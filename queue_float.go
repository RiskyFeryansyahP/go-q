package main

type queueFloat32 struct {
	// hold of the length of queue
	Size int

	// hold value of head / front element in the queue
	Head float32

	// hold value of rear / tail element in the queue
	Tail float32

	// store or hold all element value in the queue
	Data []float32
}

// NewQueueInt initialize new queue
func NewQueueFloat32() QueueFloat32 {
	queue := &queueFloat32{
		Size: 0,
		Head: 0,
		Tail: 0,
		Data: make([]float32, 0),
	}

	return queue
}

// Enqueue add new element add store it in rear / end
// position of the queue
func (q *queueFloat32) Enqueue(value float32) {
	q.Data = append(q.Data, value)
	q.Size = q.Size + 1
	q.Head = q.Data[0]
	q.Tail = q.Data[q.Size-1]
}

// Dequeue get element at front of the queue
// and remove the element at front of the queue
func (q *queueFloat32) Dequeue() float32 {
	// no need to enqueu, because
	// length of the data is 0 / empty
	// and return -1 that means no data in queue
	if q.Size == 0 {
		return -1
	}

	enqueueVal := q.Data[0]

	q.Data = q.Data[1:]
	q.Size = q.Size - 1

	// emtpy head and tail value
	// because the data was empty
	if q.Size == 0 {
		q.Head = 0
		q.Tail = 0
	}

	// move head value to next data
	if q.Size > 0 {
		q.Head = q.Data[0]
	}

	return enqueueVal
}

// Length get the size of the queue
func (q queueFloat32) Length() int {
	return q.Size
}

// IsEmpty check if the queue is empty
func (q queueFloat32) IsEmpty() bool {
	return q.Size == 0
}

// HeadValue get the front / head value of the queue
func (q queueFloat32) HeadValue() float32 {
	return q.Head
}

// TailValue get the rear / tail value of the queue
func (q queueFloat32) TailValue() float32 {
	return q.Tail
}

// FindIndex find the index of given value as param
// and return the index if value matches with value
// in the queue
func (q queueFloat32) FindIndex(value float32) (int, bool) {
	for k, v := range q.Data {
		if v == value {
			return k, true
		}
	}

	// no match value found
	// so we return -1 and false
	return -1, false
}

type queueFloat64 struct {
	// hold of the length of queue
	Size int

	// hold value of head / front element in the queue
	Head float64

	// hold value of rear / tail element in the queue
	Tail float64

	// store or hold all element value in the queue
	Data []float64
}

// NewQueueInt32 initialize new queue
func NewQueueFloat64() QueueFloat64 {
	queue := &queueFloat64{
		Size: 0,
		Head: 0,
		Tail: 0,
		Data: make([]float64, 0),
	}

	return queue
}

// Enqueue add new element add store it in rear / end
// position of the queue
func (q *queueFloat64) Enqueue(value float64) {
	q.Data = append(q.Data, value)
	q.Size = q.Size + 1
	q.Head = q.Data[0]
	q.Tail = q.Data[q.Size-1]
}

// Dequeue get element at front of the queue
// and remove the element at front of the queue
func (q *queueFloat64) Dequeue() float64 {
	// no need to enqueu, because
	// length of the data is 0 / empty
	// and return -1 that means no data in queue
	if q.Size == 0 {
		return -1
	}

	enqueueVal := q.Data[0]

	q.Data = q.Data[1:]
	q.Size = q.Size - 1

	// emtpy head and tail value
	// because the data was empty
	if q.Size == 0 {
		q.Head = 0
		q.Tail = 0
	}

	// move head value to next data
	if q.Size > 0 {
		q.Head = q.Data[0]
	}

	return enqueueVal
}

// Length get the size of the queue
func (q queueFloat64) Length() int {
	return q.Size
}

// IsEmpty check if the queue is empty
func (q queueFloat64) IsEmpty() bool {
	return q.Size == 0
}

// HeadValue get the front / head value of the queue
func (q queueFloat64) HeadValue() float64 {
	return q.Head
}

// TailValue get the rear / tail value of the queue
func (q queueFloat64) TailValue() float64 {
	return q.Tail
}

// FindIndex find the index of given value as param
// and return the index if value matches with value
// in the queue
func (q queueFloat64) FindIndex(value float64) (int, bool) {
	for k, v := range q.Data {
		if v == value {
			return k, true
		}
	}

	// no match value found
	// so we return -1 and false
	return -1, false
}
