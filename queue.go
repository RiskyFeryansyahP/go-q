package main

type queueStr struct {
	Size int
	Head string
	Tail string
	Data []string
}

func NewQueue() *queueStr {
	queue := &queueStr{
		Size: 0,
		Head: "",
		Tail: "",
		Data: make([]string, 0),
	}

	return queue
}

func (q *queueStr) PushQueue(value string) {
	q.Data = append(q.Data, value)
	q.Size = q.Size + 1
	q.Head = q.Data[0]
	q.Tail = q.Data[q.Size-1]
}

func (q *queueStr) EnQueue() string {
	// no need to enqueu, because
	// the data is 0
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

func (q queueStr) Length() int {
	return q.Size
}

func (q queueStr) IsEmpty() bool {
	return q.Size == 0
}

func (q queueStr) HeadValue() string {
	return q.Head
}

func (q queueStr) TailValue() string {
	return q.Tail
}
