package list

import (
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