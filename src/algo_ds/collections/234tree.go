package collections

import "fmt"

const (
	TOTAL_ELEMENTS_234 = 3
	TOTAL_LINKS_234    = TOTAL_ELEMENTS_234 + 1
)

// A 2-3-4 Tree  structure
type tree234 struct {
	root        *node234
	size        int
	compareElms func(a, b interface{}) int
}

func (t *tree234) Insert(key interface{}) error {
	if t.root == nil {
		t.createRoot(key)
		t.size++
		return nil
	}

	node, _ := t.findNode234WithSplit(key)
	node.addElement(key)
	t.size++
	return nil
}

func (t *tree234) Find(key interface{}) interface{} {
	return t.findNode234(key)
}
func (t *tree234) Delete(key interface{}) (interface{}, error) {
	return nil, nil
}
func (t *tree234) Size() int {
	return t.size
}

func New234Tree(compare func(a, b interface{}) int) Tree {
	var t interface{} = &tree234{nil, 0, compare}
	tree, _ := t.(Tree)
	return tree
}

// A node structure
type node234 struct {
	parent      *node234
	elements    []interface{}
	links       []*node234
	compareElms func(a, b interface{}) int
	elmsSize    int
	linksSize   int
}

func (n *node234) isFull() bool {
	return n.elmsSize == TOTAL_ELEMENTS_234
}

func (n *node234) createNode() *node234 {
	nnode := createNode(n.compareElms)
	nnode.parent = n
	return nnode
}

func (n *node234) addElement(element interface{}) {
	if element == nil {
		return
	}
	pos := n.findElementPosition(element)
	shiftRight(n.elements, pos, TOTAL_ELEMENTS_234)
	n.elements[pos] = element
	n.elmsSize++
}

func (n *node234) removeElement(element interface{}) {
	if element == nil {
		return
	}
	pos := n.findElementPosition(element)
	n.elements[pos] = nil
	shiftLeft(n.elements, pos, TOTAL_ELEMENTS_234)
	n.elements[TOTAL_ELEMENTS_234-1] = nil
	n.elmsSize--
}

func (n *node234) addLink(link *node234) {

	if link == nil {
		return
	}

	linkPos := n.findLinkPosition(link.elements[0])
	n.links[linkPos] = link
	link.parent = n
	n.linksSize++
}

// Find a position in the given node for given key
func (n *node234) findElementPosition(key interface{}) int {
	pos := 0
	for _, val := range n.elements {
		// if exisitng element is greater than given key, that is where we need to insert the key
		if val == nil || n.compareElms(val, key) <= 0 {
			return pos
		}
		pos++
	}
	return pos
}

func (n *node234) findElement(key interface{}) interface{} {
	for _, val := range n.elements {
		if val != nil && n.compareElms(val, key) == 0 {
			return val
		}
	}
	return nil
}

func (n *node234) findLink(key interface{}) *node234 {
	pos := n.findElementPosition(key)
	return n.links[pos]
}

func (n *node234) findLinkPosition(key interface{}) int {
	return n.findElementPosition(key)
}

func (n *node234) removeLink(link *node234) {
	if link == nil {
		return
	}

	pos := n.findLinkPosition(link.elements[0])
	n.links[pos].parent = nil
	n.linksSize--
}

func shiftRight(arr []interface{}, pos int, size int) {
	for i := size - 2; i >= pos; i-- {
		arr[i+1] = arr[i]
	}
}

func shiftLeft(arr []interface{}, pos int, size int) {
	for i := pos; i > size-1; i++ {
		arr[i] = arr[i+1]
	}
}

func (n *node234) String() string {

	parent := ""
	if n.parent != nil {
		parent = fmt.Sprintf("%v", n.parent.elements[0])
	}
	return fmt.Sprintf("{elms=[%v,%v,%v],parent=%v}", n.elements[0], n.elements[1], n.elements[2], parent)
}

func (n *node234) linksString() string {
	return fmt.Sprintf("L0=%v,L1=%v,L2=%v,L3=%v]", n.links[0], n.links[1], n.links[2], n.links[3])
}

// It does the split if it finds any full node while searching for right node to insert a new element
func (t *tree234) findNode234WithSplit(key interface{}) (*node234, *node234) {
	n := t.root
	var parent *node234 = nil

	if n.isFull() {
		new_root := t.splitNode(n)
		t.root = new_root
		n = new_root
	}

	for true {
		if n == nil {
			return parent, parent.parent
		}

		if n.isFull() {
			n = t.splitNode(n)
		}

		t := n.findLink(key)
		parent = n
		n = t
	}
	return nil, nil
}

func (t *tree234) findNode234(key interface{}) interface{} {
	n := t.root

	for true {
		if n == nil {
			return nil
		}

		k := n.findElement(key)
		if k != nil {
			return k
		}

		n = n.findLink(key)
	}
	return nil
}

func (t *tree234) splitNode(n *node234) *node234 {

	elmB := n.elements[1]
	elmC := n.elements[2]

	link2 := n.links[2]
	link3 := n.links[3]

	// Create  a new root node if n is root
	var par *node234 = nil
	if t.root == n {
		par = t.createNode()
		par.addLink(n)
	} else {
		par = n.parent
	}
	par.addElement(elmB)

	// Create a new right node
	new_right_node := t.createNode()
	new_right_node.addElement(elmC)

	// Add links 2 and 3 into new right node
	new_right_node.addLink(link2)
	new_right_node.addLink(link3)
	par.addLink(new_right_node)

	// Clean up existing links as they are moved to next right node
	n.removeElement(elmB)
	n.removeElement(elmC)

	n.removeLink(link2)
	n.removeLink(link3)
	return par
}

func (t *tree234) createNode() *node234 {
	return createNode(t.compareElms)
}

func (t *tree234) createRoot(element interface{}) {
	r := t.createNode()
	r.addElement(element)
	t.root = r

}

func createNode(compareElms func(a, b interface{}) int) *node234 {
	elms := make([]interface{}, TOTAL_ELEMENTS_234)
	links := make([]*node234, TOTAL_LINKS_234)
	nnode := &node234{nil, elms, links, compareElms, 0, 0}
	return nnode
}
