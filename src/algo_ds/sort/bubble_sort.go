package sort

import "algo_ds/util"

type BubbleSort struct {
	compare func(a interface{}, b interface{}) int
}

func (b BubbleSort) Sort(arr []interface{}) {

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if b.compare(arr[j], arr[j+1]) > 0 {
				util.Swap(arr, j, j+1)
			}
		}
	}
}

func NewBubbleSort(compare func(a interface{}, b interface{}) int) *BubbleSort {
	return &BubbleSort{compare}
}
