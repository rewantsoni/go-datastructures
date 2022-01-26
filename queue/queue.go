package queue

type Queue interface {
	Clear()
	Dequeue() int
	Empty() bool
	Enqueue(element int) bool
	Peek() int
	Size() int
}
