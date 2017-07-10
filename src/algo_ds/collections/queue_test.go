package collections

import "testing"

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue(8)
	if !q.IsEmpty() {
		t.Error("Newly created queue shoudl be empty, but it is not!")
	}
}

func TestQueueIsFull(t *testing.T) {
	q := NewQueue(4)
	for i := 0; i < 4; i++ {
		q.Insert(i)
	}

	if !q.IsFull() {
		t.Error("After we added all elementes which are equal to its capacity, queue should be full, but it is not!")
	}
}

func TestQueueInsert(t *testing.T) {
	q := NewQueue(4)

	for i := 1; i < 5; i++ {
		q.Insert(i)
		v, _ := q.Peek()
		val, _ := v.(int)

		if val != 1 {
			t.Error("Expected value recently inserted is ", i, " but got ", val)
		}
	}
}

func TestQueueInsertWhenFull(t *testing.T) {
	q := NewQueue(4)

	for i := 1; i < 5; i++ {
		q.Insert(i)
	}

	err := q.Insert(5)
	if err == nil {
		t.Error("Error should be thrown when we try to insert an item in a queue which is full, but it is not")
	}
}

func TestQueueRemove(t *testing.T) {
	q := NewQueue(4)

	for i := 1; i < 5; i++ {
		q.Insert(i)
	}

	for i := 1; i < 5; i++ {
		v, _ := q.Remove()
		val, _ := v.(int)

		if val != i {
			t.Error("Expected value recently inserted is ", i, " but got ", val)
		}
	}
}

func TestQueueRemoveWhenEmpty(t *testing.T) {
	q := NewQueue(4)
	_, err := q.Remove()

	if err == nil {
		t.Error("Error should be thrown while trying to remove an item from queue. But it is not ")
	}
}

func TestQueueCircle(t *testing.T) {
	q := NewQueue(4)

	q.Insert(1)
	q.Insert(2)
	q.Insert(3)
	q.Insert(4)

	q.Remove()
	q.Remove()

	q.Insert(5)
	q.Insert(6)

	q.Remove()
	q.Remove()

	q.Insert(7)
	q.Insert(8)

	if q.Size() != 4 {
		t.Error("Expected Size is 4, but got ", q.Size())
	}
}
