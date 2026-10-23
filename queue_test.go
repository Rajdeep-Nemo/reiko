package reiko

import "testing"

func TestQueueZeroValue(t *testing.T) {
	var q Queue[int]

	if !q.IsEmpty() {
		t.Error("zero value queue should be empty")
	}
	if q.Len() != 0 {
		t.Errorf("Len() = %d, want 0", q.Len())
	}

	q.Enqueue(1)
	if v, ok := q.Dequeue(); !ok || v != 1 {
		t.Errorf("Dequeue() = (%d, %t), want (1, true)", v, ok)
	}
}

func TestQueueEnqueueDequeueOrder(t *testing.T) {
	var q Queue[int]
	for i := 1; i <= 3; i++ {
		q.Enqueue(i)
	}

	// First in, first out
	for _, want := range []int{1, 2, 3} {
		got, ok := q.Dequeue()
		if !ok {
			t.Fatalf("Dequeue() ok = false, want true (expected %d)", want)
		}
		if got != want {
			t.Errorf("Dequeue() = %d, want %d", got, want)
		}
	}

	if !q.IsEmpty() {
		t.Error("queue should be empty after dequeuing everything")
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	var q Queue[string]

	v, ok := q.Dequeue()
	if ok {
		t.Error("Dequeue() on empty queue: ok = true, want false")
	}
	if v != "" {
		t.Errorf("Dequeue() on empty queue returned %q, want zero value", v)
	}

	// Dequeuing an emptied queue must behave the same way
	q.Enqueue("a")
	q.Dequeue()
	if _, ok := q.Dequeue(); ok {
		t.Error("Dequeue() after emptying the queue: ok = true, want false")
	}
}

func TestQueueInterleavedOperations(t *testing.T) {
	var q Queue[int]

	q.Enqueue(1)
	q.Enqueue(2)
	if v, _ := q.Dequeue(); v != 1 {
		t.Errorf("Dequeue() = %d, want 1", v)
	}

	q.Enqueue(3)
	if v, _ := q.Dequeue(); v != 2 {
		t.Errorf("Dequeue() = %d, want 2", v)
	}
	if v, _ := q.Dequeue(); v != 3 {
		t.Errorf("Dequeue() = %d, want 3", v)
	}
	if !q.IsEmpty() {
		t.Error("queue should be empty")
	}
}

func TestQueuePeek(t *testing.T) {
	var q Queue[int]
	q.Enqueue(10)
	q.Enqueue(20)

	got, ok := q.Peek()
	if !ok || got != 10 {
		t.Errorf("Peek() = (%d, %t), want (10, true)", got, ok)
	}

	// Peek must not remove the element
	if q.Len() != 2 {
		t.Errorf("Len() after Peek = %d, want 2", q.Len())
	}
	if got, _ := q.Peek(); got != 10 {
		t.Errorf("second Peek() = %d, want 10", got)
	}
}

func TestQueuePeekEmpty(t *testing.T) {
	var q Queue[int]

	v, ok := q.Peek()
	if ok {
		t.Error("Peek() on empty queue: ok = true, want false")
	}
	if v != 0 {
		t.Errorf("Peek() on empty queue returned %d, want zero value", v)
	}
}

func TestQueueLenAndIsEmpty(t *testing.T) {
	var q Queue[int]

	if q.Len() != 0 || !q.IsEmpty() {
		t.Fatalf("new queue: Len() = %d, IsEmpty() = %t", q.Len(), q.IsEmpty())
	}

	q.Enqueue(1)
	q.Enqueue(2)
	if q.Len() != 2 || q.IsEmpty() {
		t.Errorf("after 2 enqueues: Len() = %d, IsEmpty() = %t", q.Len(), q.IsEmpty())
	}

	q.Dequeue()
	if q.Len() != 1 || q.IsEmpty() {
		t.Errorf("after 1 dequeue: Len() = %d, IsEmpty() = %t", q.Len(), q.IsEmpty())
	}

	q.Dequeue()
	if q.Len() != 0 || !q.IsEmpty() {
		t.Errorf("after 2 dequeues: Len() = %d, IsEmpty() = %t", q.Len(), q.IsEmpty())
	}
}

func TestQueueClear(t *testing.T) {
	var q Queue[int]
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if !q.IsEmpty() || q.Len() != 0 {
		t.Errorf("after Clear: Len() = %d, IsEmpty() = %t", q.Len(), q.IsEmpty())
	}
	if _, ok := q.Dequeue(); ok {
		t.Error("Dequeue() after Clear: ok = true, want false")
	}

	// Queue must still be usable after Clear
	q.Enqueue(42)
	if v, ok := q.Dequeue(); !ok || v != 42 {
		t.Errorf("Dequeue() after reuse = (%d, %t), want (42, true)", v, ok)
	}

	// Clearing an already empty queue is a no-op
	q.Clear()
	q.Clear()
}

func TestQueueWithStructAndPointerTypes(t *testing.T) {
	type point struct{ X, Y int }

	var qs Queue[point]
	qs.Enqueue(point{1, 2})
	if got, ok := qs.Dequeue(); !ok || got != (point{1, 2}) {
		t.Errorf("Dequeue() = (%v, %t), want ({1 2}, true)", got, ok)
	}

	var qp Queue[*point]
	p := &point{3, 4}
	qp.Enqueue(p)
	if got, ok := qp.Dequeue(); !ok || got != p {
		t.Errorf("Dequeue() = (%p, %t), want (%p, true)", got, ok, p)
	}
}

func TestQueueDequeueReleasesReference(t *testing.T) {
	var q Queue[*int]
	n := 7
	q.Enqueue(&n)

	// Keep a view of the original backing array before dequeuing,
	// because Dequeue re-slices items past the first element.
	backing := q.items[:1]
	q.Dequeue()

	// The dequeued slot must be zeroed so the garbage collector
	// can reclaim the value.
	if backing[0] != nil {
		t.Error("Dequeue() left a reference in the backing array")
	}
}

func BenchmarkQueueEnqueueDequeue(b *testing.B) {
	var q Queue[int]
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}