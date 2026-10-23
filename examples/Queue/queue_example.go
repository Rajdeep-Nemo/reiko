package main

import (
	"fmt"

	"github.com/Rajdeep-Nemo/reiko"
)

func main() {
	var q reiko.Queue[int]

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	front, _ := q.Peek()
	fmt.Println("front:", front) // front: 1

	for !q.IsEmpty() {
		n, _ := q.Dequeue()
		fmt.Println(n) // 1, 2, 3
	}

	// Empty case: zero value and false
	n, ok := q.Dequeue()
	fmt.Println(n, ok) // 0 false
}
