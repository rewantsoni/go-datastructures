package list

import (
	"fmt"
	"github.com/rewantsoni/go-datastructures/iterator"
	"github.com/rewantsoni/go-datastructures/operators"
	"strings"
)

type ArrayList struct {
	capacity        int
	upperLoadFactor float64
	lowerLoadFactor float64
	scalingFactor   int
	size            int
	data            []int
}

type arrayListIterator struct {
	currentIndex int
	al           List
}

func NewArrayList(elements ...int) List {
	al := &ArrayList{
		size:            nought,
		capacity:        initialCapacity,
		upperLoadFactor: upperLoadFactor,
		lowerLoadFactor: lowerLoadFactor,
		scalingFactor:   scalingFactor,
		data:            make([]int, initialCapacity),
	}

	if len(elements) == 0 {
		return al
	}

	if !al.AddAll(elements...) {
		return nil
	}

	return al
}

func (al *ArrayList) Add(element int) bool {
	return al.addAll(al.Size(), element)
}

func (al *ArrayList) AddAll(elements ...int) bool {
	return al.addAll(al.Size(), elements...)
}

func (al *ArrayList) AddAt(index int, element int) bool {
	return al.addAll(index, element)
}

func (al *ArrayList) Clear() {
	for i := 0; i < al.Size(); i++ {
		al.data[i] = 0
	}
	al.size = nought
}

func (al *ArrayList) Clone() (bool, List) {
	if al.IsEmpty() {
		return true, NewArrayList()
	}
	return al.SubList(nought, al.Size())
}

func (al *ArrayList) Contains(element int) bool {
	return al.IndexOf(element) != -1
}

func (al *ArrayList) ContainsAll(elements ...int) bool {
	for _, element := range elements {
		if !al.Contains(element) {
			return false
		}
	}
	return true
}

func (al *ArrayList) GetAt(index int) int {
	if al.IsEmpty() || index < 0 || index >= al.Size() {
		panic(fmt.Sprintf("panic: index %d is out of bound length is %d", index, al.Size()))
	}

	return al.data[index]
}

func (al *ArrayList) IndexOf(element int) int {
	return al.find(element)
}

//TODO: test
func (al *ArrayList) IsEmpty() bool {
	return al.Size() == 0
}

func (al *ArrayList) Iterator() iterator.Iterator {
	return newArrayListIterator(al)
}

//TODO: test
func (al *ArrayList) LastIndexOf(element int) int {
	return al.findLast(element)
}

func (al *ArrayList) Remove(element int) bool {
	index := al.IndexOf(element)
	if index == -1 {
		return false
	}

	al.shiftLeft(index)

	return true
}

//TODO: Can return a panic
func (al *ArrayList) RemoveAt(index int) (int, bool) {
	if al.IsEmpty() || index < 0 || index >= al.Size() {
		return -1, false
	}

	e := al.data[index]
	al.shiftLeft(index)

	return e, true
}

func (al *ArrayList) RemoveAll(elements ...int) {
	al.filterArrayList(false, elements...)
}

func (al *ArrayList) Replace(oldElement int, newElement int) bool {
	if al.IsEmpty() {
		return false
	}

	ok := false
	for i := 0; i < al.Size(); i++ {
		if al.data[i] == oldElement {
			al.data[i] = newElement
			ok = true
		}
	}

	return ok
}

func (al *ArrayList) ReplaceAll(operator operators.UnaryOperator) {
	for i := 0; i < al.Size(); i++ {
		al.data[i] = operator.Apply(al.data[i])
	}
}

func (al *ArrayList) RetainAll(elements ...int) {
	al.filterArrayList(true, elements...)
}

//TODO: Can return a panic
func (al *ArrayList) Set(index int, newElement int) bool {
	if al.IsEmpty() || index < 0 || index >= al.Size() {
		return false
	}

	al.data[index] = newElement
	return true
}

func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) SubList(start, end int) (bool, List) {

	if (start >= end) || (start < 0 || start >= al.Size()) || (end < 0 || end > al.Size()) {
		return false, nil
	}

	tempList := NewArrayList()
	for i := start; i < end; i++ {
		if !tempList.Add(al.GetAt(i)) {
			return false, nil
		}
	}

	return true, tempList
}

func (al *ArrayList) String() string {
	sb := strings.Builder{}

	for i := 0; i < al.Size(); i++ {
		sb.WriteString(fmt.Sprintf("%d ", al.data[i]))
	}

	return sb.String()
}

func (ali *arrayListIterator) HasNext() bool {
	return ali.currentIndex < ali.al.Size()
}

func (ali *arrayListIterator) Next() int {
	e := ali.al.GetAt(ali.currentIndex)

	ali.currentIndex++
	return e
}

//Helper Functions
func (al *ArrayList) checkAndIncreaseLimit() {
	if al.Size() >= int(float64(al.capacity)*al.upperLoadFactor) {
		al.capacity *= al.scalingFactor
		al.data = resize(al.capacity, al.data)
	}
}

func (al *ArrayList) checkAndDecreaseLimit() {
	if al.Size() <= int(float64(al.capacity)*al.lowerLoadFactor) && al.capacity != initialCapacity {
		al.capacity /= al.scalingFactor
		al.data = resize(al.capacity, al.data)
	}
}

func (al *ArrayList) addAll(index int, elements ...int) bool {

	for i, d := range elements {
		if !al.add(index+i, d) {
			return false
		}
	}
	return true
}

func (al *ArrayList) add(index int, element int) bool {

	if !(index >= 0 && index <= al.Size()) {
		return false
	}

	al.checkAndIncreaseLimit()

	for i := al.Size(); i > index; i-- {
		al.data[i] = al.data[i-1]
	}

	al.data[index] = element
	al.size++

	return true
}

func (al *ArrayList) find(element int) int {

	if al.IsEmpty() {
		return -1
	}

	for i := 0; i < al.Size(); i++ {
		if al.data[i] == element {
			return i
		}
	}
	return -1
}

func (al *ArrayList) findLast(element int) int {

	if al.IsEmpty() {
		return -1
	}

	for i := al.Size() - 1; i > 0; i-- {
		if al.data[i] == element {
			return i
		}
	}
	return -1
}

func (al *ArrayList) shiftLeft(index int) {

	al.checkAndDecreaseLimit()

	for i := index; i < al.Size(); i++ {
		al.data[i] = al.data[i+1]
	}

	al.size--
}

//TODO: Improve the logic for filtering the arrayList
func (al *ArrayList) filterArrayList(retain bool, elements ...int) {
	cache := map[int]bool{}
	temp := make([]int, al.capacity)
	j := 0

	for _, e := range elements {
		cache[e] = true
	}

	for i := 0; i < al.Size(); i++ {
		if cache[al.data[i]] {
			if retain {
				temp[j] = al.data[i]
				j++
			}
		} else {
			if !retain {
				temp[j] = al.data[i]
				j++
			}
		}
	}
	al.data = temp
	al.size = j
}

func resize(capacity int, data []int) []int {
	temp := make([]int, capacity)

	sz := len(data)
	if capacity < sz {
		sz = capacity
	}

	for i := 0; i < sz; i++ {
		temp[i] = data[i]
	}
	return temp
}

func newArrayListIterator(al *ArrayList) *arrayListIterator {
	return &arrayListIterator{
		currentIndex: 0,
		al:           al,
	}
}
