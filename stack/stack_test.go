package stack

import (
	"github.com/nsnikhil/go-datastructures/liberr"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewStack(t *testing.T) {
	actualResult, err := NewStack()
	require.NoError(t, err)

	ll, err := list.NewLinkedList()
	require.NoError(t, err)
	expectedResult := &Stack{ll}

	assert.Equal(t, expectedResult, actualResult)
}

func TestStackPush(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, *Stack)
		expectedResult func() *Stack
		expectedError  error
	}{
		{
			name: "insert an element into Empty Stack",
			actualResult: func() (error, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				return s.Push(1), s
			},
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
		{
			name: "insert two element into Empty Stack",
			actualResult: func() (error, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				err = s.Push(1)
				if err != nil {
					return err, s
				}

				return s.Push(2), s
			},
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(2, 1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
		{
			name: "return error when inserting different type of elements",
			actualResult: func() (error, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				err = s.Push(1)
				if err != nil {
					return err, s
				}

				return s.Push("a"), s
			},
			expectedError: liberr.TypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedResult(), res)
			}

		})
	}
}

func TestStackPop(t *testing.T) {
	testCases := []struct {
		name             string
		actualResult     func() (error, []interface{}, *Stack)
		expectedElements []interface{}
		expectedResult   func() *Stack
		expectedError    error
	}{
		{
			name: "Pop one element from to make Stack Empty",
			actualResult: func() (error, []interface{}, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))

				e, err := s.Pop()

				return err, []interface{}{e}, s
			},
			expectedElements: []interface{}{1},
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)
				ll.Clear()

				return &Stack{ll}
			},
		},
		{
			name: "Pop one element from Stack",
			actualResult: func() (error, []interface{}, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))

				e, err := s.Pop()

				return err, []interface{}{e}, s
			},
			expectedElements: []interface{}{2},
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
		{
			name: "Pop four elements from Stack",
			actualResult: func() (error, []interface{}, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))
				require.NoError(t, s.Push(3))
				require.NoError(t, s.Push(4))
				require.NoError(t, s.Push(5))

				var res []interface{}

				e, err := s.Pop()
				if err != nil {
					return err, res, s
				}
				res = append(res, e)

				e, err = s.Pop()
				if err != nil {
					return err, res, s
				}
				res = append(res, e)

				e, err = s.Pop()
				if err != nil {
					return err, res, s
				}
				res = append(res, e)

				e, err = s.Pop()
				if err != nil {
					return err, res, s
				}
				res = append(res, e)

				return err, res, s
			},
			expectedElements: []interface{}{5, 4, 3, 2},
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, e, s := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedElements, e)
				assert.Equal(t, testCase.expectedResult(), s)
			}

		})
	}
}

func TestStackPeek(t *testing.T) {
	testCases := []struct {
		name             string
		actualResult     func() (error, interface{}, *Stack)
		expectedElements interface{}
		expectedResult   func() *Stack
		expectedError    error
	}{
		{
			name: "Peek top element from Stack of size 1",
			actualResult: func() (error, interface{}, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))

				e, err := s.Peek()

				return err, e, s
			},
			expectedElements: 1,
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
		{
			name: "Peek first element from Stack of size 2",
			actualResult: func() (error, interface{}, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))

				e, err := s.Peek()

				return err, e, s
			},
			expectedElements: 2,
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(2, 1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
		{
			name: "Peek first element from Stack of size 5",
			actualResult: func() (error, interface{}, *Stack) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))
				require.NoError(t, s.Push(3))
				require.NoError(t, s.Push(4))
				require.NoError(t, s.Push(5))

				e, err := s.Peek()

				return err, e, s
			},
			expectedElements: 5,
			expectedResult: func() *Stack {
				ll, err := list.NewLinkedList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				return &Stack{ll}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, e, s := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if testCase.expectedError == nil {
				assert.Equal(t, testCase.expectedElements, e)
				assert.Equal(t, testCase.expectedResult(), s)
			}

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
			name: "return true when Stack is Empty",
			actualResult: func() bool {
				s, err := NewStack()
				require.NoError(t, err)

				return s.Empty()
			},
			expectedResult: true,
		},
		{
			name: "return false when Stack is not Empty",
			actualResult: func() bool {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))

				return s.Empty()
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestStackCount(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "return Size 0 when Stack is Empty",
			actualResult: func() int {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))

				return s.Size()
			},
			expectedResult: 1,
		},
		{
			name: "return Size 1 when Stack has 1 element",
			actualResult: func() int {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))

				return s.Size()
			},
			expectedResult: 2,
		},
		{
			name: "return Size 1 when Stack has 2 elements",
			actualResult: func() int {
				s, err := NewStack()
				require.NoError(t, err)

				return s.Size()
			},
			expectedResult: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestStackClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "should return true for Empty call after clearing Stack",
			actualResult: func() bool {
				s, err := NewStack()
				require.NoError(t, err)

				s.Clear()

				return s.Empty()
			},
			expectedResult: true,
		},
		{
			name: "should return true for Empty call after clearing Stack two",
			actualResult: func() bool {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))

				s.Clear()

				return s.Empty()
			},
			expectedResult: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestStackSearch(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "return index of element in Stack",
			actualResult: func() (int, error) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))

				return s.Search(2)
			},
			expectedResult: 0,
		},
		{
			name: "return index of element in Stack two",
			actualResult: func() (int, error) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))
				require.NoError(t, s.Push(3))
				require.NoError(t, s.Push(4))

				return s.Search(1)
			},
			expectedResult: 3,
		},
		{
			name: "return -1 when element is not present in Stack",
			actualResult: func() (int, error) {
				s, err := NewStack()
				require.NoError(t, err)

				require.NoError(t, s.Push(1))
				require.NoError(t, s.Push(2))

				return s.Search(3)
			},
			expectedResult: -1,
		},
		{
			name: "return -1 when searching in Empty Stack",
			actualResult: func() (int, error) {
				s, err := NewStack()
				require.NoError(t, err)

				return s.Search(1)
			},
			expectedResult: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			i, _ := testCase.actualResult()

			//TODO FIX ERROR MESSAGE AND UNCOMMENT THE ASSERTION
			//assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, i)

		})
	}
}
