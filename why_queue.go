// this should be changed to package name instead
// but I'm only here for the memes
package main

import (
	"math/rand"
	"time"
)

// Serobot is ingrained in our culture. Why queue when you
// can just jump the queue?
func (q *queue[T]) Serobot(value T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(q.Size - 1)
	chunk := q.Data[n:]
	q.Data = q.Data[:n]
	q.Data = append(q.Data, value)
	q.Data = append(q.Data, chunk...)
	q.Size = q.Size + 1
	q.Head = q.Data[0]
	q.Tail = q.Data[q.Size-1]
}

// OrDal is a more powerful version of Serobot, as it guarantees
// that you get in front of the queue no matter how long it is.
func (q *queue[T]) OrDal(value T) {
	getFkd := q.Data
	q.Data = nil
	q.Data = append(q.Data, value)
	q.Data = append(q.Data, getFkd...)
	q.Size = q.Size + 1
	q.Head = value
	q.Tail = q.Data[q.Size-1]
}
