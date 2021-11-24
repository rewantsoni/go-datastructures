package list

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/rewantsoni/go-datastructures/errors"
	"github.com/rewantsoni/go-datastructures/utils"
)

type ArrayList struct {
	arrayListType   string
	capacity        int
	upperLoadFactor float64
	lowerLoadFactor float64
	scalingFactor   int
	size            int
	data            []interface{}
}

func NewArrayList(elements ...interface{}) (*ArrayList, error) {
	al := &ArrayList{
		arrayListType:   utils.NA,
		size:            nought,
		capacity:        initialCapacity,
		upperLoadFactor: upperLoadFactor,
		lowerLoadFactor: lowerLoadFactor,
		scalingFactor:   scalingFactor,
		data:            make([]interface{}, initialCapacity)}

	if len(elements) == 0 {
		return al, nil
	}

	if err := al.AddAll(elements...); err != nil {
		return nil, err
	}

	return al, nil
}

// Size returns the size of the array
func (al *ArrayList) Size() int {
	return al.size
}

func (al *ArrayList) isEmpty() bool {
	if al.size == 0 {
		return true
	}
	return false
}

// Add adds an element to the arraylist
func (al *ArrayList) Add(element interface{}) error {
	return al.AddAll(element)
}

// AddAll adds multiple element to the arraylist
func (al *ArrayList) AddAll(elements ...interface{}) error {
	return addAll(al, elements...)
}

func (al *ArrayList) AddAt(pos int, element interface{}) error {
	if al.isEmpty() {
		if pos != 0 {
			return errors.EmptyListError
		}
		err := al.Add(element)
		if err != nil {
			return err
		}
		return nil
	}

	elementType := utils.GetTypeName(element)
	if elementType != al.arrayListType {
		return errors.TypeMismatchError(al.arrayListType, elementType)
	}

	if pos >= 0 && pos <= al.size+1 {
		for i := al.size + 1; i > pos; i-- {
			al.data[i] = al.data[i-1]
		}
		al.data[pos] = element
		al.size++
		checkAndIncreaseLimit(al)
		return nil
	}

	return errors.IndexOutOfBoundError(pos)
}

// GetAt returns the element at an index
func (al *ArrayList) GetAt(i int) (interface{}, error) {
	if al.isEmpty() {
		return nil, errors.EmptyListError
	}

	if isIndexInRange(al.size, i) {
		return al.data[i], nil
	}

	return nil, errors.IndexOutOfBoundError(i)
}

// Contains returns if element is present in the arraylist
func (al *ArrayList) Contains(element interface{}) (bool, error) {
	if al.isEmpty() {
		return false, errors.EmptyListError
	}
	elementType := utils.GetTypeName(element)
	if al.arrayListType != elementType {
		return false, errors.TypeMismatchError(al.arrayListType, elementType)
	}
	for _, d := range al.data {
		if cmp.Equal(d, element) {
			return true, nil
		}
	}
	return false, nil
}

// IndexOf returns the first index of an element if present along with an error
func (al *ArrayList) IndexOf(element interface{}) (int, error) {
	return find(al, element)
}

// LastIndexOf returns the last index of an element if present along with an error
func (al *ArrayList) LastIndexOf(element interface{}) (int, error) {
	return findLast(al, element)
}

// FindAllOccurrencesOf returns all the occurrences of an element
func (al *ArrayList) FindAllOccurrencesOf(element interface{}) ([]int, error) {
	return findAll(al, element)
}

// Replace updates the first occurrence of an element with the new one
func (al *ArrayList) Replace(old interface{}, new interface{}) error {
	i, err := al.IndexOf(old)
	if err != nil {
		return err
	}

	newType := utils.GetTypeName(new)
	if newType != al.arrayListType {
		return errors.TypeMismatchError(al.arrayListType, newType)
	}

	al.data[i] = new
	return nil
}

// ReplaceFromEnd updates the last occurrence of an element with the new one
func (al *ArrayList) ReplaceFromEnd(old interface{}, new interface{}) error {
	i, err := al.LastIndexOf(old)
	if err != nil {
		return err
	}

	newType := utils.GetTypeName(new)
	if newType != al.arrayListType {
		return errors.TypeMismatchError(al.arrayListType, newType)
	}

	al.data[i] = new
	return nil
}

// ReplaceAllOccurrencesOf updates all the occurrence of an element with the new one
func (al *ArrayList) ReplaceAllOccurrencesOf(old interface{}, new interface{}) error {
	indexes, err := al.FindAllOccurrencesOf(old)
	if err != nil {
		return err
	}

	newType := utils.GetTypeName(new)
	if newType != al.arrayListType {
		return errors.TypeMismatchError(al.arrayListType, newType)
	}

	for _, i := range indexes {
		al.data[i] = new
	}
	return nil
}

// Set updates the value at a particular index
func (al *ArrayList) Set(i int, new interface{}) error {
	if al.isEmpty() {
		return errors.EmptyListError
	}

	newType := utils.GetTypeName(new)
	if newType != al.arrayListType {
		return errors.TypeMismatchError(al.arrayListType, newType)
	}

	if isIndexInRange(al.size, i) {
		al.data[i] = new
		return nil
	}

	return errors.IndexOutOfBoundError(i)
}

// Remove removes the first occurrence of the element
func (al *ArrayList) Remove(element interface{}) (bool, error) {
	i, err := al.IndexOf(element)
	if err != nil {
		return false, err
	}

	shiftLeft(al, i)
	return true, nil
}

// RemoveFromEnd removes the last occurrence of the element
func (al *ArrayList) RemoveFromEnd(element interface{}) (bool, error) {
	i, err := al.LastIndexOf(element)
	if err != nil {
		return false, err
	}

	shiftLeft(al, i)
	return true, nil
}

// RemoveAt removes the element at an index
func (al *ArrayList) RemoveAt(i int) (interface{}, error) {
	if al.isEmpty() {
		return nil, errors.EmptyListError
	}

	if isIndexInRange(al.size, i) {
		e := al.data[i]
		shiftLeft(al, i)
		return e, nil
	}
	return nil, errors.IndexOutOfBoundError(i)
}


func (al *ArrayList) Print() {
	fmt.Println("Array Type: ", al.arrayListType)
	fmt.Println("Array Size: ", al.size)
	fmt.Println("Array Capacity: ", al.capacity)
	fmt.Println("Printing Array: ")
	for i := 0; i < al.size; i++ {
		fmt.Print(al.data[i], ", ")
	}
}
