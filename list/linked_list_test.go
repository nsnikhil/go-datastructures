package list

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/functions/comparator"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewLinkedList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult List
		expectedError  error
	}{
		{
			name: "test create new empty linked list",
			actualResult: func() (List, error) {
				return newLinkedList()
			},
			expectedResult: &linkedList{
				typeURL: "na",
			},
		},
		{
			name: "test create new linked list with values",
			actualResult: func() (List, error) {
				return newLinkedList(1, 2)
			},
			expectedResult: &linkedList{
				typeURL: "int",
				root: &node{
					data: 1,
					next: &node{
						data: 2,
					},
				},
			},
		},
		{
			name: "test failed to create new linked list when type mis matches",
			actualResult: func() (List, error) {
				return newLinkedList(1, "a")
			},
			expectedResult: (*linkedList)(nil),
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

func TestLinkedListAdd(t *testing.T) {
	testCases := []struct {
		name         string
		actualSize   func() int
		expectedSize int
	}{
		{
			name: "test add one item to linked empty list",
			actualSize: func() int {
				list, err := newLinkedList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				return list.Size()
			},

			expectedSize: 1,
		},
		{
			name: "test add four item to linked empty list",
			actualSize: func() int {
				ll, err := newLinkedList()
				require.NoError(t, err)

				err = ll.Add(1)
				require.NoError(t, err)

				err = ll.Add(2)
				require.NoError(t, err)

				err = ll.Add(3)
				require.NoError(t, err)

				err = ll.Add(4)
				require.NoError(t, err)

				return ll.Size()
			},

			expectedSize: 4,
		},
		{
			name: "test failed to add item of different type",
			actualSize: func() int {
				list, err := newLinkedList()
				require.NoError(t, err)

				err = list.Add(1)
				require.NoError(t, err)

				err = list.Add("a")
				require.Error(t, err)

				return list.Size()
			},

			expectedSize: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedSize, testCase.actualSize())
		})
	}
}

func TestLinkedListGetSize(t *testing.T) {
	testCases := []struct {
		name         string
		actualSize   func() int
		expectedSize int
	}{
		{
			name: "test get list size 1",
			actualSize: func() int {
				list, err := newLinkedList(1)
				require.NoError(t, err)

				return list.Size()
			},
			expectedSize: 1,
		},
		{
			name: "test get list size 4",
			actualSize: func() int {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Size()
			},
			expectedSize: 4,
		},
		{
			name: "test get list size 0 for empty list",
			actualSize: func() int {
				list, err := newLinkedList()
				require.NoError(t, err)

				return list.Size()
			},
			expectedSize: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedSize, testCase.actualSize())
		})
	}
}

func TestLinkedListIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test iterator get all values",
			actualResult: func() interface{} {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				it := ll.Iterator()

				var res []interface{}

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 2, 3, 4},
		},
		{
			name: "test iterator get empty result for empty list",
			actualResult: func() interface{} {
				ll, err := newLinkedList()
				require.NoError(t, err)

				it := ll.Iterator()

				var res []interface{}

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}(nil),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestSortLinkedList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult func() List
	}{
		{
			name: "test Sort integer List",
			actualResult: func() List {
				ll, err := newLinkedList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				ll.Sort(comparator.NewIntegerComparator())

				return ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Sort string List with equal length",
			actualResult: func() List {
				ll, err := newLinkedList("e", "d", "c", "b", "a")
				require.NoError(t, err)

				ll.Sort(comparator.NewStringComparator())

				return ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList("a", "b", "c", "d", "e")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Sort string List with un equal length",
			actualResult: func() List {
				ll, err := newLinkedList("a", "aaa", "aaa", "a", "aaaa")
				require.NoError(t, err)

				ll.Sort(comparator.NewStringComparator())

				return ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList("a", "a", "aaa", "aaa", "aaaa")
				require.NoError(t, err)

				return ll
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), testCase.actualResult())
		})
	}
}

func TestLinkedListAddAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test Add element at index 1",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 3, 4, 5)
				require.NoError(t, err)

				return ll.AddAt(1, 2), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Add element at index 0",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.AddAt(0, 0), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(0, 1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Add element at index 3",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll.AddAt(3, 4), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add elements for empty linked list",
			actualResult: func() (error, List) {
				ll, err := newLinkedList()
				require.NoError(t, err)

				return ll.AddAt(0, 1), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return error when adding element at invalid index",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll.AddAt(4, 8), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewIndexOutOfBoundError(4),
		},
		{
			name: "test return error when adding element of invalid type",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll.AddAt(0, "a"), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, l := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), l)
		})
	}
}

func TestLinkedListAllAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test add all for empty list elements",
			actualResult: func() (error, List) {
				ll, err := newLinkedList()
				require.NoError(t, err)

				return ll.AddAll(1, 2, 3, 4), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add all for integer elements",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll.AddAll(3, 4), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add all for string elements",
			actualResult: func() (error, List) {
				ll, err := newLinkedList("a", "b")
				require.NoError(t, err)

				return ll.AddAll("c", "d"), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList("a", "b", "c", "d")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to Add all elements when type if different",
			actualResult: func() (error, List) {
				ll, err := newLinkedList("a", "b")
				require.NoError(t, err)

				return ll.AddAll("c", 5), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList("a", "b")
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("type mismatch : all elements must be of same type"),
		},
		{
			name: "test add all return nil when arguments are empty",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll.AddAll(), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, l := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), l)
		})
	}
}

func TestLinkedListClear(t *testing.T) {
	testCases := []struct {
		name   string
		result func() int
	}{
		{
			name: "test clear integer list",
			result: func() int {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()
				return ll.Size()
			},
		},
		{
			name: "test clear string list",
			result: func() int {
				ll, err := newLinkedList("a", "b")
				require.NoError(t, err)

				ll.Clear()

				return ll.Size()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, 0, testCase.result())
		})
	}
}

func TestLinkedListContains(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		{
			name: "test return true when element is present",
			actualResult: func() (bool, error) {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "test return false when element is not present",
			actualResult: func() (bool, error) {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Contains(0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "test return false when element type mismatch",
			actualResult: func() (bool, error) {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Contains("a")
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

func TestLinkedListGet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test Get nil for empty List",
			actualResult: func() interface{} {
				ll, err := newLinkedList()
				require.NoError(t, err)

				return ll.Get(0)
			},
		},
		{
			name: "test Get 0th element from the List",
			actualResult: func() interface{} {
				ll, err := newLinkedList(1)
				require.NoError(t, err)

				return ll.Get(0)
			},
			expectedResult: 1,
		},
		{
			name: "test Get 4th element from the List",
			actualResult: func() interface{} {
				ll, err := newLinkedList(0, 1, 2, 3)
				require.NoError(t, err)

				return ll.Get(3)
			},
			expectedResult: 3,
		},
		{
			name: "test get nil when index is greater than the Size of List",
			actualResult: func() interface{} {
				list, err := newLinkedList()
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

func TestLinkedListContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		{
			name: "return true when all elements are present",
			actualResult: func() (bool, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return ll.ContainsAll(6, 1, 3)
			},
			expectedResult: true,
		},
		{
			name: "return false when a element is not present",
			actualResult: func() (bool, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return ll.ContainsAll(6, 1, 0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return false when a element type mismatch",
			actualResult: func() (bool, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return ll.ContainsAll(6, 1, "a")
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

func TestLinkedListIndexOf(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "return index when element is found",
			actualResult: func() (int, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.IndexOf(2)
			},
			expectedResult: 1,
		},
		{
			name: "return -1 when element is not present",
			actualResult: func() (int, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return type mismatch error when searching for invalid type",
			actualResult: func() (int, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.IndexOf("a")
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

func TestLinkedListIsEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return true when list is empty",
			actualResult: func() bool {
				ll, err := newLinkedList()
				require.NoError(t, err)

				return ll.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return false when list is not empty",
			actualResult: func() bool {
				ll, err := newLinkedList(1)
				require.NoError(t, err)

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
		actualResult   func() (int, error)
		expectedResult int
		expectedError  error
	}{
		{
			name: "get last index of the element in List",
			actualResult: func() (int, error) {
				ll, err := newLinkedList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return ll.LastIndexOf(1)
			},
			expectedResult: 3,
		},
		{
			name: "return -1 when the element in not present",
			actualResult: func() (int, error) {
				ll, err := newLinkedList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return ll.LastIndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return type mismatch error when searching for different type",
			actualResult: func() (int, error) {
				ll, err := newLinkedList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return ll.LastIndexOf("a")
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

func TestLinkedListRemove(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test successfully remove element",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(2)
				require.NoError(t, err)
				require.True(t, ok)

				return ll, nil
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove first element",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(1)
				require.NoError(t, err)
				require.True(t, ok)

				return ll, nil
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove last element",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(4)
				require.NoError(t, err)
				require.True(t, ok)

				return ll, nil
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return false when trying to remove element which is not present",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(0)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("element 0 not found in the list"),
		},
		{
			name: "test return false when trying to Remove element of different type",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove("a")
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListRemoveAt(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test remove element at index 1",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(1)
				assert.Equal(t, 2, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove element at index 0",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(0)
				assert.Equal(t, 1, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove element at index 3",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(3)
				assert.Equal(t, 4, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to remove element at invalid index",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(4)
				assert.Nil(t, e)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewIndexOutOfBoundError(4),
		},
		{
			name: "test failed to remove element for empty list",
			actualResult: func() (List, error) {
				ll, err := newLinkedList()
				require.NoError(t, err)

				e, err := ll.RemoveAt(0)
				assert.Nil(t, e)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList()
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test successfully remove elements",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(2, 4)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove only elements which are present",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(2, 4, 5)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove when list has only one element",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				return &linkedList{
					typeURL: "int",
				}
			},
		},
		{
			name: "test successfully remove elements at start",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove elements at end",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(4)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to remove elements due to type mismatch",
			actualResult: func() (List, error) {
				al, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveAll(2, "a")
				require.False(t, ok)

				return al, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListReplaceAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test replace all on integer List with increment operator",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.ReplaceAll(testIntIncOperator{}), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test replace all on string List with concat operator",
			actualResult: func() (error, List) {
				ll, err := newLinkedList("a", "b")
				require.NoError(t, err)

				return ll.ReplaceAll(testStringConcatOperator{}), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList("aa", "ba")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test replace all fails when operator return invalid data",
			actualResult: func() (error, List) {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll.ReplaceAll(testInvalidOperator{}), ll
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test retain all from integer List",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RetainAll(2, 4)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test retain all from integer list with ignoring element which are not present",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RetainAll(2, 4, 5)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test retain all from string List",
			actualResult: func() (List, error) {
				ll, err := newLinkedList("a", "b", "c", "d")
				require.NoError(t, err)

				ok, err := ll.RetainAll("b", "d")
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList("b", "d")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return error when type mismatch",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RetainAll(2, "d")
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestLinkedListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (interface{}, error)
		expectedResult interface{}
		expectedError  error
	}{
		{
			name: "test set value at index 3",
			actualResult: func() (interface{}, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Set(3, 5)
			},
			expectedResult: 5,
		},
		{
			name: "test set value at index 0",
			actualResult: func() (interface{}, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Set(0, 2)
			},
			expectedResult: 2,
		},
		{
			name: "test set value at index 1",
			actualResult: func() (interface{}, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Set(1, 4)
			},
			expectedResult: 4,
		},
		{
			name: "test failed to set value due to invalid index",
			actualResult: func() (interface{}, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Set(5, 3)
			},
			expectedError: liberror.NewIndexOutOfBoundError(5),
		},
		{
			name: "test failed to set value due to different element type",
			actualResult: func() (interface{}, error) {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Set(1, "a")
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

func TestLinkedListSubList(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test get sublist from index 1 to 4",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(1, 4)
			},
			expectedResult: func() List {
				ll, err := newLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist from index 0 to 0",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(0, 0)
			},
			expectedResult: func() List {
				ll, err := newLinkedList()
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist from index 0 to 4",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(0, 4)
			},
			expectedResult: func() List {
				ll, err := newLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist from index 4 to 6",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(4, 6)
			},
			expectedResult: func() List {
				ll, err := newLinkedList(5, 6)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist return error for invalid start index",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(-1, 4)
			},
			expectedResult: func() List {
				return nil
			},
			expectedError: liberror.NewIndexOutOfBoundError(-1),
		},
		{
			name: "test get sublist return error for invalid end index",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(0, 10)
			},
			expectedResult: func() List {
				return nil
			},
			expectedError: liberror.NewIndexOutOfBoundError(10),
		},
		{
			name: "test get sublist return error when end is less than start",
			actualResult: func() (List, error) {
				ll, err := newLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(4, 2)
			},
			expectedResult: func() List {
				return nil
			},
			expectedError: errors.New("end cannot be smaller than start"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestAllLinkedListAPI(t *testing.T) {
	ll, err := newLinkedList()
	require.NoError(t, err)

	err = ll.Add(2)
	require.NoError(t, err)

	err = ll.Add("a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)

	err = ll.AddAt(0, 1)
	require.NoError(t, err)

	err = ll.AddAt(2, 1)
	require.Error(t, err)
	assert.Equal(t, liberror.NewIndexOutOfBoundError(2), err)

	err = ll.AddAt(0, "a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)

	err = ll.AddAll(3, 4, 5)
	require.NoError(t, err)

	err = ll.AddAll(5, 6, "a")
	require.Error(t, err)
	assert.Equal(t, "type mismatch : all elements must be of same type", err.Error())

	ok, err := ll.Contains(1)
	require.NoError(t, err)
	assert.True(t, ok)

	ok, err = ll.Contains(8)
	require.Error(t, err)
	assert.Equal(t, "element 8 not found in the list", err.Error())
	assert.False(t, ok)

	ok, err = ll.Contains("a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	assert.False(t, ok)

	ok, err = ll.ContainsAll(2, 4, 5)
	require.NoError(t, err)
	assert.True(t, ok)

	ok, err = ll.ContainsAll(4, 7)
	require.Error(t, err)
	assert.Equal(t, "element 7 not found in the list", err.Error())
	assert.False(t, ok)

	ok, err = ll.ContainsAll(4, "a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	assert.False(t, ok)

	ele := ll.Get(0)
	assert.Equal(t, 1, ele)

	ele = ll.Get(10)
	assert.Nil(t, ele)

	id, err := ll.IndexOf(2)
	require.NoError(t, err)
	assert.Equal(t, 1, id)

	id, err = ll.IndexOf(10)
	require.Error(t, err)
	assert.Equal(t, "element 10 not found in the list", err.Error())
	assert.Equal(t, -1, id)

	id, err = ll.IndexOf("a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	assert.Equal(t, -1, id)

	ok = ll.IsEmpty()
	require.False(t, ok)

	it := ll.Iterator()

	count := 1
	for it.HasNext() {
		assert.Equal(t, count, it.Next())
		count++
	}

	require.NoError(t, ll.Add(1))

	id, err = ll.LastIndexOf(1)
	require.NoError(t, err)
	assert.Equal(t, 5, id)

	id, err = ll.LastIndexOf(10)
	require.Error(t, err)
	assert.Equal(t, "element 10 not found in the list", err.Error())
	assert.Equal(t, -1, id)

	id, err = ll.LastIndexOf("a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	assert.Equal(t, -1, id)

	ok, err = ll.Remove(1)
	require.NoError(t, err)
	assert.True(t, ok)

	ok, err = ll.Remove(10)
	require.Error(t, err)
	assert.Equal(t, "element 10 not found in the list", err.Error())
	assert.False(t, ok)

	ok, err = ll.Remove("a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	assert.False(t, ok)

	ele, err = ll.RemoveAt(0)
	require.NoError(t, err)
	assert.Equal(t, 2, ele)

	ele, err = ll.RemoveAt(10)
	require.Error(t, err)
	assert.Equal(t, liberror.NewIndexOutOfBoundError(10), err)
	assert.Nil(t, ele)

	require.NoError(t, ll.Add(3))

	ok, err = ll.RemoveAll(3, 5)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = ll.RemoveAll(3, "a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	require.False(t, ok)

	err = ll.ReplaceAll(testIntIncOperator{})
	require.NoError(t, err)

	err = ll.ReplaceAll(testStringConcatOperator{})
	require.Error(t, err)
	assert.Equal(t, "type mismatch : interface conversion: interface {} is int, not string", err.Error())

	require.NoError(t, ll.AddAll(4, 6, 3, 1))

	ok, err = ll.RetainAll(6, 2, 4)
	require.NoError(t, err)
	require.True(t, ok)

	ok, err = ll.RetainAll(1, 3, 5, "a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	require.False(t, ok)

	ele, err = ll.Set(0, 7)
	require.NoError(t, err)
	assert.Equal(t, 7, ele)

	ele, err = ll.Set(3, 8)
	require.Error(t, err)
	assert.Equal(t, liberror.NewIndexOutOfBoundError(3), err)
	assert.Nil(t, ele)

	ele, err = ll.Set(1, "a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)
	assert.Nil(t, ele)

	sz := ll.Size()
	assert.Equal(t, 3, sz)

	ll.Sort(comparator.NewIntegerComparator())

	val := []int{4, 6, 7}
	i := 0
	it = ll.Iterator()
	for it.HasNext() {
		assert.Equal(t, val[i], it.Next())
		i++
	}

	sl, err := ll.SubList(0, 0)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(), sl)

	sl, err = ll.SubList(1, 1)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(), sl)

	sl, err = ll.SubList(2, 2)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(), sl)

	sl, err = ll.SubList(0, 1)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(4), sl)

	sl, err = ll.SubList(0, 2)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(4, 6), sl)

	sl, err = ll.SubList(1, 2)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(6), sl)

	sl, err = ll.SubList(1, 3)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(6, 7), sl)

	sl, err = ll.SubList(2, 3)
	require.NoError(t, err)
	assert.Equal(t, tempLinkedList(7), sl)

	ll.Clear()

	assert.Equal(t, 0, ll.Size())

	err = ll.Add("a")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)

	err = ll.AddAll("a", "b")
	require.Error(t, err)
	assert.Equal(t, liberror.NewTypeMismatchError("int", "string"), err)

	err = ll.AddAll(1, "a", "b")
	require.Error(t, err)
	assert.Equal(t, "type mismatch : all elements must be of same type", err.Error())

	id, err = ll.IndexOf(1)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	assert.Equal(t, -1, id)

	id, err = ll.LastIndexOf(1)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	assert.Equal(t, -1, id)

	ok, err = ll.Remove(1)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	require.False(t, ok)

	ele, err = ll.RemoveAt(0)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	assert.Nil(t, ele)

	ok, err = ll.RemoveAll(1, 2)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	require.False(t, ok)

	ok, err = ll.RetainAll(1, 2)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	require.False(t, ok)

	ele, err = ll.Set(0, 1)
	require.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
	assert.Nil(t, ele)

}

func tempLinkedList(data ...interface{}) *linkedList {
	ll, _ := newLinkedList(data...)
	return ll
}
