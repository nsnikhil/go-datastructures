package stack

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestCreateNewStack(t *testing.T) {
	actualResult := NewStack[int64]()

	expectedResult := &Stack[int64]{list.NewLinkedList[int64]()}

	assert.Equal(t, expectedResult, actualResult)
}

func TestStackPush(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() *Stack[int64]
		expectedResult func() *Stack[int64]
		expectedError  error
	}{
		{
			name: "should insert an element into stack",
			actualResult: func() *Stack[int64] {
				s := NewStack[int64]()
				s.Push(1)

				return s
			},
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64](1)

				return &Stack[int64]{ll}
			},
		},
		{
			name: "should insert math.MaxInt8 elements into stack",
			actualResult: func() *Stack[int64] {
				s := NewStack[int64]()

				data := internal.SliceGenerator{Size: math.MaxInt8}.Generate()
				for _, e := range data {
					s.Push(e)
				}

				return s
			},
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64](internal.SliceGenerator{Size: math.MaxInt8, Reverse: true}.Generate()...)

				return &Stack[int64]{ll}
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

func TestStackPop(t *testing.T) {
	testCases := []struct {
		name             string
		actualResult     func() (error, []int64, *Stack[int64])
		expectedElements []int64
		expectedResult   func() *Stack[int64]
		expectedError    error
	}{
		{
			name: "should pop one element from stack to make the stack empty",
			actualResult: func() (error, []int64, *Stack[int64]) {
				s := NewStack[int64]()

				s.Push(1)

				e, err := s.Pop()

				return err, []int64{e}, s
			},
			expectedElements: []int64{1},
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64]()

				return &Stack[int64]{ll}
			},
		},
		{
			name: "should pop one element from stack",
			actualResult: func() (error, []int64, *Stack[int64]) {
				s := NewStack[int64]()
				s.Push(1)
				s.Push(2)

				e, err := s.Pop()

				return err, []int64{e}, s
			},
			expectedElements: []int64{2},
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64](1)

				return &Stack[int64]{ll}
			},
		},
		{
			name: "should pop four elements from stack",
			actualResult: func() (error, []int64, *Stack[int64]) {
				s := NewStack[int64]()

				data := internal.SliceGenerator{Size: 5}.Generate()
				for _, e := range data {
					s.Push(e)
				}

				var res []int64

				for i := 0; i < 4; i++ {
					e, err := s.Pop()
					//TODO: IS ASSERTION HERE THE CORRECT THING?
					require.NoError(t, err)
					res = append(res, e)
				}

				return nil, res, s
			},
			expectedElements: []int64{4, 3, 2, 1},
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64](0)

				return &Stack[int64]{ll}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, e, s := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElements, e)
			assert.Equal(t, testCase.expectedResult(), s)
		})
	}
}

func TestStackPeek(t *testing.T) {
	testCases := []struct {
		name             string
		actualResult     func() (error, int64, *Stack[int64])
		expectedElements int64
		expectedResult   func() *Stack[int64]
		expectedError    error
	}{
		{
			name: "should peek top element from stack of size 1",
			actualResult: func() (error, int64, *Stack[int64]) {
				s := NewStack[int64]()
				s.Push(1)

				e, err := s.Peek()

				return err, e, s
			},
			expectedElements: 1,
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64](1)

				return &Stack[int64]{ll}
			},
		},
		{
			name: "should peek first element from Stack of size 2",
			actualResult: func() (error, int64, *Stack[int64]) {
				s := NewStack[int64]()

				s.Push(1)
				s.Push(2)

				e, err := s.Peek()

				return err, e, s
			},
			expectedElements: 2,
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64](2, 1)

				return &Stack[int64]{ll}
			},
		},
		{
			name: "should return error when stack is empty",
			actualResult: func() (error, int64, *Stack[int64]) {
				s := NewStack[int64]()

				e, err := s.Peek()

				return err, e, s
			},
			expectedError: errors.New("stack is empty"),
			expectedResult: func() *Stack[int64] {
				ll := list.NewLinkedList[int64]()
				return &Stack[int64]{ll}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, e, s := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElements, e)
			assert.Equal(t, testCase.expectedResult(), s)
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
			name: "should return true when stack is Empty",
			actualResult: func() bool {
				s := NewStack[int64]()

				return s.Empty()
			},
			expectedResult: true,
		},
		{
			name: "should return false when stack is not Empty",
			actualResult: func() bool {
				s := NewStack[int64]()

				s.Push(1)

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

func TestStackSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "should return size 0 when stack is empty",
			actualResult: func() int64 {
				s := NewStack[int64]()

				s.Push(1)

				return s.Size()
			},
			expectedResult: 1,
		},
		{
			name: "should return size 1 when stack has one element",
			actualResult: func() int64 {
				s := NewStack[int64]()

				s.Push(1)
				s.Push(2)

				return s.Size()
			},
			expectedResult: 2,
		},
		{
			name: "should return size MaxInt8 when stack has MaxInt8 element",
			actualResult: func() int64 {
				s := NewStack[int64]()

				data := internal.SliceGenerator{Size: math.MaxInt8}.Generate()
				for _, e := range data {
					s.Push(e)
				}

				return s.Size()
			},
			expectedResult: math.MaxInt8,
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
		name         string
		actualResult func() *Stack[int64]
	}{
		{
			name: "should clear an empty stack",
			actualResult: func() *Stack[int64] {
				s := NewStack[int64]()

				s.Clear()

				return s
			},
		},
		{
			name: "should clear an stack with elements",
			actualResult: func() *Stack[int64] {
				s := NewStack[int64]()
				s.Push(1)

				s.Clear()

				return s
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, NewStack[int64](), testCase.actualResult())
		})
	}
}
