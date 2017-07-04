package array

import (
	"testing"
)

func TestInsert(t *testing.T) {
	arr := NewOrderedArray(10)
	arr.Insert(5)
	arr.Insert(4)
	arr.Insert(3)
	arr.Insert(2)
	arr.Insert(1)

	for i := 0; i < arr.totalElms; i++ {
		if index, _ := arr.Find(i + 1); index != i {
			t.Error("Expected index ", i, " for key ", i+1, ", but got ", index)
		}
	}

}

func TestDelete(t *testing.T) {
	arr := NewOrderedArray(10)
	arr.Insert(5)
	arr.Insert(4)
	arr.Insert(3)
	arr.Insert(2)
	arr.Insert(1)

	for i := 0; i < 5; i++ {
		key := i + 1
		arr.Delete(key)
		if _, ok := arr.Find(key); ok {
			t.Error("Key ", key, " should not be exist since we deleted it, but it exists ")
		}
	}

	if arr.Size() != 0 {
		t.Error("After deleteing all elements, size should be 0, but got ", arr.Size())
	}

}

func TestGet(t *testing.T) {
	arr := NewOrderedArray(10)
	arr.Insert(5)
	arr.Insert(4)
	arr.Insert(3)
	arr.Insert(2)
	arr.Insert(1)

	for i := 0; i < 5; i++ {
		if v, _ := arr.Get(i); v != (i + 1) {
			t.Error("Expected value ", i+1, " at index:", i, " , but got key", v)
		}
	}
}
