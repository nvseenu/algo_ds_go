package util

func Swap(arr []interface{}, i int, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}
