package sort

import "algo_ds/util"

type SelectionSort sorting

func (ss *SelectionSort) Sort(arr []interface{}) {

	endIndex := len(arr) - 1
	for i := 0; i < endIndex; i++ {
		minIndex := ss.findMinIndex(arr, i, endIndex)
		util.Swap(arr, i, minIndex)
	}
}

func (ss *SelectionSort) findMinIndex(arr []interface{}, startIndex int, endIndex int) int {
	minIndex := startIndex
	for i := startIndex + 1; i <= endIndex; i++ {
		if ss.compare(arr[minIndex], arr[i]) >= 0 {
			minIndex = i
		}

	}
	return minIndex
}
