package collections

import "reflect"

type link struct {
	data interface{}
	next *link
}

type lldata struct {
	head *link
	size int
}

func (ll *lldata) Add(elm interface{}) {
	l := &link{elm, nil}

	if ll.head == nil {
		ll.head = l
	} else {
		lk := ll.head
		for lk.next != nil {
			lk = lk.next
		}
		lk.next = l
	}
	ll.size++
}

func (ll *lldata) AddAt(index int, elm interface{}) error {
	pl := ll.head
	entry := &link{elm, nil}

	//Find a previous link to given index
	for i := 1; i < index; i++ {
		pl = pl.next
	}

	if ll.size == 0 {
		ll.head = entry
	} else {
		nlink := pl.next
		pl.next = entry
		entry.next = nlink
	}
	ll.size++
	return nil
}
func (ll *lldata) GetAt(index int) (interface{}, error) {
	l := ll.head
	for i := 1; i < index; i++ {
		l = l.next
	}
	return l.data, nil
}

func (ll *lldata) Remove(elm interface{}) error {

	l := ll.head
	var pl *link = nil
	for l != nil {
		if reflect.DeepEqual(l.data, elm) {
			if ll.Size() == 1 {
				ll.head = nil
			} else {
				pl.next = l.next
				l.next = nil
			}
			ll.size--
		}
		pl = l
		l = l.next
	}
	return nil
}
func (ll *lldata) RemoveAt(index int) (interface{}, error) {
	pl := ll.head

	//Move the pointer from head to previous of given index
	for i := 1; i < index; i++ {
		pl = pl.next
	}

	var entry *link = nil
	if ll.size == 1 {
		ll.head = nil
		entry = pl
	} else {
		entry = pl.next
	}

	pl.next = entry.next
	entry.next = nil
	ll.size--
	return entry.data, nil
}
func (ll *lldata) Size() int {
	return ll.size
}

func NewLinkedList() List {
	var lld interface{} = &lldata{
		head: nil,
		size: 0,
	}
	ll, _ := lld.(List)
	return ll
}
