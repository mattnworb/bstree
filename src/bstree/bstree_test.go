package bstree

import "testing"

func TestContains(t *testing.T) {
	tree := New()

	if tree.Contains(1) {
		t.Errorf("Contains(1) on an empty tree should return false")
	}

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	for _, i := range []int{1, 2, 3} {
		if !tree.Contains(i) {
			t.Errorf("After insertion of [1, 2, 3] tree.Contains(%v) should return true", i)
		}
	}

	for _, i := range []int{4, 5, 6} {
		if tree.Contains(i) {
			t.Errorf("After insertion of [1, 2, 3] tree.Contains(%v) should return false", i)
		}
	}
}

func TestSize(t *testing.T) {
	tree := New()

	s := tree.Size()
	if s != 0 {
		t.Errorf("size of a new tree should be 0, but got %v", s)
	}

	tree.Insert(1)

	s = tree.Size()
	if s != 1 {
		t.Errorf("expected size of 1 after one insertion, but got %v", s)
	}
}
