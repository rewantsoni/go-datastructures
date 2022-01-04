package list

import (
	"fmt"
	"github.com/rewantsoni/go-datastructures/iterator"
	"github.com/rewantsoni/go-datastructures/operators"
)

type node struct {
	data int
	next *node
	prev *node
}

func newNode(element int) *node {
	return &node{
		data: element,
	}
}

type LinkedList struct {
	size int

	first *node
	last  *node
}

func NewLinkedList(elements ...int) List {
	ll := &LinkedList{
		size:  nought,
		first: nil,
	}

	if len(elements) == 0 {
		return ll
	}

	if !ll.AddAll(elements...) {
		return nil
	}
	return ll
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Size() == 0
}

func (ll *LinkedList) Add(element int) bool {
	return ll.addAll(ll.Size(), element)
}

func (ll *LinkedList) AddAll(elements ...int) bool {
	return ll.addAll(ll.Size(), elements...)
}

func (ll *LinkedList) AddAt(index int, element int) bool {
	return ll.addAll(index, element)
}

func (ll *LinkedList) GetAt(index int) int {
	if ll.IsEmpty() || index < 0 || index >= ll.Size() {
		panic(fmt.Sprintf("panic: index %d is out of bound length is %d", index, ll.Size()))
	}

	return ll.traverseTo(index).data
}

func (ll *LinkedList) Contains(element int) bool {
	return ll.IndexOf(element) != -1
}

func (ll *LinkedList) ContainsAll(elements ...int) bool {
	for _, element := range elements {
		if !ll.Contains(element) {
			return false
		}
	}
	return true
}

func (ll *LinkedList) IndexOf(element int) int {
	return ll.find(element)
}

func (ll *LinkedList) Replace(oldElement int, newElement int) bool {
	if ll.IsEmpty() {
		return false
	}

	res := false
	temp := ll.first
	for temp != nil {
		if temp.data == oldElement {
			temp.data = newElement
			res = true
		}
		temp = temp.next
	}

	return res

}

func (ll *LinkedList) Set(index int, newElement int) bool {
	if ll.IsEmpty() || index < 0 || index >= ll.Size() {
		return false
	}

	ll.traverseTo(index).data = newElement
	return true
}

func (ll *LinkedList) Remove(element int) bool {
	//TODO implement me
	panic("implement me")
}

func (ll *LinkedList) RemoveAt(index int) (int, bool) {
	//TODO implement me
	panic("implement me")
}

func (ll *LinkedList) RetainAll(elements ...int) {
	//TODO implement me
	panic("implement me")
}

func (ll *LinkedList) RemoveAll(elements ...int) {
	//TODO implement me
	panic("implement me")
}

func (ll *LinkedList) ReplaceAll(operator operators.UnaryOperator) {
	//TODO implement me
	panic("implement me")
}

func (ll *LinkedList) Iterator() iterator.Iterator {
	//TODO implement me
	panic("implement me")
}

func (ll *LinkedList) addAll(index int, elements ...int) bool {
	for i, element := range elements {
		if !ll.add(index+i, element) {
			return false
		}
	}
	return true
}

func (ll *LinkedList) add(index int, element int) bool {

	if !(index >= 0 && index <= ll.Size()) {
		return false
	}

	newData := newNode(element)
	if ll.IsEmpty() {
		ll.first = newData
		ll.last = newData
		ll.size++
		return true
	}

	if index == 0 {
		newData.next = ll.first
		ll.first.prev = newData
		ll.first = newData
		ll.size++
		return true
	}

	if index == ll.Size() {
		newData.prev = ll.last
		ll.last.next = newData
		ll.last = newData
		ll.size++
		return true
	}

	curData := ll.traverseTo(index)
	newData.next = curData
	curData.prev.next = newData
	newData.prev = curData.prev
	curData.prev = newData
	ll.size++

	return true
}

func (ll *LinkedList) traverseTo(index int) *node {
	temp := ll.first

	for i := 0; i < index; i++ {
		temp = temp.next
	}

	return temp

}

func (ll *LinkedList) find(element int) int {
	temp := ll.first
	for i := 0; i < ll.Size(); i++ {
		if temp.data == element {
			return i
		}
		temp = temp.next
	}
	return -1
}
