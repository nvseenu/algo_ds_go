package collections

type List interface {
	Add(elm interface{})
	AddAt(index int, elm interface{}) error
	GetAt(index int) (interface{}, error)
	Remove(elm interface{}) error
	RemoveAt(index int) (interface{}, error)
	Size() int
}
