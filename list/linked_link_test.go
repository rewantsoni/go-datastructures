package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewLinkedList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult *LinkedList
		expectedError  error
	}{
		{
			name: "test create new empty linked list",
			actualResult: func() List {
				return NewLinkedList()
			},
			expectedResult: &LinkedList{
				size:  0,
				first: nil,
			},
		},
		{
			name: "test create new array list with elements",
			actualResult: func() List {
				return NewLinkedList(1, 2, 3, 4, 5)
			},
			expectedResult: &LinkedList{
				size:  5,
				first: &node{data: 1, next: &node{data: 2, next: &node{data: 3, next: &node{data: 4, next: &node{data: 5, next: nil}}}}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
