package list

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/functions/predicate"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestCreateNewArrayList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int64]
		expectedResult List[int64]
	}{
		{
			name: "should create a empty array list",
			actualResult: func() List[int64] {
				return NewArrayList[int64]()
			},
			expectedResult: &ArrayList[int64]{
				factors: &factors{upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2, capacity: 16},
				data:    make([]int64, 16),
			},
		},
		{
			name: "should create array list with elements",
			actualResult: func() List[int64] {
				return NewArrayList[int64](1, 2, 3, 4, 5)
			},
			expectedResult: &ArrayList[int64]{
				size:    5,
				factors: &factors{upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2, capacity: 16},
				data:    []int64{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			name: "should create array list with variable arg elements",
			actualResult: func() List[int64] {
				return NewArrayList[int64](internal.SliceGenerator{Size: math.MaxInt8}.Generate()...)
			},
			expectedResult: &ArrayList[int64]{
				size:    int64(math.MaxInt8),
				factors: &factors{upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2, capacity: 256},
				data:    append(internal.SliceGenerator{Size: math.MaxInt8}.Generate(), internal.SliceGenerator{Size: math.MaxInt8 + 2, AbsoluteValue: 0, Absolute: true}.Generate()...),
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

func TestArrayListSize(t *testing.T) {
	testCases := []struct {
		name         string
		actualSize   func() int64
		expectedSize int64
	}{
		{
			name: "should return the size as 1 when listen contains single element",
			actualSize: func() int64 {
				list := NewArrayList[int]()
				list.Add(1)

				return list.Size()
			},
			expectedSize: 1,
		},
		{
			name: "should return size as 0 for empty list",
			actualSize: func() int64 {
				list := NewArrayList[any]()

				return list.Size()
			},
			expectedSize: 0,
		},

		{
			name: "should return the size as math.MaxInt16 when listen contains math.MaxInt16 elements",
			actualSize: func() int64 {
				list := NewArrayList[int64](internal.SliceGenerator{Size: math.MaxInt16}.Generate()...)

				return list.Size()
			},
			expectedSize: math.MaxInt16,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectedSize, testCase.actualSize())
	}
}

func TestArrayListAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int64, List[int64])
		expectedResult List[int64]
		expectedSize   int64
	}{
		{
			name: "should return the size as 1 when listen contains single element",
			actualResult: func() (int64, List[int64]) {
				al := NewArrayList[int64]()
				al.Add(1)

				return al.Size(), al
			},
			expectedSize:   1,
			expectedResult: NewArrayList[int64](1),
		},
		{
			name: "should return the size as math.MaxInt16 when listen contains math.MaxInt16 elements",
			actualResult: func() (int64, List[int64]) {
				al := NewArrayList[int64]()
				for i := int64(0); i < math.MaxInt16; i++ {
					al.Add(i)
				}

				return al.Size(), al
			},
			expectedSize:   math.MaxInt16,
			expectedResult: NewArrayList[int64](internal.SliceGenerator{Size: math.MaxInt16}.Generate()...),
		},
	}

	for _, testCase := range testCases {
		size, res := testCase.actualResult()

		assert.Equal(t, testCase.expectedSize, size)
		assert.Equal(t, testCase.expectedResult, res)
	}
}

func TestArrayListFilter(t *testing.T) {
	testCases := map[string]struct {
		inputFilter    predicate.Predicate[int64]
		inputList      List[int64]
		expectedResult List[int64]
	}{
		"should filter even numbers": {
			inputFilter:    evenFilter{},
			inputList:      NewArrayList[int64](internal.SliceGenerator{Size: 10}.Generate()...),
			expectedResult: NewArrayList[int64](0, 2, 4, 6, 8),
		},
		"should filter no elements when all elements matches filter": {
			inputFilter:    evenFilter{},
			inputList:      NewArrayList[int64](0, 2, 4, 6, 8),
			expectedResult: NewArrayList[int64](0, 2, 4, 6, 8),
		},
		"should filter all elements when no element matches filter": {
			inputFilter:    evenFilter{},
			inputList:      NewArrayList[int64](1, 3, 5, 7, 9),
			expectedResult: NewArrayList[int64](),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res := testCase.inputList.Filter(testCase.inputFilter)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListFindFirst(t *testing.T) {
	testCases := map[string]struct {
		inputFilter    predicate.Predicate[int64]
		inputList      List[int64]
		expectedResult int64
		expectedError  error
	}{
		"should return first even number": {
			inputFilter:    evenFilter{},
			inputList:      NewArrayList[int64](internal.SliceGenerator{Size: 10}.Generate()...),
			expectedResult: 0,
		},
		"should return error when no element matches filter": {
			inputFilter:   evenFilter{},
			inputList:     NewArrayList[int64](1, 3, 5, 7, 9),
			expectedError: errors.New("no element match the provided filter"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.inputList.FindFirst(testCase.inputFilter)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListGet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int64, error)
		expectedResult int64
	}{
		{
			name: "should return error when trying to get value from empty list",
			actualResult: func() (int64, error) {
				list := NewArrayList[int64]()

				return list.Get(0)
			},
		},
		{
			name: "should return first element from the list",
			actualResult: func() (int64, error) {
				list := NewArrayList[int64](1)

				return list.Get(0)
			},
			expectedResult: 1,
		},
		{
			name: "should return last element from the list",
			actualResult: func() (int64, error) {
				list := NewArrayList[int64](internal.SliceGenerator{Size: 4}.Generate()...)

				return list.Get(3)
			},
			expectedResult: 3,
		},
		{
			name: "should return error when trying to fetch value for invalid index",
			actualResult: func() (int64, error) {
				list := NewArrayList[int64](0)

				return list.Get(1)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			fmt.Println(err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListIteratorHasNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "should return false when list is empty",
			actualResult: func() bool {
				list := NewArrayList[int]()

				return list.Iterator().HasNext()
			},
		},
		{
			name: "should return true when list is not empty",
			actualResult: func() bool {
				list := NewArrayList[int](1)

				return list.Iterator().HasNext()
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

func TestArrayListIteratorNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() ([]int64, error)
		expectedResult []int64
	}{
		{
			name: "should return error for empty list",
			actualResult: func() ([]int64, error) {
				list := NewArrayList[int64]()

				res, err := list.Iterator().Next()
				return []int64{res}, err
			},
			expectedResult: []int64{0},
		},
		{
			name: "should return first element from the list",
			actualResult: func() ([]int64, error) {
				list := NewArrayList[int64](1)

				res, err := list.Iterator().Next()

				return []int64{res}, err
			},
			expectedResult: []int64{1},
		},
		{
			name: "should return all elements from the list",
			actualResult: func() ([]int64, error) {
				list := NewArrayList[int64](internal.SliceGenerator{Size: 4}.Generate()...)

				it := list.Iterator()

				var res []int64

				for it.HasNext() {
					v, _ := it.Next()
					res = append(res, v)
				}

				//TODO: CHANGE NIL
				return res, nil
			},
			expectedResult: []int64{0, 1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, _ := testCase.actualResult() //TODO: TEST ERROR
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListDescendingIteratorHasNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "should return false when list is empty",
			actualResult: func() bool {
				list := NewArrayList[int]()

				return list.DescendingIterator().HasNext()
			},
		},
		{
			name: "should return true when list is not empty",
			actualResult: func() bool {
				list := NewArrayList[int](1)

				return list.DescendingIterator().HasNext()
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

func TestArrayListDescendingIteratorNext(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() ([]int64, error)
		expectedResult []int64
	}{
		{
			name: "should return error for empty list",
			actualResult: func() ([]int64, error) {
				list := NewArrayList[int64]()

				res, err := list.DescendingIterator().Next()

				return []int64{res}, err
			},
			expectedResult: []int64{0},
		},
		{
			name: "should return first element from the list",
			actualResult: func() ([]int64, error) {
				list := NewArrayList[int64](1)

				res, err := list.DescendingIterator().Next()

				return []int64{res}, err
			},
			expectedResult: []int64{1},
		},
		{
			name: "should return all elements from the list",
			actualResult: func() ([]int64, error) {
				list := NewArrayList[int64](internal.SliceGenerator{Size: 4}.Generate()...)

				it := list.DescendingIterator()

				var res []int64

				for it.HasNext() {
					v, _ := it.Next()
					res = append(res, v)
				}

				//TODO: CHANGE NIL
				return res, nil
			},
			expectedResult: []int64{3, 2, 1, 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, _ := testCase.actualResult() //TODO: TEST ERROR
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error, List[int])
		expectedResult int
		expectedList   List[int]
		expectedError  error
	}{
		{
			name: "should set value at index 3",
			actualResult: func() (int, error, List[int]) {
				al := NewArrayList[int](1, 2, 3, 4)

				idx, err := al.Set(3, 5)

				return idx, err, al
			},
			expectedResult: 5,
			expectedList:   NewArrayList[int](1, 2, 3, 5),
		},
		{
			name: "should set value at index 0",
			actualResult: func() (int, error, List[int]) {
				al := NewArrayList[int](1, 2, 3, 4)

				idx, err := al.Set(0, 2)

				return idx, err, al
			},
			expectedResult: 2,
			expectedList:   NewArrayList[int](2, 2, 3, 4),
		},
		{
			name: "should fail to set value due to invalid index",
			actualResult: func() (int, error, List[int]) {
				al := NewArrayList[int](1, 2, 3, 4)

				idx, err := al.Set(5, 3)

				return idx, err, al
			},
			expectedError: fmt.Errorf("invalid index %d", 5),
			expectedList:  NewArrayList[int](1, 2, 3, 4),
		},
		{
			name: "should fail to set when list is empty",
			actualResult: func() (int, error, List[int]) {
				al := NewArrayList[int]()

				idx, err := al.Set(0, 3)

				return idx, err, al
			},
			expectedError: errors.New("list is empty"),
			expectedList:  NewArrayList[int](),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err, al := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList, al)
		})
	}
}

func TestArrayListSort(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int]
		expectedResult List[int]
	}{
		{
			name: "should sort integer list",
			actualResult: func() List[int] {
				al := NewArrayList[int](5, 4, 3, 2, 1)

				al.Sort(comparator.NewIntegerComparator())

				return al
			},
			expectedResult: NewArrayList[int](1, 2, 3, 4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestArrayListAddAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, int64, List[int])
		expectedResult List[int]
		expectedSize   int64
		expectedError  error
	}{
		{
			name: "should add element when list is empty",
			actualResult: func() (error, int64, List[int]) {
				al := NewArrayList[int]()

				err := al.AddAt(0, 1)

				return err, al.Size(), al
			},
			expectedResult: NewArrayList[int](1),
			expectedSize:   1,
		},
		{
			name: "should add element at index 1",
			actualResult: func() (error, int64, List[int]) {
				al := NewArrayList[int](1, 3, 4, 5)

				return al.AddAt(1, 2), al.Size(), al
			},
			expectedResult: NewArrayList[int](1, 2, 3, 4, 5),
			expectedSize:   5,
		},
		{
			name: "should add element at index 0",
			actualResult: func() (error, int64, List[int]) {
				al := NewArrayList[int](1, 2, 3, 4)

				return al.AddAt(0, 0), al.Size(), al
			},
			expectedResult: NewArrayList[int](0, 1, 2, 3, 4),
			expectedSize:   5,
		},
		{
			name: "should add element at index 3",
			actualResult: func() (error, int64, List[int]) {
				al := NewArrayList[int](1, 2, 3, 5)

				return al.AddAt(3, 4), al.Size(), al
			},
			expectedResult: NewArrayList[int](1, 2, 3, 4, 5),
			expectedSize:   5,
		},
		{
			name: "should return error when adding element at invalid index",
			actualResult: func() (error, int64, List[int]) {
				al := NewArrayList[int](1, 2, 3, 5)

				return al.AddAt(5, 8), al.Size(), al
			},
			expectedResult: NewArrayList[int](1, 2, 3, 5),
			expectedSize:   4,
			expectedError:  fmt.Errorf("invalid index %d", 5),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, sz, l := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, l)
		})
	}
}

func TestArrayListAddAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int64, List[int64])
		expectedResult List[int64]
		expectedSize   int64
	}{
		{
			name: "should add all elements to list",
			actualResult: func() (int64, List[int64]) {
				al := NewArrayList[int64](1, 2)
				al.AddAll(3, 4)

				return al.Size(), al
			},
			expectedResult: NewArrayList[int64](1, 2, 3, 4),
			expectedSize:   4,
		},
		{
			name: "should add all var args elements to list",
			actualResult: func() (int64, List[int64]) {
				al := NewArrayList[int64]()
				al.AddAll(internal.SliceGenerator{Size: 40}.Generate()...)

				return al.Size(), al
			},
			expectedResult: NewArrayList[int64](internal.SliceGenerator{Size: 40}.Generate()...),
			expectedSize:   40,
		},
		{
			name: "should add all not fail when args is empty",
			actualResult: func() (int64, List[int64]) {
				al := NewArrayList[int64](1, 2)

				al.AddAll()

				return al.Size(), al
			},
			expectedResult: NewArrayList[int64](1, 2),
			expectedSize:   2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sz, l := testCase.actualResult()

			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, l)
		})
	}
}

func TestArrayListClear(t *testing.T) {
	testCases := []struct {
		name         string
		actualResult func() (int64, []int64)
	}{
		{
			name: "should clear integer list",
			actualResult: func() (int64, []int64) {
				al := NewArrayList[int64](1, 2)
				al.Clear()

				return al.Size(), toSlice[int64](al)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sz, res := testCase.actualResult()

			assert.Equal(t, int64(0), sz)
			assert.Equal(t, []int64{}, res)
		})
	}
}

func TestArrayListClone(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int]
		expectedResult List[int]
		expectedError  error
	}{
		{
			name: "should clone list",
			actualResult: func() List[int] {
				al := NewArrayList[int](1, 2, 3, 4)

				return al.Clone()
			},
			expectedResult: NewArrayList[int](1, 2, 3, 4),
		},
		{
			name: "should clone empty array list",
			actualResult: func() List[int] {
				al := NewArrayList[int]()

				return al.Clone()
			},
			expectedResult: NewArrayList[int](),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListContains(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
		expectedError  error
	}{
		{
			name: "should return true when element is present",
			actualResult: func() bool {
				al := NewArrayList[int](1, 2)

				return al.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "should return false when element is not present",
			actualResult: func() bool {
				al := NewArrayList[int](1, 2)

				return al.Contains(0)
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

func TestArrayListContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
		expectedError  error
	}{
		{
			name: "should return true when all elements are present",
			actualResult: func() bool {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6)

				return al.ContainsAll(6, 1, 3)
			},
			expectedResult: true,
		},
		{
			name: "should return false when all elements are not present",
			actualResult: func() bool {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6)

				return al.ContainsAll(6, 1, 0)
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

func TestArrayListIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "should return index when element is present",
			actualResult: func() int64 {
				al := NewArrayList[int](1, 2, 3, 4)

				return al.IndexOf(2)
			},
			expectedResult: 1,
		},
		{
			name: "should return -1 when element is not present",
			actualResult: func() int64 {
				al := NewArrayList[int](1, 2, 3, 4)

				return al.IndexOf(0)
			},
			expectedResult: -1,
		},
		{
			name: "should return -1 when list is empty",
			actualResult: func() int64 {
				al := NewArrayList[int]()

				return al.IndexOf(0)
			},
			expectedResult: -1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "should return true when list is empty",
			actualResult: func() bool {
				al := NewArrayList[int]()

				return al.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "should return false when list is not empty",
			actualResult: func() bool {
				al := NewArrayList[int](1)

				return al.IsEmpty()
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

func TestArrayListLastIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "should return last index of the element in List",
			actualResult: func() int64 {
				al := NewArrayList[int](1, 2, 3, 1, 4)

				return al.LastIndexOf(1)
			},
			expectedResult: 3,
		},
		{
			name: "should return -1 when the element in not present",
			actualResult: func() int64 {
				al := NewArrayList[int](1, 2, 3, 1, 4)

				return al.LastIndexOf(0)
			},
			expectedResult: -1,
		},
		{
			name: "should return error when list is empty",
			actualResult: func() int64 {
				al := NewArrayList[int]()

				return al.LastIndexOf(0)
			},
			expectedResult: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveElement(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() ([]int64, int64, error)
		expectedResult []int64
		expectedSize   int64
		expectedError  error
	}{
		{
			name: "should successfully remove element",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.Remove(2)
				require.NoError(t, err)

				return toSlice[int64](al), al.Size(), nil
			},
			expectedResult: []int64{1, 3, 4},
			expectedSize:   3,
		},
		{
			name: "should return error when trying to remove element which is not present",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.Remove(0)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{1, 2, 3, 4},
			expectedSize:   4,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64]()

				err := al.Remove(0)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{},
			expectedSize:   0,
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, sz, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveAt(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() ([]int64, int64, int64, error)
		expectedResult  []int64
		expectedElement int64
		expectedSize    int64
		expectedError   error
	}{
		{
			name: "should remove element at index 1",
			actualResult: func() ([]int64, int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				e, err := al.RemoveAt(1)

				return toSlice[int64](al), e, al.Size(), err
			},
			expectedResult:  []int64{1, 3, 4},
			expectedElement: 2,
			expectedSize:    3,
		},
		{
			name: "should return error when list is empty",
			actualResult: func() ([]int64, int64, int64, error) {
				al := NewArrayList[int64]()

				e, err := al.RemoveAt(0)

				return toSlice[int64](al), e, al.Size(), err
			},
			expectedResult: []int64{},
			expectedSize:   0,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "should return error when trying to remove element at invalid index",
			actualResult: func() ([]int64, int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				e, err := al.RemoveAt(4)

				return toSlice[int64](al), e, al.Size(), err
			},
			expectedResult: []int64{1, 2, 3, 4},
			expectedSize:   4,
			expectedError:  fmt.Errorf("invalid index %d", 4),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, el, sz, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedElement, el)
		})
	}
}

func TestArrayListRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], int64, error)
		expectedResult List[int64]
		expectedSize   int64
		expectedError  error
	}{
		{
			name: "should successfully remove elements",
			actualResult: func() (List[int64], int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RemoveAll(2, 4)

				return al, al.Size(), err
			},
			expectedResult: NewArrayList[int64](1, 3),
			expectedSize:   2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, sz, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

type isEven struct{}

func (ie isEven) Test(e int) bool {
	return e%2 == 0
}

func TestArrayListRemoveIf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int], int64, error)
		expectedResult List[int]
		expectedSize   int64
		expectedError  error
	}{
		{
			name: "should remove if number is even",
			actualResult: func() (List[int], int64, error) {
				al := NewArrayList[int](1, 2, 3, 4)

				err := al.RemoveIf(isEven{})

				return al, al.Size(), err
			},
			expectedResult: NewArrayList[int](1, 3),
			expectedSize:   2,
		},
		{
			name: "should remove no element when no element match predicate",
			actualResult: func() (List[int], int64, error) {
				al := NewArrayList[int](1, 3, 5, 7)

				err := al.RemoveIf(isEven{})

				return al, al.Size(), err
			},
			expectedResult: NewArrayList[int](1, 3, 5, 7),
			expectedSize:   4,
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (List[int], int64, error) {
				al := NewArrayList[int]()

				err := al.RemoveIf(isEven{})

				return al, al.Size(), err
			},
			expectedResult: NewArrayList[int](),
			expectedSize:   0,
			expectedError:  errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, sz, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveRange(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() ([]int64, int64, error)
		expectedResult []int64
		expectedSize   int64
		expectedError  error
	}{
		{
			name: "should remove elements from range 0 to 3",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4, 5, 6, 7, 8)

				err := al.RemoveRange(0, 3)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{5, 6, 7, 8},
			expectedSize:   4,
		},
		{
			name: "should remove elements from range 2 to 3",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RemoveRange(2, 3)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{1, 2},
			expectedSize:   2,
		},
		{
			name: "should remove single element when to and from are the same",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RemoveRange(0, 0)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{2, 3, 4},
			expectedSize:   3,
		},
		{
			name: "should return error when to is smaller than from",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RemoveRange(1, 0)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{1, 2, 3, 4},
			expectedSize:   4,
			expectedError:  errors.New("end cannot be smaller than start"),
		},
		{
			name: "should return error when start is an invalid index",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RemoveRange(-1, 2)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{1, 2, 3, 4},
			expectedSize:   4,
			expectedError:  fmt.Errorf("invalid index %d", -1),
		},
		{
			name: "should return error when end is an invalid index",
			actualResult: func() ([]int64, int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RemoveRange(1, 10)

				return toSlice[int64](al), al.Size(), err
			},
			expectedResult: []int64{1, 2, 3, 4},
			expectedSize:   4,
			expectedError:  fmt.Errorf("invalid index %d", 10),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, sz, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

type testIntIncOperator struct{}

func (ti testIntIncOperator) Apply(e int) int { return e + 1 }

func TestArrayListReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List[int])
		expectedResult List[int]
		expectedError  error
	}{
		{
			name: "should replace a given value with new one",
			actualResult: func() (error, List[int]) {
				al := NewArrayList[int](1)

				return al.Replace(1, 2), al
			},
			expectedResult: NewArrayList[int](2),
		},
		{
			name: "should return error when item is not found in the list",
			actualResult: func() (error, List[int]) {
				al := NewArrayList[int](1, 2)

				return al.Replace(5, 3), al
			},
			expectedResult: NewArrayList[int](1, 2),
			expectedError:  errors.New("element 5 not found in the list"),
		},
		{
			name: "should return error when list is empty",
			actualResult: func() (error, List[int]) {
				al := NewArrayList[int]()

				return al.Replace(1, 2), al
			},
			expectedResult: NewArrayList[int](),
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

func TestArrayListReplaceAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List[int]
		expectedResult List[int]
		expectedError  error
	}{
		{
			name: "should replace all on integer list with increment operator",
			actualResult: func() List[int] {
				al := NewArrayList[int](1, 2, 3, 4)

				al.ReplaceAll(testIntIncOperator{})

				return al
			},
			expectedResult: NewArrayList[int](2, 3, 4, 5),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int64], int64, error)
		expectedResult List[int64]
		expectedSize   int64
		expectedError  error
	}{
		{
			name: "should retain given values from integer list",
			actualResult: func() (List[int64], int64, error) {
				al := NewArrayList[int64](1, 2, 3, 4)

				err := al.RetainAll(2, 4)

				return al, al.Size(), err
			},
			expectedResult: NewArrayList[int64](2, 4),
			expectedSize:   2,
		},
		{
			name: "should retain given values from integer list second scenario",
			actualResult: func() (List[int64], int64, error) {
				al := NewArrayList[int64](internal.SliceGenerator{Size: 22}.Generate()...)

				err := al.RetainAll(internal.SliceGenerator{Size: 11}.Generate()...)

				return al, al.Size(), err
			},
			expectedResult: NewArrayList[int64](0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
			expectedSize:   11,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, sz, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedSize, sz)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListSubList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List[int], error)
		expectedResult List[int]
		expectedError  error
	}{
		{
			name: "should get sublist from index 1 to 4",
			actualResult: func() (List[int], error) {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

				return al.SubList(1, 4)
			},
			expectedResult: NewArrayList[int](2, 3, 4, 5),
		},
		{
			name: "should get sublist from index 0 to 0",
			actualResult: func() (List[int], error) {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

				return al.SubList(0, 0)
			},
			expectedResult: NewArrayList[int](1),
		},
		{
			name: "should get sublist from index 0 to 2",
			actualResult: func() (List[int], error) {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

				return al.SubList(0, 2)
			},
			expectedResult: NewArrayList[int](1, 2, 3),
		},
		{
			name: "should get sublist from index 4 to 6",
			actualResult: func() (List[int], error) {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

				return al.SubList(4, 6)
			},
			expectedResult: NewArrayList[int](5, 6, 7),
		},
		{
			name: "should return error for invalid start index",
			actualResult: func() (List[int], error) {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

				return al.SubList(-1, 6)
			},
			expectedError: fmt.Errorf("invalid index %d", -1),
		},
		{
			name: "should return error for invalid end index",
			actualResult: func() (List[int], error) {
				al := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

				return al.SubList(0, 7)
			},
			expectedError: fmt.Errorf("invalid index %d", 7),
		},
		{
			name: "should return error when end is less than start",
			actualResult: func() (List[int], error) {
				ll := NewArrayList[int](1, 2, 3, 4, 5, 6, 7)

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
