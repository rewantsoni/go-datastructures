package list

import (
	"github.com/google/go-cmp/cmp"
	"github.com/rewantsoni/go-datastructures/errors"
	"github.com/rewantsoni/go-datastructures/utils"
)

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
func addAll(al *ArrayList, data ...interface{}) error {

	for _, d := range data {
		if err := add(al, d); err != nil {
			return err
		}
	}
	return nil
}

func add(al *ArrayList, e interface{}) error {

	elementType := utils.GetTypeName(e)

	if al.isEmpty() && al.arrayListType == utils.NA {
		al.arrayListType = elementType
	} else if al.arrayListType != elementType {
		return errors.TypeMismatchError(al.arrayListType, elementType)
	}

	checkAndIncreaseLimit(al)

	al.data[al.size] = e
	al.size++

	return nil
}

func find(al *ArrayList, val interface{}) (int, error) {

	elementType := utils.GetTypeName(val)
	if al.isEmpty() {
		return -1, errors.EmptyListError
	} else if al.arrayListType != elementType {
		return -1, errors.TypeMismatchError(al.arrayListType, elementType)
	}
	for i := 0; i < al.size; i++ {
		if cmp.Equal(al.data[i], val) {
			return i, nil
		}
	}
	return -1, errors.NotFondError(val, "arrayList")
}

func findLast(al *ArrayList, val interface{}) (int, error) {

	elementType := utils.GetTypeName(val)
	if al.isEmpty() {
		return -1, errors.EmptyListError
	}else if al.arrayListType != elementType {
		return -1, errors.TypeMismatchError(al.arrayListType, elementType)
	}
	for i := al.size - 1; i >= 0; i-- {
		if cmp.Equal(al.data[i], val) {
			return i, nil
		}
	}
	return -1, errors.NotFondError(val, "arrayList")
}

func findAll(al *ArrayList, val interface{}) ([]int, error) {
	if al.isEmpty() {
		return nil, errors.EmptyListError
	}

	elementType := utils.GetTypeName(val)
	if al.arrayListType != elementType {
		return nil, errors.TypeMismatchError(al.arrayListType, elementType)
	}

	indexes, err := NewArrayList()
	if err != nil {
		return nil, err
	}

	for i, d := range al.data {
		if cmp.Equal(d, val) {
			err := indexes.Add(i)
			if err != nil {
				return nil, err
			}
		}
	}
	if indexes.isEmpty() {
		return nil, errors.NotFondError(val, "arrayList")
	}
	return convertToInt(indexes), nil
}

func convertToInt(a *ArrayList) []int {
	b := make([]int, a.size)
	for i := 0; i < a.size; i++ {
		b[i] = a.data[i].(int)

	}
	return b
}

func resize(capacity int, data []interface{}) []interface{} {
	temp := make([]interface{}, capacity)

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
	for i := index; i < al.size; i++ {
		al.data[i] = al.data[i+1]
	}
	al.size--
	checkAndDecreaseLimit(al)
}

func isIndexInRange(size, i int) bool {
	return i >= 0 && i < size
}
