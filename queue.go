// Package reiko provides basic generic data structures.
package reiko

type Queue[T any] struct{
	items []T
}

// Enqueue
func (queue *Queue[T]) Enqueue(value T) {
	queue.items = append(queue.items, value)
}