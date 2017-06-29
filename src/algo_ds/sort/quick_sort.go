package sort

import "algo_ds/util"

type QuickSort struct {
}

func (q QuickSort) Sort(arr []interface{}) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []interface{}, p int, r int) {
	if p < r {
		q := partition(arr, p, r)
		quickSort(arr, p, q-1)
		quickSort(arr, q+1, r)
	}
}

func partition(arr []interface{}, p int, r int) int {
	x := util.IntType(arr[r])
	i := p - 1

	for j := p; j < r; j++ {

		if util.IntType(arr[j]) <= x {
			i++
			util.Swap(arr, i, j)
		}
	}
	util.Swap(arr, i+1, r)

	return i + 1
}
