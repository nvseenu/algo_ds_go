package sort

type InsertionSort sorting

func (is *InsertionSort) Sort(arr []interface{}) {
	for i := 1; i < len(arr); i++ {

		key := arr[i]
		j := i - 1

		for j >= 0 && is.compare(arr[j], key) > 0 {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
}
