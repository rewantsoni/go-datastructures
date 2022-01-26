package queue

import (
	"github.com/rewantsoni/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewLinkedListQueue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue
		expectedResult func() Queue
	}{
		{
			name: "test linked list queue create new empty linked list queue",
			actualResult: func() Queue {
				return NewLinkedListQueue()
			},
			expectedResult: func() Queue {
				ll := list.NewLinkedList()
				return &LinkedListQueue{
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

func TestLinkedListQueueClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Queue
		expectedResult func() Queue
	}{
		{
			name: "test linked list queue clear with empty linked list queue",
			actualResult: func() Queue {
				return NewLinkedListQueue()
			},
			expectedResult: func() Queue {
				ll := list.NewLinkedList()
				return &LinkedListQueue{
					ll: ll,
				}
			},
		},
		{
			name: "test linked list queue clear after adding elements",
			actualResult: func() Queue {
				s := NewLinkedListQueue()
				s.Enqueue(1)
				s.Clear()
				return s
			},
			expectedResult: func() Queue {
				return NewLinkedListQueue()
			},
		},
		{
			name: "test linked list queue clear after adding and removing elements elements",
			actualResult: func() Queue {
				s := NewLinkedListQueue()
				s.Enqueue(1)
				s.Enqueue(2)
				s.Dequeue()
				s.Clear()
				return s
			},
			expectedResult: func() Queue {
				return NewLinkedListQueue()
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

func TestLinkedListQueueEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test is empty on new queue",
			actualResult: func() bool {
				q := NewLinkedListQueue()
				return q.Empty()
			},
			expectedResult: true,
		},
		{
			name: "test is empty on queue after adding elements",
			actualResult: func() bool {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				q.Enqueue(2)
				return q.Empty()
			},
			expectedResult: false,
		},
		{
			name: "test is empty on queue after adding and removing elements",
			actualResult: func() bool {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				q.Dequeue()
				return q.Empty()
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

func TestLinkedListQueuePeek(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectedPanic  bool
	}{
		{
			name: "test linked list queue peek on empty linked list queue",
			actualResult: func() int {
				s := NewLinkedListQueue()
				return s.Peek()
			},
			expectedPanic: true,
		},
		{
			name: "test linked list queue peek on linked list queue after adding elements",
			actualResult: func() int {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				q.Enqueue(2)
				return q.Peek()
			},
			expectedResult: 1,
		},
		{
			name: "test peek on linked list queue after adding one element",
			actualResult: func() int {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				return q.Peek()
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

func TestLinkedListQueueDequeue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
		expectedPanic  bool
	}{
		{
			name: "test linked list queue dequeue on empty linked list queue",
			actualResult: func() int {
				q := NewLinkedListQueue()
				return q.Dequeue()
			},
			expectedPanic: true,
		},
		{
			name: "test linked list queue dequeue on linked list queue after adding elements",
			actualResult: func() int {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				q.Enqueue(2)
				return q.Dequeue()
			},
			expectedResult: 1,
		},
		{
			name: "test dequeue on linked list queue after adding one element",
			actualResult: func() int {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				return q.Dequeue()
			},
			expectedResult: 1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != testCase.expectedPanic {
					t.Errorf("Dequeue() paniced when it didn't expect to panic")
				}
			}()

			res := testCase.actualResult()
			if !testCase.expectedPanic {
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func TestLinkedListQueueEnqueue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test push on new linked list queue",
			actualResult: func() bool {
				q := NewLinkedListQueue()
				return q.Enqueue(2)
			},
			expectedResult: true,
		},
		{
			name: "test push on linked list queue after adding elements",
			actualResult: func() bool {
				q := NewLinkedListQueue()
				q.Enqueue(1)
				return q.Enqueue(2)
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

func TestNewLinkedListQueueSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test linked list queue size on empty linked list queue",
			actualResult: func() int {
				q := NewLinkedListQueue()
				return q.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test linked list queue size on linked list queue after adding elements",
			actualResult: func() int {
				s := NewLinkedListQueue()
				s.Enqueue(3)
				s.Enqueue(4)
				return s.Size()
			},
			expectedResult: 2,
		},
		{
			name: "test size on linked list queue after adding one element",
			actualResult: func() int {
				s := NewLinkedListQueue()
				s.Enqueue(1)
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
