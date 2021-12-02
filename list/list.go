package list

import (
	"github.com/rewantsoni/go-datastructures/iterator"
	"github.com/rewantsoni/go-datastructures/operators"
)

type List interface {
	Size() int
	IsEmpty() bool
	Add(element int) bool
	AddAll(elements ...int) bool
	AddAt(index int, element int) bool
	GetAt(i int) int
	Contains(element int) bool
	IndexOf(element int) int
	Replace(oldElement int, newElement int) bool
	Set(index int, newElement int) bool
	Remove(element int) bool
	RemoveAt(index int) (int, bool)
	RetainAll(elements ...int)
	RemoveAll(elements ...int)
	ReplaceAll(operator operators.UnaryOperator)
	Iterator() iterator.Iterator
}
