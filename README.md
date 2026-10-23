# reiko

Small, generic data structures for Go that the standard library doesn't provide.

> Currently available: Stack and Queue, Set is planned.

## Installation

```sh
go get github.com/Rajdeep-Nemo/reiko
```

Requires Go 1.18 or later (generics).

## Stack

The zero value is an empty stack, ready to use. A `Stack` is **not** safe for concurrent use.

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

## Queue

The zero value is an empty queue, ready to use. A `Queue` is **not** safe for concurrent use.

### Methods

| Method                | Description                                                                                            |
| --------------------- | ------------------------------------------------------------------------------------------------------ |
| `Enqueue(value T)`    | Adds `value` to the back of the queue.                                                                 |
| `Dequeue() (T, bool)` | Removes and returns the front value. Returns the zero value and `false` if the queue is empty.         |
| `Peek() (T, bool)`    | Returns the front value without removing it. Returns the zero value and `false` if the queue is empty. |
| `Len() int`           | Returns the number of elements.                                                                        |
| `IsEmpty() bool`      | Reports whether the queue has no elements.                                                             |
| `Clear()`             | Removes all elements.                                                                                  |

`Dequeue` and `Peek` return a `bool` instead of an error because an empty queue is an expected state, like a missing map key. Check the `ok` value before using the result.
