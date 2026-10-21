# reiko

Small, generic data structures for Go that the standard library doesn't provide.

> Currently available: **Stack**. Queue and Set are planned.

## Installation

```sh
go get github.com/Rajdeep-Nemo/reiko
```

Requires Go 1.18 or later (generics).

## Stack

A last-in, first-out (LIFO) collection. The last value pushed is the first one popped.

The zero value is an empty stack, ready to use. A `Stack` is **not** safe for concurrent use; wrap it with a `sync.Mutex` if you need that.

### Example

```go
package main

import (
	"fmt"

	"github.com/Rajdeep-Nemo/reiko"
)

func main() {
	var s reiko.Stack[int]

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if top, ok := s.Peek(); ok {
		fmt.Println("top:", top) // top: 3
	}

	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println(v) // 3, 2, 1
	}

	_, ok := s.Pop()
	fmt.Println(ok) // false, the stack is empty
}
```

### Methods

| Method             | Description                                                                                          |
| ------------------ | ---------------------------------------------------------------------------------------------------- |
| `Push(value T)`    | Adds `value` to the top of the stack.                                                                |
| `Pop() (T, bool)`  | Removes and returns the top value. Returns the zero value and `false` if the stack is empty.         |
| `Peek() (T, bool)` | Returns the top value without removing it. Returns the zero value and `false` if the stack is empty. |
| `Len() int`        | Returns the number of elements.                                                                      |
| `IsEmpty() bool`   | Reports whether the stack has no elements.                                                           |
| `Clear()`          | Removes all elements.                                                                                |

`Pop` and `Peek` return a `bool` instead of an error because an empty stack is an expected state, like a missing map key. Check the `ok` value before using the result.
