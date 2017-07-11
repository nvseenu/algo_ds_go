package collections

type pqData struct {
	arr     []interface{}
	front   int
	rear    int
	compare func(a interface{}, b interface{}) int
	count   int
}

type pqErr struct {
	message string
}

func (e *pqErr) Error() string {
	return e.message
}

func (pq *pqData) Insert(elm interface{}) error {

	if pq.IsFull() {
		return &pqErr{"We can not insert an element to the queue which is full"}
	}

	i := 0
	for i = pq.front - 1; i >= pq.rear; i-- {
		if pq.compare(elm, pq.arr[i]) <= 0 {
			break
		}
	}

	// move the items from i to front one cell right
	for j := pq.front - 1; j > i; j-- {

		pq.arr[j+1] = pq.arr[j]
	}
	pq.arr[i+1] = elm
	pq.front++
	pq.count++

	return nil
}

func (pq *pqData) Remove() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, &pqErr{"We can not remove an element from the queue which is empty"}
	}

	elm := pq.arr[pq.front-1]
	pq.front--
	pq.count--
	return elm, nil
}

func (pq *pqData) Peek() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, &pqErr{"We can not peek an element from the queue which is empty"}
	}
	return pq.arr[pq.front-1], nil
}
func (pq *pqData) IsFull() bool {
	return pq.Size() == len(pq.arr)
}
func (pq *pqData) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *pqData) Size() int {
	return pq.count
}

func NewPriorityQueue(size int, compare func(a interface{}, b interface{}) int) Queue {
	var pq interface{} = &pqData{arr: make([]interface{}, size),
		front:   0,
		rear:    0,
		count:   0,
		compare: compare,
	}
	q, _ := pq.(Queue)
	return q
}
