package stack

import "github.com/rewantsoni/go-datastructures/list"

type Stack struct {
	ll *list.LinkedList
}

func NewStack() *Stack {
	ll := list.NewLinkedList()
	return &Stack{
		ll: ll,
	}
}

func (s *Stack) Clear() {
	s.ll.Clear()
}

func (s *Stack) Empty() bool {
	return s.ll.IsEmpty()
}

func (s *Stack) Peek() int {
	return s.ll.GetFirst()
}

func (s *Stack) Pop() int {
	return s.ll.RemoveFirst()
}

func (s *Stack) Push(element int) bool {
	return s.ll.AddFirst(element)
}

func (s *Stack) Size() int {
	return s.ll.Size()
}
