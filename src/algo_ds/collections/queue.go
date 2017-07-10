package collections

type Queue interface {
	Insert(elm interface{}) error
	Remove() (interface{}, error)
	Peek() (interface{}, error)
	IsFull() bool
	IsEmpty() bool
	Size() int
}

type queueData struct {
	arr   []interface{}
	count int
	front int
	rear  int
}

type queueError struct {
	message string
}

func (err *queueError) Error() string {
	return err.message
}

func (qd *queueData) Insert(elm interface{}) error {

	if qd.IsFull() {
		return &queueError{"Queue is full. Hence we can not insert!!!"}
	}
	if qd.front == len(qd.arr) {
		qd.front = 0
	}

	qd.arr[qd.front] = elm
	qd.front++
	qd.count++
	return nil
}

func (qd *queueData) Remove() (interface{}, error) {

	if qd.IsEmpty() {
		return nil, &queueError{"Queue is empty. Hence we can not remove!!!"}
	}

	if qd.rear == len(qd.arr) {
		qd.rear = 0
	}

	elm := qd.arr[qd.rear]
	qd.rear++
	qd.count--
	return elm, nil
}
func (qd *queueData) Peek() (interface{}, error) {
	return qd.arr[qd.rear], nil
}
func (qd *queueData) IsFull() bool {
	return qd.count == len(qd.arr)
}
func (qd *queueData) IsEmpty() bool {
	return qd.count == 0
}
func (qd *queueData) Size() int {
	return qd.count
}

func NewQueue(size int) Queue {
	var q interface{} = &queueData{arr: make([]interface{}, size),
		count: 0,
		front: 0,
		rear:  0,
	}
	queue, _ := q.(Queue)
	return queue
}
