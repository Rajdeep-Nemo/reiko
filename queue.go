// Package reiko provides basic generic data structures.
package reiko

// Queue is a first-in, first-out (FIFO) collection of values of type T.
// The zero value is an empty queue ready to use.
// A Queue is not safe for concurrent use.
type Queue[T any] struct {
	items []T
}

// Enqueue adds value to the back of the queue.
func (queue *Queue[T]) Enqueue(value T) {
	queue.items = append(queue.items, value)
}

// Dequeue removes and returns the value at the front of the queue.
// It returns the zero value and false if the queue is empty.
func (queue *Queue[T]) Dequeue() (T, bool) {
	var zeroValue T

	if len(queue.items) == 0 {
		return zeroValue, false
	}

	value := queue.items[0]
	queue.items[0] = zeroValue // release the reference so it can be garbage collected
	queue.items = queue.items[1:]

	return value, true
}

// Peek returns the value at the front of the queue without removing it.
// It returns the zero value and false if the queue is empty.
func (queue *Queue[T]) Peek() (T, bool) {
	var zeroValue T
	if len(queue.items) == 0 {
		return zeroValue, false
	}
	return queue.items[0], true
}

// Len returns the number of elements in the queue.
func (queue *Queue[T]) Len() int {
	return len(queue.items)
}

// IsEmpty reports whether the queue has no elements.
func (queue *Queue[T]) IsEmpty() bool {
	return len(queue.items) == 0
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.items = nil
}
