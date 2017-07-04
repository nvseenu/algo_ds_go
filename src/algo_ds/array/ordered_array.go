package array

type OrderedArray struct {
	array     []interface{}
	totalElms int
	compare   func(a interface{}, b interface{}) int
}

func NewOrderedArray(capacity int, compare func(a interface{}, b interface{}) int) *OrderedArray {
	return &OrderedArray{make([]interface{}, capacity), 0, compare}
}

func (a *OrderedArray) Find(key interface{}) (interface{}, bool) {
	_, value, ok := a.find(key)
	return value, ok
}

func (a *OrderedArray) find(key interface{}) (int, interface{}, bool) {
	startIndex := 0
	endIndex := a.totalElms - 1
	currentIndex := 0

	for true {

		currentIndex = (startIndex + endIndex) / 2

		res := a.compare(a.array[currentIndex], key)

		if res == 0 {
			return currentIndex, a.array[currentIndex], true
		} else if startIndex > endIndex {
			return currentIndex, nil, false
		} else if res > 0 {
			startIndex = currentIndex + 1
		} else if res < 0 {
			endIndex = currentIndex - 1
		}
	}
	return currentIndex, nil, false
}

func (a *OrderedArray) Insert(key interface{}) {
	i := 0
	//Find an index to insert a new key
	for ; i < a.totalElms; i++ {
		res := a.compare(a.array[i], key)

		if res < 0 {
			break
		}
	}
	// Make a space by shiting each element to its next index
	for j := a.totalElms - 1; j >= i; j-- {
		a.array[j+1] = a.array[j]
	}

	a.array[i] = key
	a.totalElms++
}

func (a *OrderedArray) Size() int {
	return a.totalElms
}

func (a *OrderedArray) Get(index int) (interface{}, bool) {
	if index < 0 || index >= a.totalElms {
		return -1, false
	}

	return a.array[index], true
}

func (a *OrderedArray) Delete(key interface{}) (interface{}, bool) {
	index, value, _ := a.find(key)

	// Shrink  a space by shiting each element to its previous index
	i := 0
	for i = index; i < a.totalElms-1; i++ {
		a.array[i] = a.array[i+1]
	}
	a.array[i] = nil
	a.totalElms--

	return value, true
}
