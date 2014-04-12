package main

import "fmt"
import "math/rand"
import "time"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Height() int {
    height := 1
    if (n.left != nil) {
        height += n.left.Height()
    }
    if (n.right != nil) {
        r := n.right.Height()
        if r > height {
            height = r
        }
    }
    return height
}

var maxSize int = 1000

func value() int {
	return rand.Intn(maxSize)
}

func makeTree(size int) Node {
	var root Node = Node{value(), nil, nil}
    current := &root
	for i := 0; i < size; i++ {
        current.left = &Node{value(), nil, nil}
        current.right = &Node{value(), nil, nil}
        current = current.left
	}
	return root
}

func printTree(prefix string, root *Node) {
    fmt.Printf("%s.value = %v\n", prefix, root.value)
    if root.left != nil {
        printTree(prefix + ".left", root.left)
    }
    if root.right != nil {
        printTree(prefix + ".right", root.right)
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

	//tree := makeTree(10)
	//fmt.Println(tree)
    //printTree("root", &tree)
}
