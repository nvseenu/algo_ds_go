package sort

import (
	"algo_ds/util"
	"reflect"
	"testing"
)

func compare(a interface{}, b interface{}) int {
	v1, _ := a.(int)
	v2, _ := b.(int)
	return v1 - v2
}

var sorters = []struct {
	name   string
	sorter Sorter
}{
	{"QuickSort", NewSorter(QUICK, compare)},
	{"BubbleSort", NewSorter(BUBBLE, compare)},
	{"SelectionSort", NewSorter(SELECTION, compare)},
	{"InsertionSort", NewSorter(INSERTION, compare)},
}

var dataSet = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{2, 8, 7, 1, 3, 5, 6, 4},
		output: []int{1, 2, 3, 4, 5, 6, 7, 8},
	},

	{
		input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		output: []int{1, 2, 3, 4, 5, 6, 7, 8},
	},
	{
		input:  []int{8, 7, 6, 5, 4, 3, 2, 1},
		output: []int{1, 2, 3, 4, 5, 6, 7, 8},
	},
}

func TestSort(t *testing.T) {

	for _, sdata := range sorters {
		for _, data := range dataSet {

			input := util.GenericArray(data.input)
			output := util.GenericArray(data.output)
			// Method under test
			sdata.sorter.Sort(input)
			if !reflect.DeepEqual(output, input) {
				t.Error(sdata.name, " : ", data.input, " are not sorted! expected:", output, "But got:", input)
			}
		}
	}
}
