// A simple Binary Search Tree implementation.
//
// This implementation currently only allows integers to be stored.
//
// The tree cannot contain duplicate items.
package bstree

// A single node in the binary search tree
type node struct {
	value    int
	left     *node
	right    *node
	is_black bool
}

var sentinel *node = &node{-1, nil, nil, true}

func (n *node) size() int {
	size := 1
	if n.left != sentinel {
		size += n.left.size()
	}
	if n.right != sentinel {
		size += n.right.size()
	}
	return size
}

func (n *node) contains(value int) bool {
	switch {
	case n == sentinel:
		return false
	case value == n.value:
		return true
	case value < n.value:
		return n.left.contains(value)
	case value > n.value:
		return n.right.contains(value)
	default:
		//impossible
		return false
	}
}

type BinarySearchTree struct {
	root *node
}

// Creates a new empty tree
func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Inserts a value to the tree
func (tree *BinarySearchTree) Insert(value int) {
	if tree.root == nil {
		//root is always black
		tree.root = &node{value, sentinel, sentinel, true}
	} else {
		innerInsert(tree.root, value)
	}
}

func innerInsert(n *node, value int) {
	if value < n.value {
		if n.left == sentinel {
			// add red node
			n.left = &node{value, sentinel, sentinel, false}
		} else {
			innerInsert(n.left, value)
		}
	} else if value > n.value {
		if n.right == sentinel {
			// add red node
			n.right = &node{value, sentinel, sentinel, false}
		} else {
			innerInsert(n.right, value)
		}
	}
}

func (tree *BinarySearchTree) Contains(value int) bool {
	if tree.IsEmpty() {
		return false
	}
	return tree.root.contains(value)
}

func (tree *BinarySearchTree) Remove() {
}

func (tree *BinarySearchTree) Size() int {
	switch {
	case tree.IsEmpty():
		return 0
	default:
		return tree.root.size()
	}
}

func (tree *BinarySearchTree) IsEmpty() bool {
	return tree.root == nil
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
	if n.left != sentinel {
		contents = traverse(contents, n.left)
	}
	contents = append(contents, n.value)
	if n.right != sentinel {
		contents = traverse(contents, n.right)
	}
	return contents
}
