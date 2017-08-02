package collections

import "fmt"

type direction int

const (
	EQUAL direction = iota
	LEFT
	RIGHT
)

type treeNode struct {
	data  interface{}
	left  *treeNode
	right *treeNode
}

func (t *treeNode) isLeafNode() bool {
	return t.left == nil && t.right == nil
}

func (t *treeNode) hasOneChild() bool {
	return t.left == nil || t.right == nil
}

func (t *treeNode) findSuccessor() (*treeNode, *treeNode) {
	var parentSucc *treeNode = nil
	succ := t.right
	for succ.left != nil {
		parentSucc = succ
		succ = succ.left
	}
	return succ, parentSucc
}

func (t *treeNode) String() string {
	var ldata interface{} = nil
	if t.left == nil {
		ldata = nil
	} else {
		ldata = t.left.data
	}

	var rdata interface{} = nil
	if t.right == nil {
		rdata = nil
	} else {
		rdata = t.right.data
	}

	return fmt.Sprintf("[%v <- %v -> %v]", ldata, t.data, rdata)
}

type binarySearchTree struct {
	root    *treeNode
	size    int
	compare func(a, b interface{}) int
}

func (bst *binarySearchTree) Insert(key interface{}) error {

	nnode := &treeNode{key, nil, nil}

	if bst.root == nil {
		bst.root = nnode
		bst.size++
		return nil
	}

	node, pnode := bst.findNode(key)
	if node != nil {
		return &collError{fmt.Sprintf("Key [%v] is duplicate", key)}
	}

	if bst.isLeftNode(key, pnode.data) {
		pnode.left = nnode
	} else {
		pnode.right = nnode
	}
	bst.size++
	return nil
}
func (bst *binarySearchTree) Find(key interface{}) interface{} {
	node, _ := bst.findNode(key)
	if node == nil {
		return nil
	}

	return node.data
}

func (bst *binarySearchTree) findNode(key interface{}) (*treeNode, *treeNode) {
	node := bst.root
	var pnode *treeNode = nil

	for true {
		if node == nil {
			return node, pnode
		}
		res := bst.findDirection(key, node.data)

		if res == EQUAL {
			return node, pnode
		} else if res == RIGHT {
			pnode = node
			node = node.right
		} else {
			pnode = node
			node = node.left
		}

	}
	return nil, pnode
}

func (bst *binarySearchTree) Delete(key interface{}) (interface{}, error) {
	node, parentNode := bst.findNode(key)

	// When node to be deleted is a leaf node
	if node.isLeafNode() {
		bst.deleteLeafNode(node, parentNode)
	} else if node.hasOneChild() { // When node to be deleted has one child.
		bst.deleteNodeWithOneChild(node, parentNode)
	} else {
		// When node to be deleted has two childrens
		successor, parentSuccessor := node.findSuccessor()

		if bst.isSuccessorRightChild(node, successor) {
			bst.deleteNodeWithRightChildAsSuccessor(node, parentNode, successor)
		} else {
			bst.deleteNodeWithLeftDescendantOfRightChildAsSuccessor(node, parentNode, successor, parentSuccessor)
		}
	}
	bst.size--
	return node.data, nil
}

func (bst *binarySearchTree) deleteLeafNode(node *treeNode, parentNode *treeNode) {

	if parentNode == nil {
		bst.root = nil
	} else if bst.isLeftNode(node.data, parentNode.data) {
		parentNode.left = nil
	} else {
		parentNode.right = nil
	}

}
func (bst *binarySearchTree) deleteNodeWithOneChild(node *treeNode, parentNode *treeNode) {

	var childNode *treeNode = nil
	if node.left != nil {
		childNode = node.left
	} else {
		childNode = node.right
	}

	if parentNode == nil {
		bst.root = childNode
	} else if bst.isLeftNode(node.data, parentNode.data) {
		parentNode.left = childNode
	} else {
		parentNode.right = childNode
	}
}

func (bst *binarySearchTree) deleteNodeWithRightChildAsSuccessor(node *treeNode, parentNode *treeNode, successor *treeNode) {

	if parentNode == nil {
		bst.root = successor
	} else if bst.isLeftNode(node.data, parentNode.data) {
		parentNode.left = successor
	} else {
		parentNode.right = successor
	}
	successor.left = node.left

}

func (bst *binarySearchTree) deleteNodeWithLeftDescendantOfRightChildAsSuccessor(node *treeNode, parentNode *treeNode, successor *treeNode, parentSuccessor *treeNode) {

	if parentNode == nil {
		bst.root = successor
	} else if bst.isLeftNode(node.data, parentNode.data) {
		parentNode.left = successor
	} else {
		parentNode.right = successor
	}
	parentSuccessor.left = successor.right
	successor.right = node.right
	successor.left = node.left
}

func (bst *binarySearchTree) Size() int {
	return bst.size
}

type inorderEntry struct {
	node        *treeNode
	leftVisited bool
}

func (ie *inorderEntry) String() string {
	return fmt.Sprintf("[%v, leftVisited=%v]", ie.node, ie.leftVisited)
}

func (bst *binarySearchTree) print() {

	if bst.root == nil {
		fmt.Println("No Root")
		return
	}
	st, _ := NewStack()
	count := 0
	st.Push(&inorderEntry{bst.root, false})
	for !st.IsEmpty() {

		if count == 100 {
			break
		}
		count++

		entry, _ := st.Peek()
		e, _ := entry.(*inorderEntry)
		if !e.leftVisited && e.node.left != nil {
			e.leftVisited = true
			st.Push(&inorderEntry{e.node.left, false})
			continue
		}

		entry, _ = st.Pop()
		e, _ = entry.(*inorderEntry)

		fmt.Print(e.node.data, "->")

		if e.node.right != nil {
			st.Push(&inorderEntry{e.node.right, false})
			continue
		}
	}
	fmt.Println("")
}

func NewBinarySearchTree(compare func(a, b interface{}) int) Tree {
	var bst interface{} = &binarySearchTree{nil, 0, compare}
	tree, _ := bst.(Tree)
	return tree
}

func (bst *binarySearchTree) findDirection(nodeData interface{}, parentNodeData interface{}) direction {
	r := bst.compare(nodeData, parentNodeData)
	if r < 0 {
		return LEFT
	} else if r > 0 {
		return RIGHT
	} else {
		return EQUAL
	}
}

func (bst *binarySearchTree) isSuccessorRightChild(node *treeNode, successor *treeNode) bool {
	return node.right == successor
}

func (bst *binarySearchTree) isLeftNode(nodeData interface{}, parentNodeData interface{}) bool {
	return bst.findDirection(nodeData, parentNodeData) == LEFT
}

func (bst *binarySearchTree) isRightNode(nodeData interface{}, parentNodeData interface{}) bool {
	return !bst.isLeftNode(nodeData, parentNodeData)
}
