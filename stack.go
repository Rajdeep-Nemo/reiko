// Package reiko provides basic generic data structures.
package reiko

// Stack is a last-in, first-out (LIFO) collection of values of type T.
// The zero value is an empty stack ready to use.
// A Stack is not safe for concurrent use.
type Stack[T any] struct {
	items []T
}

// Push adds value to the top of the stack.
func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

// Pop removes and returns the top value of the stack.
// It returns the zero value and false if the stack is empty.
func (stack *Stack[T]) Pop() (T, bool) {
	var zeroValue T

	// Nothing to pop
	if len(stack.items) == 0 {
		return zeroValue, false
	}

	// Take the top element and shrink the stack
	value := stack.items[len(stack.items)-1]
	stack.items[len(stack.items)-1] = zeroValue // release the reference so it can be garbage collected
	stack.items = stack.items[:len(stack.items)-1]

	return value, true
}

// Peek returns the top value of the stack without removing it.
// It returns the zero value and false if the stack is empty.
func (stack *Stack[T]) Peek() (T, bool) {
	var zeroValue T

	// Nothing to peek at
	if len(stack.items) == 0 {
		return zeroValue, false
	}

	return stack.items[len(stack.items)-1], true
}

// Len returns the number of elements in the stack.
func (stack *Stack[T]) Len() int {
	return len(stack.items)
}

// IsEmpty reports whether the stack has no elements.
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

// Clear removes all elements from the stack.
func (stack *Stack[T]) Clear() {
	stack.items = nil
}
