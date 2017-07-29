package collections

import "testing"

func newLists() []List {
	return []List{
		NewLinkedList(),
		NewDoublyLinkedList(),
	}
}

func testLists(f func(List, *testing.T), t *testing.T) {
	for _, ls := range newLists() {
		f(ls, t)
	}
}

func TestNewList(t *testing.T) {
	testLists(testNewList, t)
}

func testNewList(ls List, t *testing.T) {
	if ls.Size() != 0 {
		t.Errorf("[%T] => Newly created list should have 0 elements, but got %v", ls, ls.Size())
	}
}

func TestListAdd(t *testing.T) {
	testLists(testListAdd, t)
}

func testListAdd(ls List, t *testing.T) {
	size := 10

	for i := 1; i <= size; i++ {
		ls.AddAt(0, i)
	}

	if ls.Size() != size {
		t.Errorf("[%T] => After inserting %v elements, size should be %v , but got %v", ls, size, size, ls.Size())
	}
}

func TestListGet(t *testing.T) {
	testLists(testListGet, t)
}

func testListGet(ls List, t *testing.T) {
	size := 10

	for i := 1; i <= size; i++ {
		ls.AddAt(0, i)
	}
	j := 10
	for i := 0; i < ls.Size(); i++ {
		v, _ := ls.GetAt(i)
		val, _ := v.(int)
		if val != j {
			t.Errorf("[%T] => Expected %v at index %v, but got %v", ls, j, i, val)
		}
		j--
	}
}

func TestListRemoveAt(t *testing.T) {
	testLists(testListRemoveAt, t)
}

func testListRemoveAt(ls List, t *testing.T) {
	size := 10
	for i := 1; i <= size; i++ {
		ls.AddAt(0, i)
	}

	for i := ls.Size(); i >= 1; i-- {
		data, _ := ls.RemoveAt(0)
		dval, _ := data.(int)
		if dval != i {
			t.Errorf("[%T] => Expected removed element  is %v at index %v, but got %v ", ls, i, 0, dval)
		}
	}
}
