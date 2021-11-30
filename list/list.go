package list

import "github.com/rewantsoni/go-datastructures/iterator"

type List interface {
	Size() int
	IsEmpty() bool
	Add(element int) bool
	AddAll(elements ...int) bool
	AddAt(i int, element int) bool
	GetAt(i int) (int, bool)
	Contains(element int) bool
	IndexOf(element int) int
	Replace(old int, new int) bool
	Set(i int, new int) bool
	Remove(element int) bool
	RemoveAt(i int) (int, bool)
	RetainAll(elements ...int) List
	RemoveAll(elements ...int) List
	ReplaceAll(f func(e int) int)
	Iterator() iterator.Iterator
}

//TODO: RetainAll, TDD
//TODO: ContainAll(elements ...), ReplaceAll(func(e)), RemoveALL(elements ...){RetainALL}
//TODO: iterator.. iterateList