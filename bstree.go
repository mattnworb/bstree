// A simple Binary Search Tree implementation.
//
// This implementation currently only allows integers to be stored.
//
// The tree cannot contain duplicate items. Adding an integer that is already
// stored in the tree results in no change being made to the tree (the result
// of Size() stays the same).
//
// This implementation is not safe for concurrent access.
//
// Internally, this package implements a red-black tree (with implementation
// heavily borrowed from java.util.TreeMap). This allows bstree to offer O(log
// n) access for inserts, removals, and searches in the average case.
package bstree

// A single node in the binary search tree
type node struct {
	value    int
	left     *node
	right    *node
	parent   *node
	is_black bool
}

func (n *node) contains(value int) bool {
	switch {
	case n == nil:
		return false
	case value < n.value:
		return n.left.contains(value)
	case value > n.value:
		return n.right.contains(value)
	default:
		// value == n.value
		return true
	}
}

type BinarySearchTree struct {
	root *node
	size int
}

// Creates a new empty tree
func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Inserts a value to the tree
func (tree *BinarySearchTree) Insert(value int) {
	if tree.root == nil {
		//root is always black
		tree.root = &node{value, nil, nil, nil, true}
		tree.size++
	} else {
		newNode, inserted := tree.innerInsert(tree.root, value)
		//only rebalance tree when a new value was actually added
		if inserted {
			tree.fixAfterInsertion(newNode)
			tree.size++
		}
	}
}

// inserts the value beneath node `n` (possibly calling itself recursively),
// returning the new node struct when done.
func (tree *BinarySearchTree) innerInsert(n *node, value int) (*node, bool) {
	if value < n.value {
		if n.left == nil {
			// add red node
			newNode := &node{value, nil, nil, n, false}
			n.left = newNode
			return newNode, true
		} else {
			return tree.innerInsert(n.left, value)
		}
	} else if value > n.value {
		if n.right == nil {
			// add red node
			newNode := &node{value, nil, nil, n, false}
			n.right = newNode
			return newNode, true
		} else {
			return tree.innerInsert(n.right, value)
		}
	} else {
		//n.value == value, don't insert a duplicate into the tree
		return n, false
	}
}

func (tree *BinarySearchTree) fixAfterInsertion(n *node) {
	n.is_black = false
	for n != nil && n != tree.root && !n.parent.is_black {
		if parentOf(n) == safe_left(parentOf(parentOf(n))) {
			r := safe_right(parentOf(parentOf(n)))
			if !safe_colorOf(r) {
				safe_setColor(parentOf(n), true)
				safe_setColor(r, true)
				safe_setColor(parentOf(parentOf(n)), false)
				n = parentOf(parentOf(n))
			} else {
				if n == safe_right(parentOf(n)) {
					n = parentOf(n)
					tree.rotate_left(n)
				}
				safe_setColor(parentOf(n), true)
				safe_setColor(parentOf(parentOf(n)), false)
				tree.rotate_right(parentOf(parentOf(n)))
			}
		} else {
			l := safe_left(parentOf(parentOf(n)))
			if !safe_colorOf(l) {
				safe_setColor(parentOf(n), true)
				safe_setColor(l, true)
				safe_setColor(parentOf(parentOf(n)), false)
				n = parentOf(parentOf(n))
			} else {
				if n == safe_left(parentOf(n)) {
					n = parentOf(n)
					tree.rotate_right(n)
				}
				safe_setColor(parentOf(n), true)
				safe_setColor(parentOf(parentOf(n)), false)
				tree.rotate_left(parentOf(parentOf(n)))
			}
		}
	}
	tree.root.is_black = true
}

func safe_left(n *node) *node {
	if n == nil {
		return nil
	}
	return n.left
}

func safe_right(n *node) *node {
	if n == nil {
		return nil
	}
	return n.right
}

func parentOf(n *node) *node {
	if n == nil {
		return nil
	}
	return n.parent
}

func safe_setColor(n *node, black bool) {
	if n != nil {
		n.is_black = black
	}
}

func safe_colorOf(n *node) bool {
	if n == nil {
		return true
	}
	return n.is_black
}

func (tree *BinarySearchTree) rotate_left(n *node) {
	if n != nil {
		r := n.right
		n.right = r.left
		if r.left != nil {
			r.left.parent = n
		}
		r.parent = n.parent
		if n.parent == nil {
			tree.root = r
		} else if n.parent.left == n {
			n.parent.left = r
		} else {
			n.parent.right = r
		}
		r.left = n
		n.parent = r
	}
}

func (tree *BinarySearchTree) rotate_right(n *node) {
	if n != nil {
		l := n.left
		n.left = l.right
		if l.right != nil {
			l.right.parent = n
		}
		l.parent = n.parent
		if n.parent == nil {
			tree.root = l
		} else if n.parent.right == n {
			n.parent.right = l
		} else {
			n.parent.left = l
		}
		l.right = n
		n.parent = l
	}
}

// Tests if the value is stored in the tree.
func (tree *BinarySearchTree) Contains(value int) bool {
	if tree.IsEmpty() {
		return false
	}
	return tree.root.contains(value)
}

// Removes value from tree. If value is not already in the tree, then TODO
func (tree *BinarySearchTree) Remove() {
	//TODO implement
}

// Returns the number of elements stored in the tree.
func (tree *BinarySearchTree) Size() int {
	return tree.size
}

// Tests if the tree is empty, i.e. if Size() == 0
func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.size == 0
}

// Returns the contents of the tree, from an in-order traversal
func (tree *BinarySearchTree) Contents() []int {
	if tree.IsEmpty() {
		return []int{}
	}
	return traverse([]int{}, tree.root)
}

// Traverses the tree in-order, adding the value of each node in-order to the `contents` slice and returning it
func traverse(contents []int, n *node) []int {
	if n.left != nil {
		contents = traverse(contents, n.left)
	}
	contents = append(contents, n.value)
	if n.right != nil {
		contents = traverse(contents, n.right)
	}
	return contents
}

// Returns the maximum value currently stored in the tree. If the tree is
// empty, returns -1.
func (tree *BinarySearchTree) Max() int {
	if tree.IsEmpty() {
		return -1
	}
	n := tree.root
	for n.right != nil {
		n = n.right
	}
	return n.value
}

// Returns the minimum value currently stored in the tree. If the tree is
// empty, returns -1.
func (tree *BinarySearchTree) Min() int {
	if tree.IsEmpty() {
		return -1
	}
	n := tree.root
	for n.left != nil {
		n = n.left
	}
	return n.value
}
