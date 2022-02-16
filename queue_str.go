package main

type queueStr struct {
	// hold of the length of queue
	Size int

	// hold value of head / front element in the queue
	Head string

	// hold value of rear / tail element in the queue
	Tail string

	// store or hold all element value in the queue
	Data []string
}

// NewQueue initialize new queue
func NewQueue() QueueStr {
	queue := &queueStr{
		Size: 0,
		Head: "",
		Tail: "",
		Data: make([]string, 0),
	}

	return queue
}

// Enqueue add new element add store it in rear / end
// position of the queue
func (q *queueStr) Enqueue(value string) {
	q.Data = append(q.Data, value)
	q.Size = q.Size + 1
	q.Head = q.Data[0]
	q.Tail = q.Data[q.Size-1]
}

// Dequeue get element at front of the queue
// and remove the element at front of the queue
func (q *queueStr) Dequeue() string {
	// no need to enqueu, because
	// length of the data is 0 / empty
	if q.Size == 0 {
		return ""
	}

	enqueueVal := q.Data[0]

	q.Data = q.Data[1:]
	q.Size = q.Size - 1

	// emtpy head and tail value
	// because the data was empty
	if q.Size == 0 {
		q.Head = ""
		q.Tail = ""
	}

	// move head value to next data
	if q.Size > 0 {
		q.Head = q.Data[0]
	}

	return enqueueVal
}

// Length get the size of the queue
func (q queueStr) Length() int {
	return q.Size
}

// IsEmpty check if the queue is empty
func (q queueStr) IsEmpty() bool {
	return q.Size == 0
}

// HeadValue get the front / head value of the queue
func (q queueStr) HeadValue() string {
	return q.Head
}

// TailValue get the rear / tail value of the queue
func (q queueStr) TailValue() string {
	return q.Tail
}

// FindIndex find the index of given value as param
// and return the index if value matches with value
// in the queue
func (q queueStr) FindIndex(value string) (int, bool) {
	for k, v := range q.Data {
		if v == value {
			return k, true
		}
	}

	// no match value found
	// so we return -1 and false
	return -1, false
}
