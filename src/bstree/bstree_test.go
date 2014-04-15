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
