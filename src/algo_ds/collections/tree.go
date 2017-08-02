package collections

type Tree interface {
	Insert(key interface{}) error
	Find(key interface{}) interface{}
	Delete(key interface{}) (interface{}, error)
	Size() int
}
