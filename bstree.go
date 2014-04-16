// Package bstree provides a simple Binary Search Tree implementation.
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
	value   int
	left    *node
	right   *node
	parent  *node
	isBlack bool
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

// BinarySearchTree provides all of the tree functionality from this package.
// Create a new empty BinarySearchTree with `bstree.New()`. Functions that
// operate on the tree structure are defined on this type.
type BinarySearchTree struct {
	root *node
	size int
}

// New creates a new empty tree
func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert adds a value to the tree. If the value is already stored in the tree,
// then Insert() has no effect.
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
		}
		return tree.innerInsert(n.left, value)
	} else if value > n.value {
		if n.right == nil {
			// add red node
			newNode := &node{value, nil, nil, n, false}
			n.right = newNode
			return newNode, true
		}
		return tree.innerInsert(n.right, value)
	} else {
		//n.value == value, don't insert a duplicate into the tree
		return n, false
	}
}

func (tree *BinarySearchTree) fixAfterInsertion(n *node) {
	n.isBlack = false
	for n != nil && n != tree.root && !n.parent.isBlack {
		if parentOf(n) == safeLeft(parentOf(parentOf(n))) {
			r := safeRight(parentOf(parentOf(n)))
			if !colorOf(r) {
				setColor(parentOf(n), true)
				setColor(r, true)
				setColor(parentOf(parentOf(n)), false)
				n = parentOf(parentOf(n))
			} else {
				if n == safeRight(parentOf(n)) {
					n = parentOf(n)
					tree.rotateLeft(n)
				}
				setColor(parentOf(n), true)
				setColor(parentOf(parentOf(n)), false)
				tree.rotateRight(parentOf(parentOf(n)))
			}
		} else {
			l := safeLeft(parentOf(parentOf(n)))
			if !colorOf(l) {
				setColor(parentOf(n), true)
				setColor(l, true)
				setColor(parentOf(parentOf(n)), false)
				n = parentOf(parentOf(n))
			} else {
				if n == safeLeft(parentOf(n)) {
					n = parentOf(n)
					tree.rotateRight(n)
				}
				setColor(parentOf(n), true)
				setColor(parentOf(parentOf(n)), false)
				tree.rotateLeft(parentOf(parentOf(n)))
			}
		}
	}
	tree.root.isBlack = true
}

func safeLeft(n *node) *node {
	if n == nil {
		return nil
	}
	return n.left
}

func safeRight(n *node) *node {
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

func setColor(n *node, black bool) {
	if n != nil {
		n.isBlack = black
	}
}

func colorOf(n *node) bool {
	if n == nil {
		return true
	}
	return n.isBlack
}

func (tree *BinarySearchTree) rotateLeft(n *node) {
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

func (tree *BinarySearchTree) rotateRight(n *node) {
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

// Contains tests if the value is stored in the tree.
func (tree *BinarySearchTree) Contains(value int) bool {
	if tree.IsEmpty() {
		return false
	}
	return tree.root.contains(value)
}

// Remove removes value from tree. If value is not already in the tree, then TODO
func (tree *BinarySearchTree) Remove() {
	//TODO implement
}

// Size returns the number of elements stored in the tree.
func (tree *BinarySearchTree) Size() int {
	return tree.size
}

// IsEmpty tests if the tree is empty, i.e. if Size() == 0
func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.size == 0
}

// Contents returns the contents of the tree, from an in-order traversal
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

// Max returns the maximum value currently stored in the tree. If the tree is
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

// Min returns the minimum value currently stored in the tree. If the tree is
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
