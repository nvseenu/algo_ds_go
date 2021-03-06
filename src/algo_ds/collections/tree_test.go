package collections

import (
	"algo_ds/util"
	"fmt"
	"testing"
)

func executeWithTrees(f func(tree Tree, t *testing.T), t *testing.T) {
	for _, tree := range CreateTrees() {
		f(tree, t)
	}
}

func TestNewTree(t *testing.T) {
	executeWithTrees(testNewTree, t)
}

func testNewTree(tree Tree, t *testing.T) {
	if tree.Size() != 0 {
		t.Errorf("%v => Expected 0 for size , but got %v", tree, tree.Size())
	}
}

func TestInsertInTree(t *testing.T) {
	executeWithTrees(testInsertInTree, t)
}

func testInsertInTree(tree Tree, t *testing.T) {
	for i := 1; i <= 10; i++ {
		tree.Insert(i)

	}
	if tree.Size() != 10 {
		t.Errorf("%v => Expected size 10. but got %v", tree, tree.Size())
	}
}

func TestFindInTree(t *testing.T) {
	tr := NewBinarySearchTree(util.CompareInt)
	for i := 1; i <= 10; i++ {
		tr.Insert(i)
	}

	for i := 10; i > 0; i-- {
		val := tr.Find(i)
		v, _ := val.(int)
		if v != i {
			t.Errorf("Expected %v , but got %v", i, v)
		}
	}

	val := tr.Find(11)
	if val != nil {
		t.Errorf("Expected nil, but got %v", val)
	}

}

func CreateTrees() []Tree {
	return []Tree{
		NewBinarySearchTree(util.CompareInt),
	}
}

type deleteData struct {
	description string
	elements    []int
	keys        []int
}

func (dd *deleteData) containsKey(key int) bool {
	for _, k := range dd.keys {
		if k == key {
			return true
		}
	}
	return false
}

func (dd *deleteData) String() string {
	return fmt.Sprintf("%v\n elements:%v\n keys:%v\n", dd.description, dd.elements, dd.keys)
}

var testDataForDeletion = []deleteData{
	{"Delete a right leaf node",
		[]int{2, 1, 3},
		[]int{3},
	},
	{"Delete a left leaf node",
		[]int{2, 1, 3},
		[]int{1},
	},
	{"Delete a right node which has one child",
		[]int{2, 1, 3, 5, 0},
		[]int{3},
	},
	{"Delete a left node which has one child",
		[]int{2, 1, 3, 5, 0},
		[]int{1},
	},
	{"Delete a node whose successor is right node",
		[]int{50, 75, 62, 87, 93},
		[]int{75},
	},
	{"Delete a node whose successor is immediate left descendant of its right child",
		[]int{50, 75, 62, 87, 77, 93, 79},
		[]int{75},
	},
	{"Delete a node whose successor is left descendant of its right child",
		[]int{50, 75, 62, 87, 80, 93, 79},
		[]int{75},
	},
	{"Delete a root node which is a leaf",
		[]int{2},
		[]int{2},
	},
	{"Delete a root node which has a left child",
		[]int{2, 1},
		[]int{2},
	},
	{"Delete a root node which has a right child",
		[]int{2, 3},
		[]int{2},
	},
	{"Delete a root node whose successor is leaf node",
		[]int{2, 1, 3},
		[]int{2},
	},
	{"Delete a root node whose successor is right node",
		[]int{2, 1, 3, 4},
		[]int{2},
	},
	{"Delete a root node whose successor is left descendant of its right child",
		[]int{2, 1, 4, 3},
		[]int{2},
	},
	{"Delete root nodes repeatedly",
		[]int{50, 75, 62, 87, 80, 93, 79},
		[]int{50, 75, 79, 80, 87, 93, 62},
	},
}

func TestDeleteNodeInTree(t *testing.T) {
	for _, td := range testDataForDeletion {
		tr := NewBinarySearchTree(util.CompareInt)
		for _, elm := range td.elements {
			tr.Insert(elm)
		}

		for _, k := range td.keys {
			res, _ := tr.Delete(k)
			r, _ := res.(int)
			if r != k {
				t.Errorf("%v => Expected %v but got %v", td, k, r)
			}
		}

		expSize := len(td.elements) - len(td.keys)
		if tr.Size() != expSize {
			t.Errorf("%v => Expected size is %v but got %v", td, expSize, tr.Size())
		}

		//verify other elements are there in the tree
		for _, elm := range td.elements {
			// if key is deleted, it should not exist in the tree
			if td.containsKey(elm) {
				if tr.Find(elm) != nil {
					t.Errorf("%v => Expected %v should not be in the tree but it did ", td, elm)
				}
				continue
			}

			res := tr.Find(elm)
			if res == nil {
				t.Errorf("%v => Expected %v is in the tree, but it did not", td.description, elm)
			} else {
				r, _ := res.(int)
				if r != elm {
					t.Errorf("%v => Expected %v is in the tree, but got %v", td.description, elm, r)
				}
			}
		}

	}

}
