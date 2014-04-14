package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	s := NewStack()
	if empty := s.IsEmpty(); !empty {
		t.Errorf("new stack should be empty, but got %v expecting %v", empty, false)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	if err == nil {
		t.Error("Expected error popping from empty stack")
	}
}

func TestPopNonEmptyStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	val, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error from Pop(), got %v", err)
	} else {
		if val != 2 {
			t.Errorf("Pop() returned wrong value, expected 2 got %v", val)
		}
		val, err := s.Pop()
        if err != nil {
            t.Errorf("Unexpected error from Pop(), got %v", err)
        } else if val != 1 {
			t.Errorf("Pop() returned wrong value, expected 1 got %v", val)
		}
	}
}
