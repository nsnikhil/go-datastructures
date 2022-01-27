package list

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestCreateNewLinkedList(t *testing.T) {
	testCases := map[string]struct {
		data           []int64
		expectedResult func() *LinkedList[int64]
		expectedError  error
	}{
		"should create new empty linked list": {
			data:           make([]int64, 0),
			expectedResult: func() *LinkedList[int64] { return &LinkedList[int64]{size: 0} },
		},
		"should create new linked list with values": {
			data: []int64{1, 2, 3, 4},
			expectedResult: func() *LinkedList[int64] {
				l := &LinkedList[int64]{size: 4}
				l.first, l.last = createNodes[int64](1, 2, 3, 4)
				return l
			},
		},
		"should create new linked list with var args": {
			data: internal.SliceGenerator{Size: math.MaxInt8}.Generate(),
			expectedResult: func() *LinkedList[int64] {
				l := &LinkedList[int64]{size: math.MaxInt8}
				l.first, l.last = createNodes[int64](internal.SliceGenerator{Size: math.MaxInt8}.Generate()...)
				return l
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), NewLinkedList(testCase.data...))
		})
	}

}

func TestLinkedListAdd(t *testing.T) {
	addFunc := func(t *testing.T, ll *LinkedList[int64], data []int64) {
		for _, e := range data {
			ll.Add(e)
		}
	}

	testCases := map[string]struct {
		input          func() (*LinkedList[int64], []int64)
		expectedResult *LinkedList[int64]
	}{
		"should add item to new list": {
			input: func() (*LinkedList[int64], []int64) {
				ll := NewLinkedList[int64]()

				return ll, []int64{1}
			},
			expectedResult: NewLinkedList[int64](1),
		},
		"should add item to non empty list": {
			input: func() (*LinkedList[int64], []int64) {
				ll := NewLinkedList[int64](1)

				return ll, []int64{2}
			},
			expectedResult: NewLinkedList[int64](1, 2),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ll, data := testCase.input()
			addFunc(t, ll, data)
			assert.Equal(t, testCase.expectedResult, ll)
		})
	}
}

func TestLinkedListGetSize(t *testing.T) {
	testCases := map[string]struct {
		data         []int64
		expectedSize int64
	}{
		"should return 1 for list with single element": {
			data:         []int64{1},
			expectedSize: 1,
		},
		"should return 4 for list with four element": {
			data:         []int64{1, 2, 3, 4},
			expectedSize: 4,
		},
		"should return 0 for empty list": {
			data:         []int64{},
			expectedSize: 0,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			data := testCase.data
			ll := NewLinkedList[int64](data...)

			assert.Equal(t, testCase.expectedSize, ll.Size())
		})
	}
}

func TestLinkedListIterator(t *testing.T) {
	ll := NewLinkedList[int64](1, 2, 3, 4)

	it := ll.Iterator()

	var res []int64
	for it.HasNext() {
		v, _ := it.Next()
		res = append(res, v)
	}

	assert.Equal(t, []int64{1, 2, 3, 4}, res)
}

func TestLinkedListDecrementIterator(t *testing.T) {
	ll := NewLinkedList[int64](1, 2, 3, 4)

	it := ll.DescendingIterator()

	var res []int64

	for it.HasNext() {
		v, _ := it.Next()
		res = append(res, v)
	}

	assert.Equal(t, []int64{4, 3, 2, 1}, res)
}

func TestSortLinkedList(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() List[int]
		expectedResult List[int]
	}{
		"should sort integer List": {
			actualResult: func() List[int] {
				ll := NewLinkedList[int](5, 4, 3, 2, 1)

				ll.Sort(comparator.NewIntegerComparator())

				return ll
			},
			expectedResult: NewLinkedList[int](1, 2, 3, 4, 5),
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestLinkedListAddAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List[int64])
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should add element at index mid",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 3, 4, 5)

				return ll.AddAt(1, 2), ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4, 5),
		},
		{
			name: "should add element at index start",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.AddAt(0, 0), ll
			},
			expectedResult: NewLinkedList[int64](0, 1, 2, 3, 4),
		},
		{
			name: "should add element at index end",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.AddAt(4, 5), ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4, 5),
		},
		{
			name: "should add elements at beginning of empty linked list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64]()

				return ll.AddAt(0, 1), ll
			},
			expectedResult: NewLinkedList[int64](1),
		},
		{
			name: "should return error when adding element at invalid index",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.AddAt(5, 5), ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
			expectedError:  errors.New("invalid index 5"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, l := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, l)
		})
	}
}

func TestLinkedListAddFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List[int64])
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should add element to start of list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](2, 3, 4)

				return ll.AddFirst(1), ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "should add element to start of new list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64]()

				return ll.AddFirst(1), ll
			},
			expectedResult: NewLinkedList[int64](1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListAddLast(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List[int64])
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "test add element to end of list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 2, 3)

				return ll.AddLast(4), ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "test add element to end of new list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64]()

				return ll.AddLast(1), ll
			},
			expectedResult: NewLinkedList[int64](1),
		},
		{
			name: "test add element to start of empty list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 2)

				ll.Clear()

				return ll.AddLast(1), ll
			},
			expectedResult: NewLinkedList[int64](1),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListAllAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int64]
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should add all elements to a new list",
			actualResult: func() List[int64] {
				ll := NewLinkedList[int64]()

				ll.AddAll(1, 2, 3, 4)

				return ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "should add all to a list",
			actualResult: func() List[int64] {
				ll := NewLinkedList[int64](1, 2)

				ll.AddAll(3, 4)

				return ll
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "should add all not fail when arguments are empty",
			actualResult: func() List[int64] {
				ll := NewLinkedList[int64](1, 2)

				ll.AddAll()

				return ll
			},
			expectedResult: NewLinkedList[int64](1, 2),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, l)
		})
	}
}

func TestLinkedListClear(t *testing.T) {
	testCases := []struct {
		name   string
		result func() *LinkedList[int64]
	}{
		{
			name: "should clear list",
			result: func() *LinkedList[int64] {
				ll := NewLinkedList[int64](1, 2)

				ll.Clear()

				return ll
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.result()
			assert.Nil(t, res.first)
			assert.Nil(t, res.last)
			assert.Equal(t, int64(0), res.size)
		})
	}
}

func TestLinkedListClone(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int64]
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should clone integer linked list",
			actualResult: func() List[int64] {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.Clone()
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "should clone empty linked list",
			actualResult: func() List[int64] {
				ll := NewLinkedList[int64]()

				return ll.Clone()
			},
			expectedResult: NewLinkedList[int64](),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
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
			name: "should return true when element is present",
			actualResult: func() bool {
				ll := NewLinkedList[int64](1, 2)

				return ll.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "should return false when element is not present",
			actualResult: func() bool {
				ll := NewLinkedList[int64](1, 2)

				return ll.Contains(0)
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

func TestLinkedListGet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int64, error)
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should return error for empty List",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64]()

				return ll.Get(0)
			},
			expectedError: errors.New("list is empty"),
		},
		{
			name: "should return first element from the List",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64](1)

				return ll.Get(0)
			},
			expectedResult: 1,
		},
		{
			name: "should return 4th element from the List",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64](0, 1, 2, 3)

				return ll.Get(3)
			},
			expectedResult: 3,
		},
		{
			name: "should return error for invalid index",
			actualResult: func() (int64, error) {
				list := NewLinkedList[int64](1)

				return list.Get(1)
			},
			expectedError: errors.New("invalid index 1"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListGetFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int64, error)
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should get the first element",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.GetFirst()
			},
			expectedResult: 1,
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64]()

				return ll.GetFirst()
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListGetLast(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int64, error)
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should return last element",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.GetLast()
			},
			expectedResult: 4,
		},
		{
			name: "return return error when list is empty",
			actualResult: func() (int64, error) {
				ll := NewLinkedList[int64]()

				return ll.GetLast()
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
		expectedError  error
	}{
		{
			name: "should return true when all elements are present",
			actualResult: func() bool {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6)

				return ll.ContainsAll(6, 1, 3)
			},
			expectedResult: true,
		},
		{
			name: "should return false when all elements are not present",
			actualResult: func() bool {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6)

				return ll.ContainsAll(6, 1, 0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
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
		actualResult   func() int64
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should return index when element is found",
			actualResult: func() int64 {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.IndexOf(2)
			},
			expectedResult: 1,
		},
		{
			name: "should return -1 when element is not present",
			actualResult: func() int64 {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				return ll.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() int64 {
				ll := NewLinkedList[int64]()

				return ll.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListIsEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "should return true when list is empty",
			actualResult: func() bool {
				ll := NewLinkedList[int64]()

				return ll.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "should return false when list is not empty",
			actualResult: func() bool {
				ll := NewLinkedList[int64](1)

				return ll.IsEmpty()
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

func TestLinkedListLastIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should return last index of the element in list",
			actualResult: func() int64 {
				ll := NewLinkedList[int64](1, 2, 3, 1, 4)

				return ll.LastIndexOf(1)
			},
			expectedResult: 3,
		},
		{
			name: "should return -1 when the element in not present",
			actualResult: func() int64 {
				ll := NewLinkedList[int64](1, 2, 3, 1, 4)

				return ll.LastIndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListRemove(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], error)
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should successfully remove element",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.Remove(2)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 3, 4),
		},
		{
			name: "should successfully remove first element",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.Remove(1)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 3, 4),
		},
		{
			name: "should successfully remove last element",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.Remove(4)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3),
		},
		{
			name: "should return error when trying to remove element which is not present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.Remove(0)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
			expectedError:  errors.New("element 0 not found in the list"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListRemoveAt(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (List[int64], int64, error)
		expectedResult  List[int64]
		expectedElement int64
		expectedError   error
	}{
		{
			name: "should remove element at index 1",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				e, err := ll.RemoveAt(1)

				return ll, e, err
			},
			expectedResult:  NewLinkedList[int64](1, 3, 4),
			expectedElement: 2,
		},
		{
			name: "should remove element at index 0",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				e, err := ll.RemoveAt(0)

				return ll, e, err
			},
			expectedResult:  NewLinkedList[int64](2, 3, 4),
			expectedElement: 1,
		},
		{
			name: "should remove element at index 3",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				e, err := ll.RemoveAt(3)

				return ll, e, err
			},
			expectedResult:  NewLinkedList[int64](1, 2, 3),
			expectedElement: 4,
		},
		{
			name: "should remove element at index 0 when list only contains one element",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1)

				e, err := ll.RemoveAt(0)
				assert.Equal(t, int64(1), e)

				return ll, e, err
			},
			expectedResult:  NewLinkedList[int64](),
			expectedElement: 1,
		},
		{
			name: "should fail to remove element at invalid index",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				e, err := ll.RemoveAt(4)

				return ll, e, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
			expectedError:  fmt.Errorf("invalid index %d", 4),
		},
		{
			name: "should fail to remove element for empty list",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64]()

				e, err := ll.RemoveAt(0)

				return ll, e, err
			},
			expectedResult: NewLinkedList[int64](),
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, e, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedElement, e)
		})
	}
}

func TestLinkedListRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], error)
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should successfully remove elements",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveAll(2, 4)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 3),
		},
		{
			name: "should successfully remove only elements which are present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveAll(2, 4, 5)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 3),
		},
		{
			name: "test successfully remove when list has only one element",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1)

				err := ll.RemoveAll(1)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](),
		},
		{
			name: "should successfully remove elements at start",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveAll(1)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 3, 4),
		},
		{
			name: "should successfully remove elements at end",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveAll(4)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3),
		},
		{
			name: "should remove all keeps all element when argument list is empty",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveAll()

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "should remove all fails when list is empty",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64]()

				err := ll.RemoveAll(1, 2)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](),
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListRemoveFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], int64, error)
		expectedList   List[int64]
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should get and remove first element",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ele, err := ll.RemoveFirst()

				return ll, ele, err
			},
			expectedResult: 1,
			expectedList:   NewLinkedList[int64](2, 3, 4),
		},
		{
			name: "should get and remove first element when list contains one element",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1)

				ele, err := ll.RemoveFirst()

				return ll, ele, err
			},
			expectedResult: 1,
			expectedList:   NewLinkedList[int64](),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64]()

				ele, err := ll.RemoveFirst()

				return ll, ele, err
			},
			expectedList:  NewLinkedList[int64](),
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l, res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList, l)
		})
	}
}

func TestLinkedListRemoveLast(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], int64, error)
		expectedList   List[int64]
		expectedResult int64
		expectedError  error
	}{
		{
			name: "should get and remove last element",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ele, err := ll.RemoveLast()

				return ll, ele, err
			},
			expectedResult: 4,
			expectedList:   NewLinkedList[int64](1, 2, 3),
		},
		{
			name: "should get and remove last element when list contains one element",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1)

				ele, err := ll.RemoveLast()

				return ll, ele, err
			},
			expectedResult: 1,
			expectedList:   NewLinkedList[int64](),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64]()

				ele, err := ll.RemoveLast()

				return ll, ele, err
			},
			expectedList:  NewLinkedList[int64](),
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l, res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList, l)
		})
	}
}

func TestLinkedListRemoveFirstOccurrence(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], error)
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should remove first occurrence of 1 when multiple occurrence of 1 is present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 1, 3, 1)

				err := ll.RemoveFirstOccurrence(1)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 1, 3, 1),
		},
		{
			name: "should remove first occurrence of 1 when single occurrence of 1 is present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveFirstOccurrence(1)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 3, 4),
		},
		{
			name: "should return error when element is not present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RemoveFirstOccurrence(5)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
			expectedError:  errors.New("element 5 not found in the list"),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64]()

				err := ll.RemoveFirstOccurrence(1)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](),
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListRemoveLastOccurrence(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], error)
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should remove last occurrence of 1 when multiple occurrence of 1 is present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 1, 3, 1)

				ok, err := ll.RemoveLastOccurrence(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 1, 3),
		},
		{
			name: "should remove last occurrence of 1 when single occurrence of 1 is present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ok, err := ll.RemoveLastOccurrence(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 3, 4),
		},
		{
			name: "should return error when element is not present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ok, err := ll.RemoveLastOccurrence(5)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4),
			expectedError:  errors.New("element 5 not found in the list"),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64]()

				ok, err := ll.RemoveLastOccurrence(1)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](),
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List[int64])
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should replace a given value with new one",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1)

				return ll.Replace(1, 2), ll
			},
			expectedResult: NewLinkedList[int64](2),
		},
		{
			name: "should return error when item is not found in the list",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64](1, 2)

				return ll.Replace(5, 3), ll
			},
			expectedResult: NewLinkedList[int64](1, 2),
			expectedError:  errors.New("element 5 not found in the list"),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (error, List[int64]) {
				ll := NewLinkedList[int64]()

				return ll.Replace(1, 2), ll
			},
			expectedResult: NewLinkedList[int64](),
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListReplaceAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int]
		expectedResult List[int]
	}{
		{
			name: "should replace all on integer List with increment operator",
			actualResult: func() List[int] {
				ll := NewLinkedList[int](1, 2, 3, 4)

				ll.ReplaceAll(testIntIncOperator{})

				return ll
			},
			expectedResult: NewLinkedList[int](2, 3, 4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], error)
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should retain all from integer List",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RetainAll(2, 4)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 4),
		},
		{
			name: "should retain all from integer list with ignoring element which are not present",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				err := ll.RetainAll(2, 4, 5)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](2, 4),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64]()

				err := ll.RetainAll(1, 2)

				return ll, err
			},
			expectedResult: NewLinkedList[int64](),
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestLinkedListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], int64, error)
		expectedResult int64
		expectedList   List[int64]
		expectedError  error
	}{
		{
			name: "should set value at index 3",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ele, err := ll.Set(3, 5)

				return ll, ele, err
			},
			expectedResult: 5,
			expectedList:   NewLinkedList[int64](1, 2, 3, 5),
		},
		{
			name: "should set value at index 0",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ele, err := ll.Set(0, 2)

				return ll, ele, err
			},
			expectedResult: 2,
			expectedList:   NewLinkedList[int64](2, 2, 3, 4),
		},
		{
			name: "should set value at index 1",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ele, err := ll.Set(1, 4)

				return ll, ele, err
			},
			expectedResult: 4,
			expectedList:   NewLinkedList[int64](1, 4, 3, 4),
		},
		{
			name: "should return error due to invalid index",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64](1, 2, 3, 4)

				ele, err := ll.Set(5, 3)

				return ll, ele, err
			},
			expectedError: fmt.Errorf("invalid index %d", 5),
			expectedList:  NewLinkedList[int64](1, 2, 3, 4),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int64], int64, error) {
				ll := NewLinkedList[int64]()

				ele, err := ll.Set(0, 1)

				return ll, ele, err
			},
			expectedError: errors.New("list is empty"),
			expectedList:  NewLinkedList[int64](),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l, res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList, l)
		})
	}
}

func TestLinkedListSubList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], error)
		expectedResult List[int64]
		expectedError  error
	}{
		{
			name: "should get sublist from index 1 to 4",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(1, 4)
			},
			expectedResult: NewLinkedList[int64](2, 3, 4, 5),
		},
		{
			name: "should get sublist from index 0 to 0",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(0, 0)
			},
			expectedResult: NewLinkedList[int64](1),
		},
		{
			name: "should get sublist from index 0 to 4",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(0, 4)
			},
			expectedResult: NewLinkedList[int64](1, 2, 3, 4, 5),
		},
		{
			name: "should get sublist from index 4 to 6",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(4, 6)
			},
			expectedResult: NewLinkedList[int64](5, 6, 7),
		},
		{
			name: "should return error for invalid start index",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(-1, 4)
			},
			expectedError: fmt.Errorf("invalid index %d", -1),
		},
		{
			name: "should return error for invalid end index",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(0, 10)
			},
			expectedError: fmt.Errorf("invalid index %d", 10),
		},
		{
			name: "should return error when end is less than start",
			actualResult: func() (List[int64], error) {
				ll := NewLinkedList[int64](1, 2, 3, 4, 5, 6, 7)

				return ll.SubList(4, 2)
			},
			expectedError: errors.New("end cannot be smaller than start"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func createNodes[T comparable](data ...T) (*node[T], *node[T]) {
	var first, curr, prev *node[T]

	for _, e := range data {
		curr = newNode(e)

		if first == nil {
			first = curr
		}

		if prev != nil {
			curr.prev = prev
			prev.next = curr
		}

		prev = curr
	}

	return first, curr
}
