package whatever

import (
	"errors"
)

// Stack type interface, holds data in a slice
type Stack[T any] struct {
	items []T
}

// Push a value onto the stack
func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

// Pop the top value from the stack
func (stack *Stack[T]) Pop() (T, error) {
	// Default value of type T
	var zeroValue T
	// If empty return error
	if len(stack.items) == 0 {
		return zeroValue, errors.New("stack is empty")
	}
	// Get the top element and remove it from the stack
	value := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	// Return the top element
	return value, nil
}

// Returns the top element without removing from the stack
func (stack *Stack[T]) Peek() (T, error) {
	// Default value of type T
	var zeroValue T
	// If empty return error
	if len(stack.items) == 0 {
		return zeroValue, errors.New("stack is empty")
	}
	// Get the top element
	value := stack.items[len(stack.items)-1]
	// Return the top element
	return value, nil
}
