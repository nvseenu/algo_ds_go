package collections

type Stack interface {
	Push(a interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)

	IsEmpty() bool
	IsFull() bool

	Size() int
}

type stackData struct {
	arr   []interface{}
	count int
}

type errorMsg struct {
	message string
}

func (e *errorMsg) Error() string {
	return e.message
}

func (sd *stackData) Push(a interface{}) error {
	if sd.IsFull() {
		return &errorMsg{"Stack is full. we can not push!"}
	}
	sd.arr[sd.count] = a
	sd.count++
	return nil
}

func (sd *stackData) Pop() (interface{}, error) {
	if sd.IsEmpty() {
		return nil, &errorMsg{"Stack is empty. we can not pop!"}
	}
	a := sd.arr[sd.count-1]
	sd.count--
	return a, nil
}

func (sd *stackData) Peek() (interface{}, error) {
	if sd.IsEmpty() {
		return nil, &errorMsg{"Stack is empty. we can not peek!"}
	}
	return sd.arr[sd.count-1], nil
}

func (sd *stackData) IsEmpty() bool {
	return sd.count == 0
}

func (sd *stackData) IsFull() bool {
	return sd.count == len(sd.arr)
}

func (sd *stackData) Size() int {
	return sd.count
}

func NewStack() (Stack, bool) {
	arr := make([]interface{}, 16)
	var s interface{} = &stackData{arr, 0}
	stack, ok := s.(Stack)
	return stack, ok
}
