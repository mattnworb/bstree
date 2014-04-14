package main

import (
    "container/list"
	"fmt"
    //"stack" 
	"math/rand"
	"time"
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
        //pop
        n := stack.Back()
        stack.Remove(n)

        switch node := n.Value.(type) {
        case *Node:
            fmt.Printf("popped node %v from stack\n", node)
            if node.value == val {
                return node
            }
            stack.PushBack(node.left)
            stack.PushBack(node.right)
        default:
            panic("Wrong type in node stack")
        }
    }
    return nil
}

var maxSize int = 1000

func value() int {
	return rand.Intn(maxSize)
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

func makeRandomTree(size int) *Node {
	root := &Node{value(), nil, nil}
	current := root
	for i := 0; i < size; i++ {
		current.left = &Node{value(), nil, nil}
		current.right = &Node{value(), nil, nil}
		current = current.left
	}
	return root
}

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
	rand.Seed(time.Now().UnixNano())

	var root Node = Node{value(), nil, nil}
	fmt.Printf("Height of tree with one element is %v\n", root.Height())

	root.left = &Node{2, nil, nil}
	fmt.Printf("Height of tree with two elements is %v\n", root.Height())

	root.right = &Node{2, nil, nil}
	fmt.Printf("Height of tree with three elements balanced is %v\n", root.Height())

	// printTree2(&root)
	// printTree2(makeRandomTree(10))

    fmt.Println()
    size := 10
    utree := makeUniformTree(size)
    fmt.Printf("Uniform tree of size %v has height %v\n", size, utree.Height())
    printTree3(utree)

    fmt.Printf("Searched for 5 in tree, answer was %v\n", Search(utree, 5))
    fmt.Printf("Searched for 11 in tree, answer was %v\n", Search(utree, 11))
}
