package sort

import "algo_ds/util"

type QuickSort sorting

func (qs *QuickSort) Sort(arr []interface{}) {
	qs.quickSort(arr, 0, len(arr)-1)
}

func (qs *QuickSort) quickSort(arr []interface{}, p int, r int) {
	if p < r {
		q := qs.partition(arr, p, r)
		qs.quickSort(arr, p, q-1)
		qs.quickSort(arr, q+1, r)
	}
}

func (qs *QuickSort) partition(arr []interface{}, p int, r int) int {
	x := arr[r]
	i := p - 1

	for j := p; j < r; j++ {

		if qs.compare(arr[j], x) <= 0 {
			i++
			util.Swap(arr, i, j)
		}
	}
	util.Swap(arr, i+1, r)

	return i + 1
}
