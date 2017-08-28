package collections

import (
	"algo_ds/util"
	"testing"
)

func TestNew234Tree(t *testing.T) {
	tree := create234Tree()
	if tree.Size() != 0 {
		t.Errorf("%v => Expected 0 for size , but got %v", tree, tree.Size())
	}
}

type insertTestData struct {
	desc string
	data []int
}

var insertTestDataSet = []insertTestData{
	{"Root split", []int{1, 2, 3, 4}},
	{"Node level split", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
}

func TestInsertIn234Tree(t *testing.T) {
	for _, td := range insertTestDataSet {
		tree := create234Tree()

		for _, elm := range td.data {
			tree.Insert(elm)
		}

		if tree.Size() != len(td.data) {
			t.Errorf("%v => Expected size %v , but got %v", td.desc, len(td.data), tree.Size())
		}

		for _, elm := range td.data {
			if tree.Find(elm) == nil {
				t.Errorf("%v => Expected key %v must exist in the tree , but it is not ", td.desc, elm)
			}
		}

	}
}

func create234Tree() Tree {
	return New234Tree(util.CompareInt)
}
