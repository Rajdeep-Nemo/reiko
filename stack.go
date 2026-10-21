package whatever

// Stack struct
type Stack[T any] struct {
	items []T
}

// Push a value onto the stack
func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

// Pop the top value from the stack
func (stack *Stack[T]) Pop() (T, bool) {
	// Default value of type T
	var zeroValue T
	// If empty return zero and false
	if len(stack.items) == 0 {
		return zeroValue, false
	}
	// Get the top element and remove it from the stack
	value := stack.items[len(stack.items)-1]
	stack.items[len(stack.items)-1] = zeroValue
	stack.items = stack.items[:len(stack.items)-1]
	// Return the top element
	return value, true
}

// Peek returns the top element without removing from the stack
func (stack *Stack[T]) Peek() (T, bool) {
	// Default value of type T
	var zeroValue T
	// If empty return error
	if len(stack.items) == 0 {
		return zeroValue, false
	}
	// Get the top element
	value := stack.items[len(stack.items)-1]
	// Return the top element
	return value, true
}

// Len returns the current size of the stack
func (stack *Stack[T]) Len() int {
	return len(stack.items)
}

// IsEmpty returns if the stack is empty
func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

// Clear removes every element from the stack
func (stack *Stack[T]) Clear() {
	stack.items = nil
}
