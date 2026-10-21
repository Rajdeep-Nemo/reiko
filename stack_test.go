package whatever

import "testing"

func TestStackZeroValue(t *testing.T) {
	var s Stack[int]

	if !s.IsEmpty() {
		t.Error("zero value stack should be empty")
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}

	s.Push(1)
	if v, ok := s.Pop(); !ok || v != 1 {
		t.Errorf("Pop() = (%d, %t), want (1, true)", v, ok)
	}
}

func TestStackPushPopOrder(t *testing.T) {
	var s Stack[int]
	for i := 1; i <= 3; i++ {
		s.Push(i)
	}

	// Last in, first out
	for _, want := range []int{3, 2, 1} {
		got, ok := s.Pop()
		if !ok {
			t.Fatalf("Pop() ok = false, want true (expected %d)", want)
		}
		if got != want {
			t.Errorf("Pop() = %d, want %d", got, want)
		}
	}

	if !s.IsEmpty() {
		t.Error("stack should be empty after popping everything")
	}
}

func TestStackPopEmpty(t *testing.T) {
	var s Stack[string]

	v, ok := s.Pop()
	if ok {
		t.Error("Pop() on empty stack: ok = true, want false")
	}
	if v != "" {
		t.Errorf("Pop() on empty stack returned %q, want zero value", v)
	}

	// Popping an emptied stack must behave the same way
	s.Push("a")
	s.Pop()
	if _, ok := s.Pop(); ok {
		t.Error("Pop() after emptying the stack: ok = true, want false")
	}
}

func TestStackPeek(t *testing.T) {
	var s Stack[int]
	s.Push(10)
	s.Push(20)

	got, ok := s.Peek()
	if !ok || got != 20 {
		t.Errorf("Peek() = (%d, %t), want (20, true)", got, ok)
	}

	// Peek must not remove the element
	if s.Len() != 2 {
		t.Errorf("Len() after Peek = %d, want 2", s.Len())
	}
	if got, _ := s.Peek(); got != 20 {
		t.Errorf("second Peek() = %d, want 20", got)
	}
}

func TestStackPeekEmpty(t *testing.T) {
	var s Stack[int]

	v, ok := s.Peek()
	if ok {
		t.Error("Peek() on empty stack: ok = true, want false")
	}
	if v != 0 {
		t.Errorf("Peek() on empty stack returned %d, want zero value", v)
	}
}

func TestStackLenAndIsEmpty(t *testing.T) {
	var s Stack[int]

	if s.Len() != 0 || !s.IsEmpty() {
		t.Fatalf("new stack: Len() = %d, IsEmpty() = %t", s.Len(), s.IsEmpty())
	}

	s.Push(1)
	s.Push(2)
	if s.Len() != 2 || s.IsEmpty() {
		t.Errorf("after 2 pushes: Len() = %d, IsEmpty() = %t", s.Len(), s.IsEmpty())
	}

	s.Pop()
	if s.Len() != 1 || s.IsEmpty() {
		t.Errorf("after 1 pop: Len() = %d, IsEmpty() = %t", s.Len(), s.IsEmpty())
	}

	s.Pop()
	if s.Len() != 0 || !s.IsEmpty() {
		t.Errorf("after 2 pops: Len() = %d, IsEmpty() = %t", s.Len(), s.IsEmpty())
	}
}

func TestStackClear(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if !s.IsEmpty() || s.Len() != 0 {
		t.Errorf("after Clear: Len() = %d, IsEmpty() = %t", s.Len(), s.IsEmpty())
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop() after Clear: ok = true, want false")
	}

	// Stack must still be usable after Clear
	s.Push(42)
	if v, ok := s.Pop(); !ok || v != 42 {
		t.Errorf("Pop() after reuse = (%d, %t), want (42, true)", v, ok)
	}

	// Clearing an already empty stack is a no-op
	s.Clear()
	s.Clear()
}

func TestStackWithStructAndPointerTypes(t *testing.T) {
	type point struct{ X, Y int }

	var ps Stack[point]
	ps.Push(point{1, 2})
	if got, ok := ps.Pop(); !ok || got != (point{1, 2}) {
		t.Errorf("Pop() = (%v, %t), want ({1 2}, true)", got, ok)
	}

	var pp Stack[*point]
	p := &point{3, 4}
	pp.Push(p)
	if got, ok := pp.Pop(); !ok || got != p {
		t.Errorf("Pop() = (%p, %t), want (%p, true)", got, ok, p)
	}
}

func TestStackPopReleasesReference(t *testing.T) {
	var s Stack[*int]
	n := 7
	s.Push(&n)
	s.Pop()

	// The popped slot in the backing array must be zeroed so the
	// garbage collector can reclaim the value.
	if backing := s.items[:1]; backing[0] != nil {
		t.Error("Pop() left a reference in the backing array")
	}
}

func BenchmarkStackPushPop(b *testing.B) {
	var s Stack[int]
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}