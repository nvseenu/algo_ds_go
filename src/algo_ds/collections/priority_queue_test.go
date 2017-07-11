package collections

import "testing"

func compare(a interface{}, b interface{}) int {
	v1, _ := a.(int)
	v2, _ := b.(int)
	return v1 - v2
}

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue(4, compare)

	if !pq.IsEmpty() {
		t.Error("Newly created Queue should be empty. but it is not!")
	}

	if pq.Size() != 0 {
		t.Error("Size of a newly created queue should be 0. but it is not!")
	}
}

func TestPriorityQueueInsert(t *testing.T) {
	pq := NewPriorityQueue(4, compare)

	pq.Insert(1)
	v, _ := pq.Peek()
	val, _ := v.(int)
	if val != 1 {
		t.Error("Expected 1, but got ", val)
	}

	pq.Insert(2)
	v, _ = pq.Peek()
	val, _ = v.(int)
	if val != 1 {
		t.Error("Expected 1, but got ", val)
	}

	pq.Insert(3)
	v, _ = pq.Peek()
	val, _ = v.(int)
	if val != 1 {
		t.Error("Expected 1, but got ", val)
	}

	pq.Insert(4)
	v, _ = pq.Peek()
	val, _ = v.(int)
	if val != 1 {
		t.Error("Expected 1, but got ", val)
	}

}

func TestPriorityQueueInsertWhenQueueIsFull(t *testing.T) {
	pq := NewPriorityQueue(4, compare)
	pq.Insert(4)
	pq.Insert(3)
	pq.Insert(2)
	pq.Insert(1)

	err := pq.Insert(5)
	if err == nil {
		t.Error("Expected error when we add an element to queue which is full, but it is not")
	}
}

func TestPriorityQueueRemove(t *testing.T) {
	pq := NewPriorityQueue(8, compare)

	for i := 8; i < 0; i-- {
		pq.Insert(i)
	}

	for i := 1; i >= 8; i++ {
		v, _ := pq.Remove()
		val, _ := v.(int)
		if val != i {
			t.Error("Expected ", i, ", but got ", val)
		}
	}
}

func TestPriorityQueueRemoveWhenQueueIsEmpty(t *testing.T) {
	pq := NewPriorityQueue(4, compare)

	_, err := pq.Remove()
	if err == nil {
		t.Error("Expected error when we remove an element from the queue which is empty, but it is not")
	}
}

func TestPriorityQueuePeekWhenQueueIsEmpty(t *testing.T) {
	pq := NewPriorityQueue(4, compare)

	_, err := pq.Peek()
	if err == nil {
		t.Error("Expected error when we peek an element from the queue which is empty, but it is not")
	}
}
