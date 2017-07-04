package array

type OrderedArray struct {
	array     []interface{}
	totalElms int
}

func NewOrderedArray(capacity int) *OrderedArray {
	return &OrderedArray{make([]interface{}, capacity), 0}
}

func (a *OrderedArray) Find(key int) (int, bool) {
	startIndex := 0
	endIndex := a.totalElms - 1
	currentIndex := 0

	for true {

		currentIndex = (startIndex + endIndex) / 2

		v, _ := a.array[currentIndex].(int)

		if v == key {
			return currentIndex, true
		} else if startIndex > endIndex {
			return currentIndex, false
		} else if v < key {
			startIndex = currentIndex + 1
		} else if v > key {
			endIndex = currentIndex - 1
		}
	}
	return currentIndex, false
}

func (a *OrderedArray) Insert(key int) {
	i := 0
	//Find an index to insert a new key
	for ; i < a.totalElms; i++ {
		v, _ := a.array[i].(int)

		if v > key {
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

func (a *OrderedArray) Get(index int) (int, bool) {
	if index < 0 || index >= a.totalElms {
		return -1, false
	}

	return a.array[index].(int), true
}

func (a *OrderedArray) Delete(key int) (int, bool) {
	index, _ := a.Find(key)

	// Shrink  a space by shiting each element to its previous index
	i := 0
	for i = index; i < a.totalElms-1; i++ {
		a.array[i] = a.array[i+1]
	}
	a.array[i] = nil
	a.totalElms--

	return key, true
}
