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