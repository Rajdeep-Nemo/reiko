// Package reiko provides basic generic data structures.
package reiko

type Queue[T any] struct{
	items []T
}

// Enqueue
func (queue *Queue[T]) Enqueue(value T) {
	queue.items = append(queue.items, value)
}

// Dequeue
func (queue *Queue[T]) Dequeue() (T, bool) {
	var zeroValue T

	if len(queue.items) == 0 {
		return zeroValue, false
	}

	value := queue.items[0]
	queue.items[0] = zeroValue
	queue.items = queue.items[1:]

	return value, true
}