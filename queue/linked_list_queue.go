package queue

import "github.com/rewantsoni/go-datastructures/list"

type LinkedListQueue struct {
	ll *list.LinkedList
}

func NewLinkedListQueue() Queue {
	ll := list.NewLinkedList()
	return &LinkedListQueue{
		ll: ll,
	}
}

func (lq LinkedListQueue) Clear() {
	lq.ll.Clear()
}

func (lq LinkedListQueue) Dequeue() int {
	return lq.ll.RemoveFirst()
}

func (lq LinkedListQueue) Empty() bool {
	return lq.ll.IsEmpty()
}

func (lq LinkedListQueue) Enqueue(element int) bool {
	return lq.ll.AddLast(element)
}

func (lq LinkedListQueue) Peek() int {
	return lq.ll.GetFirst()
}

func (lq LinkedListQueue) Size() int {
	return lq.ll.Size()
}
