package main

import (
    "container/list"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Height() int {
	height := 1
	if n.left != nil {
		height += n.left.Height()
	}
	if n.right != nil {
		r := n.right.Height()
		if r > height {
			height = r
		}
	}
	return height
}


/** Searches for node with given value */
func Search(root *Node, val int) *Node {
    //s := stack.NewStack()
    // use list as a stack
    stack := list.New()
    stack.PushBack(root)
    for stack.Len() > 0 {
        //pop from back of list
        //does a type assertion at the end
        node := stack.Remove(stack.Back()).(*Node)

        //fmt.Printf("popped node %v from stack\n", node)
        if node.value == val {
            return node
        }
        if node.left != nil {
            stack.PushBack(node.left)
        }
        if node.right != nil {
            stack.PushBack(node.right)
        }
    }
    return nil
}

func newNode(val int) *Node {
    return &Node{val, nil, nil}
}

func makeUniformTree(size int) *Node {
    root := newNode(1)
    remaining := size - 1
    counter := 2

    queue := list.New()
    queue.PushBack(root)

    for remaining > 0 {

        // take first element from queue
        el := queue.Remove(queue.Front())
        node := el.(*Node)

        // add left child to node and push that left child onto queue
        node.left = newNode(counter)
        counter += 1
        remaining -= 1
        queue.PushBack(node.left)

        // do same for right
        if remaining > 0 {
            node.right = newNode(counter)
            counter += 1
            remaining -= 1
            queue.PushBack(node.right)
        }
    }
    return root
}

/*
// a very dumb way to print trees
func printTree(prefix string, root *Node) {
	fmt.Printf("%s.value = %v\n", prefix, root.value)
	if root.left != nil {
		printTree(prefix+".left", root.left)
	}
	if root.right != nil {
		printTree(prefix+".right", root.right)
	}
}

func printTree2(root *Node) {
	printTree2_("", root)
}
func printTree2_(prefix string, root *Node) {
	fmt.Printf("%s%v\n", prefix, root.value)
	if root.left != nil {
		printTree2_(prefix+"  ", root.left)
	}
	if root.right != nil {
		printTree2_(prefix+"  ", root.right)
	}
} 
*/

func printTree3(root *Node) {
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        node := queue.Remove(queue.Front()).(*Node)
        fmt.Printf("value=%v, left=%v, right=%v\n", node.value, node.left, node.right)
        if (node.left != nil) {
            queue.PushBack(node.left)
        }
        if (node.right != nil) {
            queue.PushBack(node.right)
        }
    }
}

func main() {
    size := 10
    utree := makeUniformTree(size)
    fmt.Printf("Uniform tree of size %v has height %v\n", size, utree.Height())
    printTree3(utree)

    fmt.Printf("Searched for 5 in tree, answer was %v\n", Search(utree, 5))
    fmt.Printf("Searched for 11 in tree, answer was %v\n", Search(utree, 11))
}
