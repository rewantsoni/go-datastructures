package list

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewLinkedList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test create new empty linked list",
			actualResult: func() List {
				return NewLinkedList()
			},
			expectedResult: func() List {
				ll := &LinkedList{size: 0, first: nil, last: nil}
				return ll
			},
		},
		{
			name: "test create new linked list with elements",
			actualResult: func() List {
				return NewLinkedList(1, 2, 3, 4, 5)
			},
			expectedResult: func() List {
				ll := &LinkedList{size: 5}
				ll.first, ll.last = testCreateNodes(1, 2, 3, 4, 5)
				return ll
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test size on creating an new linked list",
			actualResult: func() int {
				ll := NewLinkedList()
				return ll.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test size on creating an new linked list with elements",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.Size()
			},
			expectedResult: 5,
		},
		{
			name: "test size on creating an new linked list with 100 elements",
			actualResult: func() int {
				data := make([]int, 100)
				for i := 0; i < 100; i++ {
					data[i] = i
				}
				ll := NewLinkedList(data...)
				return ll.Size()
			},
			expectedResult: 100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, size)
		})
	}
}

func TestLinkedListClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult List
	}{
		{
			name: "test clear on creating an new Linked list",
			actualResult: func() List {
				ll := NewLinkedList()
				ll.Clear()
				return ll
			},
			expectedResult: NewLinkedList(),
		},
		{
			name: "test clear on creating an new Linked list with elements",
			actualResult: func() List {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				ll.Clear()
				return ll
			},
			expectedResult: NewLinkedList(),
		},
		{
			name: "test clear on creating an new Linked list with 100 elements",
			actualResult: func() List {
				data := make([]int, 100)
				for i := 0; i < 100; i++ {
					data[i] = i
				}
				ll := NewLinkedList(data...)
				ll.Clear()
				return ll
			},
			expectedResult: NewLinkedList(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, size)
		})
	}
}

func TestLinkedListAdd(t *testing.T) {
	testCases := []struct {
		name               string
		actualResult       func() (int, bool, List)
		expectedLinkedList func() List
		expectedResult     bool
		expectedSize       int
	}{
		{
			name: "test Size is 1 after adding one element",
			actualResult: func() (int, bool, List) {
				ll := NewLinkedList()
				res := ll.Add(1)
				return ll.Size(), res, ll
			},
			expectedSize: 1,
			expectedLinkedList: func() List {
				ll := NewLinkedList(1)
				return ll
			},
			expectedResult: true,
		},
		{
			name: "test Size is 2 after adding two element",
			actualResult: func() (int, bool, List) {
				ll := NewLinkedList()

				res := ll.Add(1)

				res = ll.Add(2)
				return ll.Size(), res, ll
			},
			expectedSize: 2,
			expectedLinkedList: func() List {
				ll := NewLinkedList(1, 2)
				return ll
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size, res, resLinkedList := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedSize, size)
			assert.Equal(t, testCase.expectedLinkedList(), resLinkedList)
		})
	}
}

func TestLinkedListAddAll(t *testing.T) {
	testCases := []struct {
		name               string
		actualResult       func() (int, bool, List)
		expectedLinkedList func() List
		expectedResult     bool
		expectedSize       int
	}{
		{
			name: "test Size is 5 after adding five element",
			actualResult: func() (int, bool, List) {
				ll := NewLinkedList()
				res := ll.AddAll(1, 2, 3, 4, 5)
				return ll.Size(), res, ll
			},
			expectedSize: 5,
			expectedLinkedList: func() List {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll
			},
			expectedResult: true,
		},
		{
			name: "test Size is 17 after adding seventeen element",
			actualResult: func() (int, bool, List) {
				ll := NewLinkedList()
				data := make([]int, 17)
				for i := 0; i < 17; i++ {
					data[i] = i
				}
				res := ll.AddAll(data...)
				return ll.Size(), res, ll
			},
			expectedSize: 17,
			expectedLinkedList: func() List {
				ll := NewLinkedList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)
				return ll
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			size, res, resLinkedList := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedSize, size)
			assert.Equal(t, testCase.expectedLinkedList(), resLinkedList)
		})
	}
}

func TestLinkedListAddAt(t *testing.T) {
	testCases := []struct {
		name               string
		actualResult       func() (bool, List)
		expectedLinkedList func() List
		expectedResult     bool
	}{
		{
			name: "test addAt index 10 when list is empty",
			actualResult: func() (bool, List) {
				ll := NewLinkedList()
				res := ll.AddAt(10, 1)
				return res, ll
			},
			expectedLinkedList: func() List {
				return NewLinkedList()
			},
			expectedResult: false,
		},
		{
			name: "test addAt index 0 when list is empty",
			actualResult: func() (bool, List) {
				ll := NewLinkedList()
				res := ll.AddAt(0, 1)
				return res, ll
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1)
			},
			expectedResult: true,
		},
		{
			name: "test addAt is first index when list is not empty",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3)

				res := ll.AddAt(0, 10)
				return res, ll
			},
			expectedLinkedList: func() List {
				return NewLinkedList(10, 1, 2, 3)
			},
			expectedResult: true,
		},
		{
			name: "test addAt is last index when list is not empty",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3)

				res := ll.AddAt(3, 10)
				return res, ll
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2, 3, 10)
			},
			expectedResult: true,
		},
		{
			name: "test addAt is last index+1 when list is not empty",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3)

				res := ll.AddAt(4, 10)
				return res, ll
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2, 3)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, resLinkedList := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedLinkedList(), resLinkedList)
		})
	}
}

func TestLinkedListGetAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectPanic    bool
	}{
		{
			name: "test get when list is empty",
			actualResult: func() int {
				ll := NewLinkedList()
				fmt.Println("al:")
				return ll.GetAt(7)
			},
			expectPanic: true,
		},
		{
			name: "test get first index when list has elements",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.GetAt(0)
			},
			expectedResult: 1,
		},
		{
			name: "test get last index when list has elements",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.GetAt(4)
			},
			expectedResult: 5,
		},
		{
			name: "test get when index is out of bond and positive",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.GetAt(5)
			},
			expectPanic: true,
		},
		{
			name: "test get when index is out of bond and negative",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.GetAt(-1)
			},
			expectPanic: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != testCase.expectPanic {
					t.Errorf("GetAt() paniced when it didn't expect to panic")
				}
			}()

			res := testCase.actualResult()
			if !testCase.expectPanic {
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestLinkedListContains(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test when list is empty",
			actualResult: func() bool {
				ll := NewLinkedList()
				return ll.Contains(1)
			},
			expectedResult: false,
		},
		{
			name: "test when list contains the element",
			actualResult: func() bool {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "test when list does not contain the element",
			actualResult: func() bool {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.Contains(10)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test when list is empty",
			actualResult: func() bool {
				ll := NewLinkedList()
				return ll.ContainsAll(1, 2, 3)
			},
			expectedResult: false,
		},
		{
			name: "test when list contains all the elements",
			actualResult: func() bool {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.ContainsAll(1, 2, 5)
			},
			expectedResult: true,
		},
		{
			name: "test when list does not contain the elements",
			actualResult: func() bool {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.ContainsAll(10, 1)
			},
			expectedResult: false,
		},
		{
			name: "test when list does not contain the elements",
			actualResult: func() bool {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.ContainsAll(1, 3, 10)
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectedError  error
	}{
		{
			name: "test indexOf when list is empty",
			actualResult: func() int {
				ll := NewLinkedList()
				return ll.IndexOf(1)
			},
			expectedResult: -1,
			expectedError:  errors.New("invalid operation: list is empty"),
		},
		{
			name: "test indexOf when element at first index",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.IndexOf(1)
			},
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name: "test indexOf when element at last index",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.IndexOf(5)
			},
			expectedResult: 4,
			expectedError:  nil,
		},
		{
			name: "test indexOf when element not in list",
			actualResult: func() int {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.IndexOf(6)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 6 not found in the LinkedList"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, bool)
		expectedResult func() List
		expectedBool   bool
	}{
		{
			name: "test replace when LinkedList is empty",
			actualResult: func() (List, bool) {
				ll := NewLinkedList()
				err := ll.Replace(1, 4)
				return ll, err
			},
			expectedResult: func() List {
				ll := NewLinkedList()
				return ll
			},
			expectedBool: false,
		},
		{
			name: "test replace when LinkedList has elements",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				res := ll.Replace(1, 4)
				return ll, res
			},
			expectedResult: func() List {
				return NewLinkedList(4, 2, 3, 4, 5)
			},
			expectedBool: true,
		},
		{
			name: "test replace when LinkedList has same element more than once",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(1, 2, 1, 4, 5)
				res := ll.Replace(1, 4)
				return ll, res
			},
			expectedResult: func() List {
				return NewLinkedList(4, 2, 4, 4, 5)
			},
			expectedBool: true,
		},
		{
			name: "test replace when LinkedList does not have element",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(2, 2, 3, 4, 5)
				res := ll.Replace(1, 4)
				return ll, res
			},
			expectedResult: func() List {
				return NewLinkedList(2, 2, 3, 4, 5)
			},
			expectedBool: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, b := testCase.actualResult()
			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedResult(), res)
		})

	}
}

func TestLinkedListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, bool)
		expectedResult func() List
		expectedBool   bool
	}{
		{
			name: "test Set when LinkedList is empty",
			actualResult: func() (List, bool) {
				ll := NewLinkedList()
				res := ll.Set(1, 4)
				return ll, res
			},
			expectedBool: false,
			expectedResult: func() List {
				return NewLinkedList()
			},
		},
		{
			name: "test Set when LinkedList has elements",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				res := ll.Set(1, 4)
				return ll, res
			},
			expectedBool: true,
			expectedResult: func() List {
				return NewLinkedList(1, 4, 3, 4, 5)
			},
		},
		{
			name: "test Set when index is out of bond",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				res := ll.Set(10, 4)
				return ll, res
			},
			expectedBool: false,
			expectedResult: func() List {
				return NewLinkedList(1, 2, 3, 4, 5)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, b := testCase.actualResult()
			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedResult(), res)
		})

	}
}

func TestLinkedListRemove(t *testing.T) {
	testCases := []struct {
		name               string
		actualResult       func() (List, bool)
		expectedResult     bool
		expectedLinkedList func() List
	}{
		{
			name: "test remove element when list is empty",
			actualResult: func() (List, bool) {
				ll := NewLinkedList()

				res := ll.Remove(1)
				return ll, res
			},
			expectedLinkedList: func() List {
				return NewLinkedList()
			},
			expectedResult: false,
		},
		{
			name: "test remove first occurrence of element",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(1, 2, 3, 1, 5)

				res := ll.Remove(1)
				return ll, res
			},
			expectedLinkedList: func() List {
				ll := NewLinkedList(2, 3, 1, 5)
				return ll
			},
			expectedResult: true,
		},
		{
			name: "test remove when occurrence of element in middle",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(4, 2, 3, 1, 5)

				res := ll.Remove(1)
				return ll, res
			},
			expectedLinkedList: func() List {
				ll := NewLinkedList(4, 2, 3, 5)
				return ll
			},
			expectedResult: true,
		},
		{
			name: "test remove when occurrence of element at end",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(4, 2, 3, 1, 5)

				res := ll.Remove(5)
				return ll, res
			},
			expectedLinkedList: func() List {
				ll := NewLinkedList(4, 2, 3, 1)
				return ll
			},
			expectedResult: true,
		},
		{
			name: "test remove element when element not present",
			actualResult: func() (List, bool) {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				res := ll.Remove(6)
				return ll, res
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2, 3, 4, 5)
			},
			expectedResult: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ll, res := testCase.actualResult()
			assert.Equal(t, testCase.expectedLinkedList(), ll)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListRemoveAt(t *testing.T) {
	testCases := []struct {
		name               string
		actualResult       func() (List, int, bool)
		expectedResult     int
		expectedLinkedList func() List
		expectedBool       bool
	}{
		{
			name: "test removeAt element when list is empty",
			actualResult: func() (List, int, bool) {
				ll := NewLinkedList()

				res, err := ll.RemoveAt(1)
				return ll, res, err
			},
			expectedLinkedList: func() List {
				return NewLinkedList()
			},
			expectedBool:   false,
			expectedResult: -1,
		},
		{
			name: "test removeAt index when list has elements",
			actualResult: func() (List, int, bool) {
				ll := NewLinkedList(1, 2, 3, 1, 5)

				res, err := ll.RemoveAt(1)
				return ll, res, err
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 3, 1, 5)
			},
			expectedBool:   true,
			expectedResult: 2,
		},
		{
			name: "test removeAt index when list has elements at first",
			actualResult: func() (List, int, bool) {
				ll := NewLinkedList(1, 2, 3, 1, 5)

				res, err := ll.RemoveAt(0)
				return ll, res, err
			},
			expectedLinkedList: func() List {
				return NewLinkedList(2, 3, 1, 5)
			},
			expectedBool:   true,
			expectedResult: 1,
		},
		{
			name: "test removeAt index when list has elements at last",
			actualResult: func() (List, int, bool) {
				ll := NewLinkedList(1, 2, 3, 1, 5)

				res, err := ll.RemoveAt(4)
				return ll, res, err
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2, 3, 1)
			},
			expectedBool:   true,
			expectedResult: 5,
		},
		{
			name: "test removeAt element when index not present",
			actualResult: func() (List, int, bool) {
				ll := NewLinkedList(1, 2, 3, 4, 5)

				res, err := ll.RemoveAt(6)
				return ll, res, err
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2, 3, 4, 5)
			},
			expectedBool:   false,
			expectedResult: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ll, res, b := testCase.actualResult()
			assert.Equal(t, testCase.expectedBool, b)
			assert.Equal(t, testCase.expectedLinkedList(), ll)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListReplaceAll(t *testing.T) {

	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult func() List
	}{
		{
			name: "test replace all when Linked is empty",
			actualResult: func() List {
				ll := NewLinkedList()
				ll.ReplaceAll(testMultiply{Val: 2})
				return ll
			},
			expectedResult: func() List {
				return NewLinkedList()
			},
		},
		{
			name: "test replace all when function multiply and Linked not empty",
			actualResult: func() List {
				ll := NewLinkedList(1, 2, 3)
				ll.ReplaceAll(testMultiply{Val: 2})
				return ll
			},
			expectedResult: func() List {
				return NewLinkedList(2, 4, 6)
			},
		},
		{
			name: "test replace all when add function and Linked not empty",
			actualResult: func() List {
				ll := NewLinkedList(1, 2, 3)
				ll.ReplaceAll(testAdd{Val: 10})
				return ll
			},
			expectedResult: func() List {
				return NewLinkedList(11, 12, 13)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListSubList(t *testing.T) {
	testCases := []struct {
		name               string
		actualResult       func() (bool, List)
		expectedLinkedList func() List
		expectedResult     bool
	}{
		{
			name: "test sublist when start index is more than end index",
			actualResult: func() (bool, List) {
				ll := NewLinkedList()
				res, tempList := ll.SubList(12, 10)
				return res, tempList
			},
			expectedLinkedList: func() List {
				return nil
			},
			expectedResult: false,
		},
		{
			name: "test sublist when list is empty and start is out of bound",
			actualResult: func() (bool, List) {
				ll := NewLinkedList()
				res, tempList := ll.SubList(2, 7)
				return res, tempList
			},
			expectedLinkedList: func() List {
				return nil
			},
			expectedResult: false,
		},
		{
			name: "test sublist when list is empty and not in bound",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2)
				res, tempList := ll.SubList(2, 2)
				return res, tempList
			},
			expectedLinkedList: func() List {
				return nil
			},
			expectedResult: false,
		},
		{
			name: "test sublist when list end is out of bound",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3)
				res, tempList := ll.SubList(2, 4)
				return res, tempList
			},
			expectedLinkedList: func() List {
				return nil
			},
			expectedResult: false,
		},
		{
			name: "test sublist when list is not empty",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3, 4)
				res, tempList := ll.SubList(0, 2)
				return res, tempList
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2)
			},
			expectedResult: true,
		},
		{
			name: "test sublist when list is not empty and complete list",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3, 4)
				res, tempList := ll.SubList(0, 4)
				return res, tempList
			},
			expectedLinkedList: func() List {
				return NewLinkedList(1, 2, 3, 4)
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, resLinkedList := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
			expectedLinkedList := testCase.expectedLinkedList()
			if res == true {
				for i := 0; i < expectedLinkedList.Size(); i++ {
					assert.Equal(t, expectedLinkedList.GetAt(i), resLinkedList.GetAt(i))
				}
			}
		})
	}
}

func TestLinkedListClone(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, List)
		expectedResult func() (bool, List)
	}{
		{
			name: "test clear on creating an new array list",
			actualResult: func() (bool, List) {
				ll := NewLinkedList()
				return ll.Clone()
			},
			expectedResult: func() (bool, List) {
				return true, NewLinkedList()
			},
		},
		{
			name: "test clear on creating an new linked list with elements",
			actualResult: func() (bool, List) {
				ll := NewLinkedList(1, 2, 3, 4, 5)
				return ll.Clone()
			},
			expectedResult: func() (bool, List) {
				return true, NewLinkedList(1, 2, 3, 4, 5)
			},
		},
		{
			name: "test clear on creating an new linked list with 100 elements",
			actualResult: func() (bool, List) {
				data := make([]int, 100)
				for i := 0; i < 100; i++ {
					data[i] = i
				}
				ll := NewLinkedList(data...)
				return ll.Clone()
			},
			expectedResult: func() (bool, List) {
				data := make([]int, 100)
				for i := 0; i < 100; i++ {
					data[i] = i
				}
				return true, NewLinkedList(data...)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, list := testCase.actualResult()
			expectedRes, expectedList := testCase.expectedResult()
			assert.Equal(t, expectedRes, res)
			assert.Equal(t, expectedList, list)
		})
	}
}
