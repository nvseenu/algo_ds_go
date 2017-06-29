package sort

import (
	"algo_ds/util"
	"reflect"
	"testing"
)

var sorters = []Sorter{
	QuickSort{},
}

func TestSort(t *testing.T) {

	for _, s := range sorters {
		arr := util.GenericArray([]int{2, 8, 7, 1, 3, 5, 6, 4})
		s.Sort(arr)
		expArr := util.GenericArray([]int{1, 2, 3, 4, 5, 6, 7, 8})

		if !reflect.DeepEqual(expArr, arr) {
			t.Error("Arrays are not sorted! expected:", expArr, "But got:", arr)
		}
	}
}
