package list

import (
	"fmt"
	"github.com/rewantsoni/go-datastructures/iterator"
	"github.com/rewantsoni/go-datastructures/operators"
	"strings"
)

type node struct {
	data int
	next *node
	prev *node
}

type LinkedList struct {
	size int

	first *node
	last  *node
}

type linkedListIterator struct {
	currNode *node
	ll       *LinkedList
}

func newNode(element int) *node {
	return &node{
		data: element,
	}
}

func NewLinkedList(elements ...int) List {
	ll := &LinkedList{
		size:  nought,
		first: nil,
		last:  nil,
	}

	if len(elements) == 0 {
		return ll
	}

	if !ll.AddAll(elements...) {
		return nil
	}
	return ll
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

func (ll *LinkedList) Clear() {
	ll.first = nil
	ll.last = nil
	ll.size = 0
}

func (ll *LinkedList) Clone() (bool, List) {
	if ll.IsEmpty() {
		return true, NewLinkedList()
	}
	return ll.SubList(nought, ll.Size())
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

func (ll *LinkedList) GetAt(index int) int {
	if ll.IsEmpty() || index < 0 || index >= ll.Size() {
		panic(fmt.Sprintf("panic: index %d is out of bound length is %d", index, ll.Size()))
	}

	return ll.traverseTo(index).data
}

func (ll *LinkedList) IndexOf(element int) int {
	return ll.find(element)
}

//TODO: test
func (ll *LinkedList) IsEmpty() bool {
	return ll.Size() == 0
}

//TODO: test
func (ll *LinkedList) Iterator() iterator.Iterator {
	return newLinkedListIterator(ll)
}

//TODO: test
func (ll *LinkedList) LastIndexOf(element int) int {
	return ll.findLast(element)
}

//TODO: make it more readable
func (ll *LinkedList) Remove(element int) bool {

	if ll.IsEmpty() {
		return false
	}

	if ll.first.data == element {
		ll.first = ll.first.next
		ll.first.prev = nil
		ll.size--
		return true
	}

	if ll.last.data == element {
		ll.last = ll.last.prev
		ll.last.next = nil
		ll.size--
		return true
	}

	cur := ll.first
	for cur != nil {
		if cur.data == element {
			cur.prev.next = cur.next
			cur.next.prev = cur.prev
			ll.size--
			return true
		}
		cur = cur.next
	}
	return false
}

//TODO: test
func (ll *LinkedList) RemoveAll(elements ...int) {
	ll.filterLinkedList(false, elements...)
}

//TODO: make it more readable
func (ll *LinkedList) RemoveAt(index int) (int, bool) {
	if ll.IsEmpty() || index < 0 || index >= ll.size {
		return -1, false
	}

	if index == 0 {
		temp := ll.first.data
		ll.first = ll.first.next
		ll.first.prev = nil
		ll.size--
		return temp, true
	}

	if index == ll.size-1 {
		temp := ll.last.data
		ll.last = ll.last.prev
		ll.last.next = nil
		ll.size--
		return temp, true
	}

	cur := ll.first
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	temp := cur.data
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	ll.size--
	return temp, true
}

//TODO: make it more readable
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

func (ll *LinkedList) ReplaceAll(operator operators.UnaryOperator) {
	cur := ll.first
	for cur != nil {
		cur.data = operator.Apply(cur.data)
		cur = cur.next
	}
}

//TODO: test
func (ll *LinkedList) RetainAll(elements ...int) {
	ll.filterLinkedList(true, elements...)
}

func (ll *LinkedList) Set(index int, newElement int) bool {
	if ll.IsEmpty() || index < 0 || index >= ll.Size() {
		return false
	}

	ll.traverseTo(index).data = newElement
	return true
}

func (ll *LinkedList) Size() int {
	return ll.size
}

func (ll *LinkedList) SubList(start, end int) (bool, List) {

	if (start >= end) || (start < 0 || start >= ll.Size()) || (end < 0 || end > ll.Size()) {
		return false, nil
	}

	tempList := NewLinkedList()

	cur := ll.first

	for i := 0; i < ll.Size(); i++ {
		if i >= start && i < end {
			tempList.Add(cur.data)
		}
		cur = cur.next
	}

	return true, tempList
}

func (ll *LinkedList) String() string {
	sb := strings.Builder{}
	temp := ll.first
	for temp != nil {
		sb.WriteString(fmt.Sprintf("%d ", temp.data))
		temp = temp.next
	}

	return sb.String()
}

//TODO: test
func (lli *linkedListIterator) HasNext() bool {
	return lli.currNode != nil
}

//TODO: test
func (lli *linkedListIterator) Next() int {
	if lli.currNode == nil {
		panic("panic: linked list is empty")
	}
	temp := lli.currNode.data
	lli.currNode = lli.currNode.next
	return temp
}

//Helper Functions
func (ll *LinkedList) addAll(index int, elements ...int) bool {
	for i, element := range elements {
		if !ll.add(index+i, element) {
			return false
		}
	}
	return true
}

//TODO: make it more readable
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

func (ll *LinkedList) findLast(element int) int {
	temp := ll.last
	for i := ll.Size() - 1; i > 0; i-- {
		if temp.data == element {
			return i
		}
		temp = temp.prev
	}
	return -1
}

func (ll *LinkedList) filterLinkedList(retain bool, elements ...int) {
	cache := map[int]bool{}

	for _, e := range elements {
		cache[e] = true
	}

	cur := ll.first

	for cur != nil {
		if cache[cur.data] {
			if !retain {
				//can pass node and delete wrt node
				ll.Remove(cur.data)
			}
		} else {
			if retain {
				ll.Remove(cur.data)
			}
		}
		cur = cur.next
	}
}

func newLinkedListIterator(ll *LinkedList) *linkedListIterator {
	return &linkedListIterator{
		currNode: ll.first,
		ll:       ll,
	}
}
