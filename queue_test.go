package main

import (
	"testing"
)

func BenchmarkQueueStr(t *testing.B) {
	queue := NewQueue()

	val := []string{"Risky", "Feryansyah", "Pribadi", "Shania", "Saraswati", "Khairin", "Utari", "Anarawati"}

	for i := 0; i < t.N; i++ {
		for n := 0; n < len(val); n++ {
			queue.Enqueue(val[n])
		}
	}
}
