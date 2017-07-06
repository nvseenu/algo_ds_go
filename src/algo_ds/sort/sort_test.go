package sort

import (
	"algo_ds/util"
	"fmt"
	"reflect"
	"testing"
)

func compare(a interface{}, b interface{}) int {
	v1, _ := a.(int)
	v2, _ := b.(int)
	return v1 - v2
}

var sorters = []Sorter{
	NewQuickSort(compare), NewBubbleSort(compare),
}

func TestSort(t *testing.T) {

	for _, s := range sorters {
		arr := util.GenericArray([]int{2, 8, 7, 1, 3, 5, 6, 4})
		s.Sort(arr)
		expArr := util.GenericArray([]int{1, 2, 3, 4, 5, 6, 7, 8})

		if !reflect.DeepEqual(expArr, arr) {
			t.Error(fmt.Sprintf("%v", s), " : Arrays are not sorted! expected:", expArr, "But got:", arr)
		}
	}
}
