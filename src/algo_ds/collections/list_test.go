package collections

import "testing"

func TestNewList(t *testing.T) {
	ls := NewLinkedList()
	if ls.Size() != 0 {
		t.Error("Newly craeted list should have 0 elements, but got ", ls.Size())
	}
}

func TestListAdd(t *testing.T) {
	size := 10
	ls := NewLinkedList()

	for i := 1; i <= size; i++ {
		ls.Add(i)
	}

	if ls.Size() != size {
		t.Errorf("After inserting %v elements, size should be %v , but got %v", size, size, ls.Size())
	}
}

func TestListGet(t *testing.T) {
	size := 10
	ls := NewLinkedList()

	for i := 1; i <= size; i++ {
		ls.Add(i)
	}
	for i := 1; i <= ls.Size(); i++ {
		v, _ := ls.GetAt(i)
		val, _ := v.(int)
		if val != i {
			t.Errorf("Expected %v at index %v, but got %v", i, (i - 1), val)
		}
	}
}

func TestListRemove(t *testing.T) {
	size := 10
	ls := NewLinkedList()

	for i := 1; i <= size; i++ {
		ls.Add(i)
	}

	for i := ls.Size(); i >= 1; i-- {
		ls.Remove(i)
	}

	if ls.Size() != 0 {
		t.Errorf("After removing all inserted items, the size should be %v, but got %v", ls.Size(), ls.Size())
	}
}

func TestListRemoveAt(t *testing.T) {
	size := 10
	ls := NewLinkedList()

	for i := 1; i <= size; i++ {
		ls.AddAt(i-1, i)
	}

	for i := ls.Size() - 1; i >= 0; i-- {
		data, _ := ls.RemoveAt(i)
		dval, _ := data.(int)
		if dval != i+1 {
			t.Error("Expected removed elements is ", i+1, ", but got ", dval)
		}
	}
}

func TestListAddAndRemoveAtBeginning(t *testing.T) {
	size := 10
	ls := NewLinkedList()

	for i := 1; i <= size; i++ {
		ls.AddAt(0, i)
	}

	for i := ls.Size() - 1; i >= 0; i-- {
		data, _ := ls.RemoveAt(0)
		dval, _ := data.(int)
		if dval != i+1 {
			t.Error("Expected removed elements is ", i+1, ", but got ", dval)
		}
	}
}
