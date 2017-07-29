package collections

type link struct {
	data interface{}
	next *link
}

type singlyLinkedList struct {
	head *link
	size int
}

func (sll *singlyLinkedList) AddAt(index int, elm interface{}) error {
	pnode := sll.find(index - 1)
	nnode := &link{elm, nil}

	if pnode == nil {
		t := sll.head
		nnode.next = t
		sll.head = nnode
	} else {
		t := pnode.next
		nnode.next = t
		pnode.next = nnode
	}
	sll.size++
	return nil
}
func (sll *singlyLinkedList) GetAt(index int) (interface{}, error) {
	node := sll.find(index)
	if node == nil {
		return nil, nil

	} else {
		return node.data, nil
	}
}

func (sll *singlyLinkedList) find(index int) *link {
	if index < 0 {
		return nil
	}
	node := sll.head
	for i := 1; i <= index; i++ {
		node = node.next
	}
	return node
}

func (sll *singlyLinkedList) RemoveAt(index int) (interface{}, error) {
	pnode := sll.find(index - 1)
	var elm interface{} = nil
	if pnode == nil {
		cnode := sll.head
		sll.head = cnode.next
		elm = cnode.data
	} else {
		cnode := pnode.next
		pnode.next = cnode.next
		elm = cnode.data
	}
	sll.size--
	return elm, nil
}
func (sll *singlyLinkedList) Size() int {
	return sll.size
}

func NewLinkedList() List {
	var lld interface{} = &singlyLinkedList{
		head: nil,
		size: 0,
	}
	ll, _ := lld.(List)
	return ll
}
