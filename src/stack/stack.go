package stack

import "container/list"

type IntStack interface {
	Push(int)
	Pop() (int, error)
	IsEmpty() bool
}

type stackImpl struct {
	list *list.List
}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func NewStack() IntStack {
	s := &stackImpl{list.New()}
	return s
}

func (impl *stackImpl) IsEmpty() bool {
	return impl.list.Len() == 0
}

func (impl *stackImpl) Push(val int) {
    impl.list.PushBack(val)
}

func (impl *stackImpl) Pop() (int, error) {
	if impl.IsEmpty() {
		return -1, MyError{"stack is empty!"}
	}
	e := impl.list.Back()
	switch v := e.Value.(type) {
	default:
		return -1, MyError{"wrong type in stack!"}
	case int:
		impl.list.Remove(e)
		return v, nil
	}
}
