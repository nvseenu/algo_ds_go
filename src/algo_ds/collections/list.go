package collections

type List interface {
	AddAt(index int, elm interface{}) error
	GetAt(index int) (interface{}, error)
	RemoveAt(index int) (interface{}, error)
	Size() int
}
