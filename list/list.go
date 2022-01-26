package list

import (
	"github.com/rewantsoni/go-datastructures/iterator"
	"github.com/rewantsoni/go-datastructures/operators"
)

type List interface {
	Add(element int) bool
	AddAll(elements ...int) bool
	AddAt(index int, element int) bool
	Clear()
	Clone() (bool, List)
	Contains(element int) bool
	ContainsAll(elements ...int) bool
	GetAt(index int) int
	IndexOf(element int) int
	IsEmpty() bool
	Iterator() iterator.Iterator
	LastIndexOf(element int) int
	Remove(element int) bool
	RemoveAt(index int) (int, bool)
	RemoveAll(elements ...int)
	Replace(oldElement int, newElement int) bool
	ReplaceAll(operator operators.UnaryOperator)
	RetainAll(elements ...int)
	Set(index int, newElement int) bool
	Size() int
	SubList(start, end int) (bool, List)
}
