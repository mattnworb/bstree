package bstree

import (
	"container/list"
	"fmt"
	"math/rand"
	"testing"
)

func TestContains(t *testing.T) {
	tree := New()

	if tree.Contains(1) {
		t.Errorf("Contains(1) on an empty tree should return false")
	}

	count := 10

	for i := 1; i <= count; i++ {
		tree.Insert(i)
		t.Logf("Inserted %v, tree is now %v", i, logTree(tree))

		if s := tree.Size(); s != i {
			t.Errorf("After %v insertions expected size to be %v but got %v", i, i, s)
		}

		//test that tree contains from 1 to i
		for j := 1; j < i; j++ {
			if !tree.Contains(j) {
				t.Errorf("After insertion of 1..%v, tree.Contains(%v) should return true but was false", i, j)
			}
		}
	}

	for i := count + 1; i < count*2; i++ {
		if tree.Contains(i) {
			t.Errorf("After insertion of 1-%v tree.Contains(%v) should return false but was true", count, i)
		}
	}
}

func TestSize(t *testing.T) {
	tree := New()

	s := tree.Size()
	if s != 0 {
		t.Errorf("size of a new tree should be 0, but got %v", s)
	}
	if !tree.IsEmpty() {
		t.Error("IsEmpty() should return true for an empty tree")
	}

	tree.Insert(1)

	s = tree.Size()
	if s != 1 {
		t.Errorf("expected size of 1 after one insertion, but got %v", s)
	}
	if tree.IsEmpty() {
		t.Error("IsEmpty() should return false for an empty tree")
	}

	for i := 2; i < 100; i++ {
		tree.Insert(i)
		s = tree.Size()
		if s != i {
			t.Errorf("expected size of %v after %v insertions, but got %v", i, i, s)
		}
	}
}

func TestContents(t *testing.T) {
	tree := New()

	contents := tree.Contents()
	if len(contents) > 0 {
		t.Errorf("Contents() of an empty tree should be an empty array, but saw array of size %v: %v", len(contents), contents)
	}

	tree.Insert(1)
	contents = tree.Contents()
	if len(contents) != 1 || contents[0] != 1 {
		t.Errorf("Contents() of a tree after Insert(1) should be [1], but saw array of size %v: %v", len(contents), contents)
	}

	tree.Insert(2)
	tree.Insert(3)
	contents = tree.Contents()
	if len(contents) != 3 || !isEqual(contents, []int{1, 2, 3}) {
		t.Errorf("Contents() of a tree after Insert(1, 2, 3) should be [1, 2, 3], but saw array of size %v: %v", len(contents), contents)
	}
}

func isEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRedBlackInvaraints(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		tree := New()

		//insert a random number of random values
		size := 1 + rand.Intn(20)
		for count := 0; count < size; count++ {
			num := rand.Intn(10000)
			//fmt.Printf("adding %v to tree\n", num)
			tree.Insert(num)
		}
		t.Logf("Testing with tree of size %v and contents %v", tree.Size(), tree.Contents())
		//fmt.Printf("Testing with tree of size %v and contents %v\n", tree.Size(), logTree(tree))

		assertRedBlackInvariants(tree, t)
		if t.Failed() {
			return
		}
	}
}

func assertRedBlackInvariants(tree *BinarySearchTree, t *testing.T) {
	// the root must be black
	//fmt.Println("testing that root is black")
	if !tree.root.is_black {
		t.Errorf("Tree %v should have a black root node", tree.Contents())
	}

	//fmt.Println("testing that every red node has either 0 children or 2 black children")
	visitNodes(tree, func(n *node) {
		if !n.is_black {
			if (n.left != nil && n.right == nil) || (n.left == nil && n.right != nil) {
				if !t.Failed() {
					t.Errorf("Found a red node with only one empty child in tree %v", logTree(tree))
				}
			} else if n.left != nil && n.right != nil {
				//if a red has two children, both must be black
				if !n.left.is_black || !n.right.is_black {
					if !t.Failed() {
						t.Errorf("Found a red node with only one black child in tree %v", logTree(tree))
					}
				}
			}
		}
	})

	// not totally sure if this is a valid test:
	/*
		//fmt.Println("testing that for any given node in the tree, the path to any descendant node has the same number of red nodes")
		visitNodes(tree, func(n *node) {
			red_counts := countRedNodesInPath(n)
			if len(red_counts) > 1 {
				t.Errorf("For any given node the path to any descendant node "+
					"should have same number of red nodes, but found counts of "+
					"%v from node %v in tree %v", red_counts, n.value, logTree(tree))
			}
		})
	*/
}

type visitor func(*node)

// Visits all nodes in the tree
func visitNodes(tree *BinarySearchTree, visit visitor) {
	//visit the nodes in BFS order
	queue := list.New()
	queue.PushBack(tree.root)

	for queue.Len() > 0 {
		//pop from front of queue
		var n *node = queue.Remove(queue.Front()).(*node)
		//fmt.Printf("At node %v\n", n)
		visit(n)
		if n.left != nil {
			queue.PushBack(n.left)
		}
		if n.right != nil {
			queue.PushBack(n.right)
		}
	}
}

func logTree(tree *BinarySearchTree) string {
	var s string = fmt.Sprintf("size=%v, contents=[", tree.Size())

	if !tree.IsEmpty() {
		s += logNode(tree.root, 0)
	}

	s += "]"
	return s
}

func logNode(n *node, level int) string {
	if n == nil { //sentinel {
		return "NIL"
	} else {
		color := "black"
		if !n.is_black {
			color = "red"
		}
		spacer := ""
		for i := 0; i < level; i++ {
			spacer += "  "
		}
		return fmt.Sprintf("{value=%v, color=%v, \n%vleft=%v, \n%vright=%v}",
			n.value, color, spacer+"  ", logNode(n.left, level+1), spacer+"  ", logNode(n.right, level+1))
	}
}

// for a given node, walk the path to all descendant leaf nodes and count how
// many red nodes are encountered along the way, returned as a map of `number
// of red nodes` to `number of occurrences of this count`
func countRedNodesInPath(n *node) map[int]int {
	occurrences := make(map[int]int)
	countRedWalker(n, occurrences, 0)
	return occurrences
}

func countRedWalker(n *node, occurrences map[int]int, count int) {
	if !n.is_black {
		count++
	}
	if n.left == nil && n.right == nil {
		//at leaf, update map
		occurrences[count]++
	} else {
		if n.left != nil {
			countRedWalker(n.left, occurrences, count)
		}
		if n.right != nil {
			countRedWalker(n.right, occurrences, count)
		}
	}
}

func TestMax(t *testing.T) {
	tree := New()
	max := 0
	for i := 0; i < 100; i++ {
		newval := rand.Int()
		tree.Insert(newval)
		if newval > max {
			max = newval
		}
	}

	if ans := tree.Max(); ans != max {
		t.Errorf("Expected tree.Max() to return %v but was %v", max, ans)
	}
}

func TestMaxEmptyTree(t *testing.T) {
	tree := New()
	if ans := tree.Max(); ans != -1 {
		t.Errorf("Expected Max() of empty tree to return -1 but was %v", ans)
	}
}

func TestMin(t *testing.T) {
	tree := New()
	newval := rand.Int()
	min := newval
	for i := 0; i < 100; i++ {
		tree.Insert(newval)
		if newval < min {
			min = newval
		}
		newval = rand.Int()
	}

	if ans := tree.Min(); ans != min {
		t.Errorf("Expected tree.Min() to return %v but was %v", min, ans)
	}
}

func TestMinEmptyTree(t *testing.T) {
	tree := New()
	if ans := tree.Min(); ans != -1 {
		t.Errorf("Expected Min() of empty tree to return -1 but was %v", ans)
	}
}
