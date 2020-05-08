package list

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCreateNewArrayList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test create new empty array List",
			actualResult: func() (List, error) {
				return NewArrayList()
			},
			expectedResult: &ArrayList{typeURL: "na"},
		},
		{
			name: "test create new array List with elements",
			actualResult: func() (List, error) {
				return NewArrayList(1, 2, 3, 4, 5)
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
		{
			name: "test failed to create new array List due to element type mismatch",
			actualResult: func() (List, error) {
				return NewArrayList(1, "2")
			},
			expectedResult: (*ArrayList)(nil),
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}

}

func TestArrayListSize(t *testing.T) {
	testCases := []struct {
		name         string
		actualSize   func() int
		expectedSize int
	}{
		{
			name: "test Size is 1 after adding one element",
			actualSize: func() int {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)
				return list.Size()
			},
			expectedSize: 1,
		},
		{
			name: "test Size is 0 for a new List",
			actualSize: func() int {
				list, err := NewArrayList()
				require.NoError(t, err)

				return list.Size()
			},
			expectedSize: 0,
		},

		{
			name: "test Size is 2 after adding two elements",
			actualSize: func() int {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add("a")
				require.NoError(t, err)

				err = list.Add("b")
				require.NoError(t, err)

				return list.Size()
			},
			expectedSize: 2,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectedSize, testCase.actualSize())
	}
}

func TestArrayListAdd(t *testing.T) {
	testCases := []struct {
		name          string
		actualResult  func() (int, error)
		expectedError error
		expectedSize  int
	}{
		{
			name: "test Size is 1 after adding one element",
			actualResult: func() (int, error) {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				return list.Size(), err
			},
			expectedSize: 1,
		},
		{
			name: "test Size is 2 after adding two element",
			actualResult: func() (int, error) {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				err = list.Add(2)
				return list.Size(), err
			},
			expectedSize: 2,
		},
		{
			name: "test Size is 1 after trying to Add element of different type",
			actualResult: func() (int, error) {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				err = list.Add("name")
				return list.Size(), err
			},
			expectedSize:  1,
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test Size is 2 after adding structs",
			actualResult: func() (int, error) {
				type testStruct struct{}

				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(testStruct{})
				require.NoError(t, err)

				err = list.Add(testStruct{})
				return list.Size(), err
			},
			expectedSize: 2,
		},
	}

	for _, testCase := range testCases {
		size, err := testCase.actualResult()
		assert.Equal(t, testCase.expectedError, err)
		assert.Equal(t, testCase.expectedSize, size)
	}
}

func TestArrayListGet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test Get nil for empty List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				return list.Get(0)
			},
		},
		{
			name: "test Get 0th element from the List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				return list.Get(0)
			},
			expectedResult: 1,
		},
		{
			name: "test Get 4th element from the List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(0)
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				err = list.Add(2)
				require.NoError(t, err)

				err = list.Add(3)
				require.NoError(t, err)

				return list.Get(3)
			},
			expectedResult: 3,
		},
		{
			name: "test get nil when index is greater than the Size of List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(0)
				require.NoError(t, err)

				return list.Get(1)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
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
			name: "test has next return false for empty List",
			actualResult: func() bool {
				list, err := NewArrayList()
				require.NoError(t, err)

				return list.Iterator().HasNext()
			},
		},
		{
			name: "test has next return true for non empty List",
			actualResult: func() bool {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)
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
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test Get nil for empty List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				return list.Iterator().Next()
			},
		},
		{
			name: "test Get first item from List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				return list.Iterator().Next()
			},
			expectedResult: 1,
		},
		{
			name: "test Get all items from List",
			actualResult: func() interface{} {
				list, err := NewArrayList()
				require.NoError(t, err)

				err = list.Add(0)
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				err = list.Add(2)
				require.NoError(t, err)

				err = list.Add(3)
				require.NoError(t, err)

				i := list.Iterator()

				var res []interface{}

				for i.HasNext() {
					res = append(res, i.Next())
				}

				return res
			},
			expectedResult: []interface{}{0, 1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestArrayListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (interface{}, error)
		expectedResult interface{}
		expectedError  error
	}{
		{
			name: "test set value at index 3",
			actualResult: func() (interface{}, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)
				return al.Set(3, 5)
			},
			expectedResult: 5,
		},
		{
			name: "test set value at index 0",
			actualResult: func() (interface{}, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.Set(0, 2)
			},
			expectedResult: 2,
		},
		{
			name: "test set value at index 1",
			actualResult: func() (interface{}, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.Set(1, 4)
			},
			expectedResult: 4,
		},
		{
			name: "test failed to set value due to invalid index",
			actualResult: func() (interface{}, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.Set(5, 3)
			},
			expectedError: liberror.NewIndexOutOfBoundError(5),
		},
		{
			name: "test set fails when list is empty",
			actualResult: func() (interface{}, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.Set(0, 3)
			},
			expectedError: errors.New("list is empty"),
		},
		{
			name: "test failed to set value due to type mismatch",
			actualResult: func() (interface{}, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.Set(0, "a")
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListSort(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult List
	}{
		{
			name: "test Sort integer List",
			actualResult: func() List {
				al, err := NewArrayList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				al.Sort(comparator.NewIntegerComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
		{
			name: "test Sort string List with equal length",
			actualResult: func() List {
				al, err := NewArrayList("e", "d", "c", "b", "a")
				require.NoError(t, err)

				al.Sort(comparator.NewStringComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"a", "b", "c", "d", "e"},
			},
		},
		{
			name: "test Sort string List with un equal length",
			actualResult: func() List {
				al, err := NewArrayList("a", "aaa", "aaa", "a", "aaaa")
				require.NoError(t, err)

				al.Sort(comparator.NewStringComparator())

				return al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"a", "a", "aaa", "aaa", "aaaa"},
			},
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
		actualResult   func() (error, List)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test add element when list is empty",
			actualResult: func() (error, List) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.AddAt(0, 1), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1},
			},
		},
		{
			name: "test add element at index 1",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 3, 4, 5)
				require.NoError(t, err)

				return al.AddAt(1, 2), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
		{
			name: "test add element at index 0",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.AddAt(0, 0), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{0, 1, 2, 3, 4},
			},
		},
		{
			name: "test add element at index 3",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2, 3, 5)
				require.NoError(t, err)

				return al.AddAt(3, 4), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4, 5},
			},
		},
		{
			name: "test return error when adding element at invalid index",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2, 3, 5)
				require.NoError(t, err)

				return al.AddAt(4, 8), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 5},
			},
			expectedError: liberror.NewIndexOutOfBoundError(4),
		},
		{
			name: "test return error when adding element of invalid type",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2, 3, 5)
				require.NoError(t, err)

				return al.AddAt(0, "a"), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 5},
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, l := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, l)
		})
	}
}

func TestArrayListAddAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test Add all for integer elements",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.AddAll(3, 4), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
		},
		{
			name: "test Add all for string elements",
			actualResult: func() (error, List) {
				al, err := NewArrayList("a", "b")
				require.NoError(t, err)

				return al.AddAll("c", "d"), al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"a", "b", "c", "d"},
			},
		},
		{
			name: "test failed to Add all elements when type if different",
			actualResult: func() (error, List) {
				al, err := NewArrayList("a", "b")
				require.NoError(t, err)

				return al.AddAll("c", 5), al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"a", "b"},
			},
			expectedError: errors.New("type mismatch : all elements must be of same type"),
		},
		{
			name: "test add all return nil when arguments are empty",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.AddAll(), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
		},
		{
			name: "test add all when list is new",
			actualResult: func() (error, List) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.AddAll(1, 2), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
		},
		{
			name: "test add all when list is empty",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				al.Clear()

				return al.AddAll("a", "b"), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}(nil),
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add all when list is empty and type was set",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				al.Clear()

				return al.AddAll(1, 2), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, l := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, l)
		})
	}
}

func TestArrayListClear(t *testing.T) {
	testCases := []struct {
		name   string
		result func() int
	}{
		{
			name: "test clear integer list",
			result: func() int {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				al.Clear()
				return al.Size()
			},
		},
		{
			name: "test clear string list",
			result: func() int {
				al, err := NewArrayList("a", "b")
				require.NoError(t, err)

				al.Clear()

				return al.Size()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, 0, testCase.result())
		})
	}
}

func TestArrayListClone(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test clone integer array list",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.Clone()
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
		},
		{
			name: "test clone empty array list",
			actualResult: func() (List, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.Clone()
			},
			expectedResult: &ArrayList{
				typeURL: "na",
				data:    []interface{}(nil),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListContains(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		{
			name: "test return true when element is present",
			actualResult: func() (bool, error) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "test return false when element is not present",
			actualResult: func() (bool, error) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Contains(0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "test return false when element type mismatch",
			actualResult: func() (bool, error) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Contains("a")
			},
			expectedResult: false,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		{
			name: "return true when all elements are present",
			actualResult: func() (bool, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return al.ContainsAll(6, 1, 3)
			},
			expectedResult: true,
		},
		{
			name: "return false when a element is not present",
			actualResult: func() (bool, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return al.ContainsAll(6, 1, 0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return false when a element type mismatch",
			actualResult: func() (bool, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return al.ContainsAll(6, 1, "a")
			},
			expectedResult: false,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "return index when element is found",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.IndexOf(2)
			},
			expectedResult: 1,
		},
		{
			name: "return -1 when element is not present",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return error when list is empty",
			actualResult: func() (int, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "return type mismatch error when searching for invalid type",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.IndexOf("a")
			},
			expectedResult: -1,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
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
			name: "return true when list is empty",
			actualResult: func() bool {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return false when list is not empty",
			actualResult: func() bool {
				al, err := NewArrayList(1)
				require.NoError(t, err)

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
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "get last index of the element in List",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return al.LastIndexOf(1)
			},
			expectedResult: 3,
		},
		{
			name: "return -1 when the element in not present",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return al.LastIndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return error when list is empty",
			actualResult: func() (int, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.LastIndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "return type mismatch error when searching for different type",
			actualResult: func() (int, error) {
				al, err := NewArrayList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return al.LastIndexOf("a")
			},
			expectedResult: -1,
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveElement(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test successfully Remove element",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.Remove(2)
				require.NoError(t, err)
				require.True(t, ok)

				return al, nil
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 3, 4},
			},
		},
		{
			name: "test return false when trying to Remove element which is not present",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.Remove(0)
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: errors.New("element 0 not found in the list"),
		},
		{
			name: "test return false and error when list is empty",
			actualResult: func() (List, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				ok, err := al.Remove(0)
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "na",
				data:    []interface{}(nil),
			},
			expectedError: errors.New("list is empty"),
		},
		{
			name: "test return false when trying to Remove element of different type",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.Remove("a")
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test remove element at index 1",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := al.RemoveAt(1)
				assert.Equal(t, 2, e)
				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 3, 4},
			},
		},
		{
			name: "test return error when list is empty",
			actualResult: func() (List, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				e, err := al.RemoveAt(0)
				assert.Nil(t, e)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "na",
				data:    []interface{}(nil),
			},
			expectedError: errors.New("list is empty"),
		},
		{
			name: "test failed to remove element at invalid index",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := al.RemoveAt(4)
				assert.Nil(t, e)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: liberror.NewIndexOutOfBoundError(4),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test successfully Remove elements",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveAll(2, 4)
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 3},
			},
		},
		{
			name: "test failed to Remove elements due to type mismatch",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveAll(2, "a")
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

type isEven struct{}

func (ie isEven) Test(e interface{}) bool {
	return e.(int)%2 == 0
}

type containsA struct{}

func (ca containsA) Test(e interface{}) bool {
	return strings.Contains(e.(string), "a")
}

func TestArrayListRemoveIf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test remove if number is even",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveIf(isEven{})
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 3},
			},
		},
		{
			name: "test remove if string contains a",
			actualResult: func() (List, error) {
				al, err := NewArrayList("a", "bcd", "abg", "efg", "gb")
				require.NoError(t, err)

				ok, err := al.RemoveIf(containsA{})
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"bcd", "efg", "gb"},
			},
		},
		{
			name: "test remove no element when no element match predicate",
			actualResult: func() (List, error) {
				al, err := NewArrayList("bcd", "efg", "hij")
				require.NoError(t, err)

				ok, err := al.RemoveIf(containsA{})
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"bcd", "efg", "hij"},
			},
		},
		{
			name: "test return error when list is empty",
			actualResult: func() (List, error) {
				al, err := NewArrayList()
				require.NoError(t, err)

				ok, err := al.RemoveIf(isEven{})
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "na",
				data:    []interface{}(nil),
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRemoveRange(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test remove elements from range 0 to 3",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7, 8)
				require.NoError(t, err)

				ok, err := al.RemoveRange(0, 3)
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{4, 5, 6, 7, 8},
			},
		},
		{
			name: "test remove elements from range 2 to 4",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveRange(2, 4)
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
		},
		{
			name: "test remove no elements when to and from are the same",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveRange(0, 0)
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
		},
		{
			name: "test return error when to is smaller than from",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveRange(1, 0)
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: errors.New("to cannot be smaller than from"),
		},
		{
			name: "test return error when from is an invalid index",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveRange(-1, 2)
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: liberror.NewIndexOutOfBoundError(-1),
		},
		{
			name: "test return error when to is an invalid index",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveRange(1, 10)
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: liberror.NewIndexOutOfBoundError(10),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

type testIntIncOperator struct{}

func (ti testIntIncOperator) Apply(e interface{}) interface{} { return e.(int) + 1 }

type testStringConcatOperator struct{}

func (ts testStringConcatOperator) Apply(e interface{}) interface{} {
	return fmt.Sprintf("%s%s", e.(string), "a")
}

type testInvalidOperator struct{}

func (ts testInvalidOperator) Apply(e interface{}) interface{} {
	return fmt.Sprintf("%d", e.(int))
}

func TestArrayListReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test replace a given value with new one",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1)
				require.NoError(t, err)

				return al.Replace(1, 2), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{2},
			},
		},
		{
			name: "test replace a given value with new one two",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2, 5, 4)
				require.NoError(t, err)

				return al.Replace(5, 3), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
		},
		{
			name: "test return error when item is not found in the list",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Replace(5, 3), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
			expectedError: errors.New("element 5 not found in the list"),
		},
		{
			name: "test return error when old item has different type",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Replace('a', 3), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when new item has different type",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Replace(1, 'a'), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when old and new item has different type",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.Replace('a', 'b'), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when list is empty",
			actualResult: func() (error, List) {
				al, err := NewArrayList()
				require.NoError(t, err)

				return al.Replace(1, 2), al
			},
			expectedResult: &ArrayList{
				typeURL: "na",
				data:    []interface{}(nil),
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListReplaceAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test replace all on integer List with increment operator",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				return al.ReplaceAll(testIntIncOperator{}), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{2, 3, 4, 5},
			},
		},
		{
			name: "test replace all on string List with concat operator",
			actualResult: func() (error, List) {
				al, err := NewArrayList("a", "b")
				require.NoError(t, err)

				return al.ReplaceAll(testStringConcatOperator{}), al
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"aa", "ba"},
			},
		},
		{
			name: "test replace all fails when operator return invalid data",
			actualResult: func() (error, List) {
				al, err := NewArrayList(1, 2)
				require.NoError(t, err)

				return al.ReplaceAll(testInvalidOperator{}), al
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test retain all from integer List",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RetainAll(2, 4)
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{2, 4},
			},
		},
		{
			name: "test retain all from string List",
			actualResult: func() (List, error) {
				al, err := NewArrayList("a", "b", "c", "d")
				require.NoError(t, err)

				ok, err := al.RetainAll("b", "d")
				require.True(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "string",
				data:    []interface{}{"b", "d"},
			},
		},
		{
			name: "test retain all from string List",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RetainAll(2, "c")
				require.False(t, ok)

				return al, err
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2, 3, 4},
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}

func TestArrayListSubList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test Get sublist from index 1 to 4",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return al.SubList(1, 4)
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{2, 3, 4},
			},
		},
		{
			name: "test Get sublist from index 0 to 0",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return al.SubList(0, 0)
			},
			expectedResult: &ArrayList{
				typeURL: "na",
				data:    []interface{}(nil),
			},
		},
		{
			name: "test Get sublist from index 0 to 4",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return al.SubList(0, 2)
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{1, 2},
			},
		},
		{
			name: "test Get sublist from index 4 to 6",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return al.SubList(4, 6)
			},
			expectedResult: &ArrayList{
				typeURL: "int",
				data:    []interface{}{5, 6},
			},
		},
		{
			name: "test Get sublist return error for invalid start index",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return al.SubList(-1, 6)
			},
			expectedError: liberror.NewIndexOutOfBoundError(-1),
		},
		{
			name: "test Get sublist return error for invalid end index",
			actualResult: func() (List, error) {
				al, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return al.SubList(0, 10)
			},
			expectedError: liberror.NewIndexOutOfBoundError(10),
		},
		{
			name: "test get sublist return error when end is less than start",
			actualResult: func() (List, error) {
				ll, err := NewArrayList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(4, 2)
			},
			expectedError: errors.New("end cannot be smaller than start"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
		})
	}
}
