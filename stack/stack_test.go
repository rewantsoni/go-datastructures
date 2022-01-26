package stack

import (
	"github.com/rewantsoni/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewStack(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *Stack
		expectedResult func() *Stack
	}{
		{
			name: "test stack create new empty stack",
			actualResult: func() *Stack {
				return NewStack()
			},
			expectedResult: func() *Stack {
				ll := list.NewLinkedList()
				return &Stack{
					ll: ll,
				}
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

func TestStackClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *Stack
		expectedResult func() *Stack
	}{
		{
			name: "test stack clear with empty stack",
			actualResult: func() *Stack {
				return NewStack()
			},
			expectedResult: func() *Stack {
				ll := list.NewLinkedList()
				return &Stack{
					ll: ll,
				}
			},
		},
		{
			name: "test stack clear after adding elements",
			actualResult: func() *Stack {
				s := NewStack()
				s.Push(1)
				s.Clear()
				return s
			},
			expectedResult: func() *Stack {
				return NewStack()
			},
		},
		{
			name: "test stack clear after adding and removing elements elements",
			actualResult: func() *Stack {
				s := NewStack()
				s.Push(1)
				s.Push(2)
				s.Pop()
				s.Clear()
				return s
			},
			expectedResult: func() *Stack {
				return NewStack()
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

func TestStackEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test is empty on new stack",
			actualResult: func() bool {
				s := NewStack()
				return s.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test is empty on stack after adding elements",
			actualResult: func() bool {
				s := NewStack()
				s.Push(1)
				s.Push(2)
				return s.Empty()
			},
			expectedResult: false,
		},
		{
			name: "test is empty on stack after adding elements",
			actualResult: func() bool {
				s := NewStack()
				s.Push(1)
				s.Pop()
				return s.Empty()
			},
			expectedResult: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestStackPeek(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectedPanic  bool
	}{
		{
			name: "test stack peek on empty stack",
			actualResult: func() int {
				s := NewStack()
				return s.Peek()
			},
			expectedPanic: true,
		},
		{
			name: "test stack peek on stack after adding elements",
			actualResult: func() int {
				s := NewStack()
				s.Push(1)
				s.Push(2)
				return s.Peek()
			},
			expectedResult: 2,
		},
		{
			name: "test peek on stack after adding one element",
			actualResult: func() int {
				s := NewStack()
				s.Push(1)
				return s.Peek()
			},
			expectedResult: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != testCase.expectedPanic {
					t.Errorf("Peek() paniced when it didn't expect to panic")
				}
			}()

			res := testCase.actualResult()
			if !testCase.expectedPanic {
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectedPanic  bool
	}{
		{
			name: "test stack pop on empty stack",
			actualResult: func() int {
				s := NewStack()
				return s.Pop()
			},
			expectedPanic: true,
		},
		{
			name: "test stack pop on stack after adding elements",
			actualResult: func() int {
				s := NewStack()
				s.Push(1)
				s.Push(2)
				return s.Pop()
			},
			expectedResult: 2,
		},
		{
			name: "test pop on stack after adding one element",
			actualResult: func() int {
				s := NewStack()
				s.Push(1)
				return s.Pop()
			},
			expectedResult: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != testCase.expectedPanic {
					t.Errorf("Peek() paniced when it didn't expect to panic")
				}
			}()

			res := testCase.actualResult()
			if !testCase.expectedPanic {
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestStackPush(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test push on new stack",
			actualResult: func() bool {
				s := NewStack()
				return s.Push(2)
			},
			expectedResult: true,
		},
		{
			name: "test push on stack after adding elements",
			actualResult: func() bool {
				s := NewStack()
				s.Push(1)
				return s.Push(2)
			},
			expectedResult: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestStackSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test stack size on empty stack",
			actualResult: func() int {
				s := NewStack()
				return s.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test stack size on stack after adding elements",
			actualResult: func() int {
				s := NewStack()
				s.Push(3)
				s.Push(4)
				return s.Size()
			},
			expectedResult: 2,
		},
		{
			name: "test size on stack after adding one element",
			actualResult: func() int {
				s := NewStack()
				s.Push(1)
				return s.Size()
			},
			expectedResult: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
