package list

import (
	"fmt"
	"github.com/rewantsoni/go-datastructures/iterator"
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
	al           *ArrayList
}

func NewArrayList(elements ...int) *ArrayList {
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

func newArrayListIterator(al *ArrayList) *arrayListIterator {
	return &arrayListIterator{
		currentIndex: 0,
		al:           al,
	}
}

func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) IsEmpty() bool {
	return al.size == 0
}

func (al *ArrayList) Add(element int) bool {
	return al.AddAt(al.size, element)
}

func (al *ArrayList) AddAll(elements ...int) bool {
	return addAll(al.size, al, elements...)
}

func (al *ArrayList) AddAt(pos int, element int) bool {
	return addAll(pos, al, element)
}

func (al *ArrayList) GetAt(i int) (int, bool) {
	if al.IsEmpty() || i < 0 || i >= al.size {
		return -1, false
	}

	return al.data[i], true
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

func (al *ArrayList) IndexOf(element int) int {
	return find(al, element)
}

func (al *ArrayList) Replace(old int, new int) bool {
	i := al.IndexOf(old)
	if i == -1 {
		return false
	}

	al.data[i] = new
	return true
}

func (al *ArrayList) Set(i int, new int) bool {
	if al.IsEmpty() || i < 0 || i >= al.size {
		return false
	}

	al.data[i] = new
	return true
}

func (al *ArrayList) Remove(element int) bool {
	i := al.IndexOf(element)
	if i == -1 {
		return false
	}

	shiftLeft(al, i)

	return true
}

func (al *ArrayList) RemoveAt(i int) (int, bool) {
	if al.IsEmpty() || i < 0 || i >= al.size {
		return -1, false
	}

	e := al.data[i]
	shiftLeft(al, i)

	return e, true
}

func (al *ArrayList) RetainAll(elements ...int) List {

	temp := NewArrayList(elements...)
	for i := 0; i < al.size; i++ {
		if !temp.Contains(al.data[i]) {
			al.RemoveAt(i)
			i--
		}
	}
	return al
}

func (al *ArrayList) ReplaceAll(f func(e int) int) {
	for i := 0; i < al.size; i++ {
		al.data[i] = f(al.data[i])
	}
}

func (al *ArrayList) Iterator() iterator.Iterator {
	return newArrayListIterator(al)
}

func (al *ArrayList) String() string {
	sb := strings.Builder{}

	for i := 0; i < al.size; i++ {
		sb.WriteString(fmt.Sprintf("%d ", al.data[i]))
	}

	return sb.String()
}

func (ali *arrayListIterator) HasNext() bool {
	return ali.currentIndex < ali.al.size
}

func (ali *arrayListIterator) Next() interface{} {
	if !ali.HasNext() {
		return nil
	}
	//TODO: error?
	e, err := ali.al.GetAt(ali.currentIndex)
	if !err {
		return nil
	}
	ali.currentIndex++
	return e
}

func checkAndIncreaseLimit(al *ArrayList) {
	if al.size >= int(float64(al.capacity)*al.upperLoadFactor) {
		al.capacity *= al.scalingFactor
		al.data = resize(al.capacity, al.data)
	}
}

func checkAndDecreaseLimit(al *ArrayList) {
	if al.size <= int(float64(al.capacity)*al.lowerLoadFactor) && al.capacity != initialCapacity {
		al.capacity /= al.scalingFactor
		al.data = resize(al.capacity, al.data)
	}
}

func addAll(pos int, al *ArrayList, data ...int) bool {

	for i, d := range data {
		if !add(pos+i, al, d) {
			return false
		}
	}
	return true
}

func add(pos int, al *ArrayList, e int) bool {

	if !(pos >= 0 && pos <= al.size) {
		return false
	}

	checkAndIncreaseLimit(al)

	for i := al.Size(); i > pos; i-- {
		al.data[i] = al.data[i-1]
	}

	al.data[pos] = e
	al.size++

	return true
}

func find(al *ArrayList, val int) int {

	if al.IsEmpty() {
		return -1
	}

	for i := 0; i < al.size; i++ {
		if al.data[i] == val {
			return i
		}
	}
	return -1
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

func shiftLeft(al *ArrayList, index int) {

	checkAndDecreaseLimit(al)

	for i := index; i < al.size; i++ {
		al.data[i] = al.data[i+1]
	}

	al.size--
}
