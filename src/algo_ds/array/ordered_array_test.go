package array

import (
	"fmt"
	"testing"
)

type employee struct {
	empId int
	name  string
}

func compareEmployees(a interface{}, b interface{}) int {
	v1, ok := a.(*employee)
	if !ok {
		return 1
	}
	v2, ok := b.(*employee)
	if !ok {
		return -1
	}

	return v2.empId - v1.empId
}

func TestFind(t *testing.T) {
	arr := NewOrderedArray(10, compareEmployees)
	arr.Insert(&employee{5, "e5"})
	arr.Insert(&employee{4, "e4"})
	arr.Insert(&employee{3, "e3"})
	arr.Insert(&employee{2, "e2"})
	arr.Insert(&employee{1, "e1"})

	for i := 0; i < arr.totalElms; i++ {
		if emp, _ := arr.Find(&employee{i + 1, fmt.Sprintf("e%d", i+1)}); emp.(*employee).empId != i+1 {
			t.Error("Expected empId ", i+1, ", but got ", emp)
		}
	}

	if _, ok := arr.Find(&employee{6, "e6"}); ok {
		t.Error("Expected search result is false, but got ", ok)
	}
}

func TestInsert(t *testing.T) {
	arr := NewOrderedArray(10, compareEmployees)
	arr.Insert(&employee{5, "e5"})
	arr.Insert(&employee{4, "e4"})
	arr.Insert(&employee{3, "e3"})
	arr.Insert(&employee{2, "e2"})
	arr.Insert(&employee{1, "e1"})

	for i := 0; i < arr.totalElms; i++ {
		if emp, _ := arr.Find(&employee{i + 1, fmt.Sprintf("e%d", i+1)}); emp.(*employee).empId != i+1 {
			t.Error("Expected empId ", i+1, ", but got ", emp)
		}
	}

}

func TestDelete(t *testing.T) {
	arr := NewOrderedArray(10, compareEmployees)
	arr.Insert(&employee{5, "e5"})
	arr.Insert(&employee{4, "e4"})
	arr.Insert(&employee{3, "e3"})
	arr.Insert(&employee{2, "e2"})
	arr.Insert(&employee{1, "e1"})

	for i := 0; i < 5; i++ {
		key := i + 1
		emp := &employee{key, fmt.Sprintf("e%d", key)}
		arr.Delete(emp)
		if _, ok := arr.Find(emp); ok {
			t.Error("Key ", emp, " should not be exist since we deleted it, but it exists ")
		}
	}

	if arr.Size() != 0 {
		t.Error("After deleteing all elements, size should be 0, but got ", arr.Size())
	}

}

func TestGet(t *testing.T) {
	arr := NewOrderedArray(10, compareEmployees)
	arr.Insert(&employee{5, "e5"})
	arr.Insert(&employee{4, "e4"})
	arr.Insert(&employee{3, "e3"})
	arr.Insert(&employee{2, "e2"})
	arr.Insert(&employee{1, "e1"})

	for i := 0; i < 5; i++ {
		v, _ := arr.Get(i)
		emp, _ := v.(*employee)
		if emp.empId != (i + 1) {
			t.Error("Expected value ", i+1, " at index:", i, " , but got key", v)
		}
	}
}
