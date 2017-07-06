package sort

type Sorter interface {
	Sort(arr []interface{})
}

type sorting struct {
	compare func(a interface{}, b interface{}) int
}
