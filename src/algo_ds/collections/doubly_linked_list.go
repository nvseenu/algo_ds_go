package collections

type node struct {
	data     interface{}
	previous *node
	next     *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (dll *doublyLinkedList) Add(elm interface{}) {
	dll.AddAt(0, elm)
}

func (dll *doublyLinkedList) AddAt(index int, elm interface{}) error {
	nnode := &node{elm, nil, nil}
	pnode := dll.find(index - 1)

	if pnode == nil {
		t := dll.head
		dll.head = nnode
		nnode.next = t
		nnode.previous = nil
	} else {
		t := pnode.next
		pnode.next = nnode
		nnode.previous = pnode
		nnode.next = t
	}
	dll.size++
	return nil
}

func (dll *doublyLinkedList) GetAt(index int) (interface{}, error) {
	cnode := dll.find(index)
	if cnode == nil {
		return nil, nil
	} else {
		return cnode.data, nil
	}
}

func (dll *doublyLinkedList) find(index int) *node {
	if index < 0 {
		return nil
	}
	cnode := dll.head
	for i := 1; i <= index; i++ {
		cnode = cnode.next
	}
	return cnode
}

func (dll *doublyLinkedList) RemoveAt(index int) (interface{}, error) {
	pnode := dll.find(index - 1)
	var elm interface{} = nil

	if pnode == nil {
		cnode := dll.head
		nnode := cnode.next
		dll.head = nnode
		// When there is only one node, then there will be no next node.
		if nnode != nil {
			nnode.previous = nil
		}
		elm = cnode.data
	} else {
		cnode := pnode.next
		nnode := cnode.next
		pnode.next = nnode
		nnode.previous = pnode
		elm = cnode.data
	}
	dll.size--
	return elm, nil
}

func (dll *doublyLinkedList) Size() int {
	return dll.size
}

func NewDoublyLinkedList() List {
	var dll interface{} = &doublyLinkedList{nil, nil, 0}
	ls, _ := dll.(List)
	return ls
}
