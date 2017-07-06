package sort

import "algo_ds/util"

type BubbleSort sorting

func (bs *BubbleSort) Sort(arr []interface{}) {

	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if bs.compare(arr[j], arr[j+1]) > 0 {
				util.Swap(arr, j, j+1)
			}
		}
	}
}
