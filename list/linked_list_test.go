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
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test create new empty linked list",
			actualResult: func() (List, error) {
				return NewLinkedList()
			},
			expectedResult: func() List {
				return &LinkedList{
					typeURL: "na",
				}
			},
		},
		{
			name: "test create new linked list with one value",
			actualResult: func() (List, error) {
				return NewLinkedList(1)
			},
			expectedResult: func() List {
				l := &LinkedList{}

				o := newNode(1)

				l.first = o
				l.last = o
				l.typeURL = "int"

				return l
			},
		},
		{
			name: "test create new linked list with two values",
			actualResult: func() (List, error) {
				return NewLinkedList(1, 2)
			},
			expectedResult: func() List {
				l := &LinkedList{}

				o := newNode(1)
				t := newNode(2)

				o.next = t
				t.prev = o

				l.first = o
				l.last = t
				l.typeURL = "int"

				return l
			},
		},
		{
			name: "test create new linked list with four values",
			actualResult: func() (List, error) {
				return NewLinkedList(1, 2, 3, 4)
			},
			expectedResult: func() List {
				l := &LinkedList{}

				o := newNode(1)
				t := newNode(2)
				th := newNode(3)
				fr := newNode(4)

				o.next = t
				t.prev = o

				t.next = th
				th.prev = t

				th.next = fr
				fr.prev = th

				l.first = o
				l.last = fr
				l.typeURL = "int"

				return l
			},
		},
		{
			name: "test failed to create new linked list when type mis matches",
			actualResult: func() (List, error) {
				return NewLinkedList(1, "a")
			},
			expectedResult: func() List { return (*LinkedList)(nil) },
			expectedError:  liberror.NewTypeMismatchError("int", "string"),
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

func TestLinkedListAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() List
		expectedResult func() List
	}{
		{
			name: "test add one item to linked new list",
			actualResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				err = ll.Add(1)
				require.NoError(t, err)

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add one item to linked empty list",
			actualResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				err = ll.Add(1)
				require.NoError(t, err)

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add four item to linked empty list",
			actualResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				err = ll.Add(1)
				require.NoError(t, err)

				err = ll.Add(2)
				require.NoError(t, err)

				err = ll.Add(3)
				require.NoError(t, err)

				err = ll.Add(4)
				require.NoError(t, err)

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to add item of different type",
			actualResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				err = ll.Add(1)
				require.NoError(t, err)

				err = ll.Add("a")
				require.Error(t, err)

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
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

func TestLinkedListGetSize(t *testing.T) {
	testCases := []struct {
		name         string
		actualSize   func() int
		expectedSize int
	}{
		{
			name: "test get list size 1",
			actualSize: func() int {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll.Size()
			},
			expectedSize: 1,
		},
		{
			name: "test get list size 4",
			actualSize: func() int {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Size()
			},
			expectedSize: 4,
		},
		{
			name: "test get list size 0 for empty list",
			actualSize: func() int {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.Size()
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
				ll, err := NewLinkedList(1, 2, 3, 4)
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
				ll, err := NewLinkedList()
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

func TestLinkedListDecrementIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test iterator get all values",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				it := ll.DescendingIterator()

				var res []interface{}

				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{4, 3, 2, 1},
		},
		{
			name: "test iterator get empty result for empty list",
			actualResult: func() interface{} {
				ll, err := NewLinkedList()
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
				ll, err := NewLinkedList(5, 4, 3, 2, 1)
				require.NoError(t, err)

				ll.Sort(comparator.NewIntegerComparator())

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Sort string List with equal length",
			actualResult: func() List {
				ll, err := NewLinkedList("e", "d", "c", "b", "a")
				require.NoError(t, err)

				ll.Sort(comparator.NewStringComparator())

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList("a", "b", "c", "d", "e")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Sort string List with un equal length",
			actualResult: func() List {
				ll, err := NewLinkedList("a", "aaa", "aaa", "a", "aaaa")
				require.NoError(t, err)

				ll.Sort(comparator.NewStringComparator())

				return ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList("a", "a", "aaa", "aaa", "aaaa")
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
				ll, err := NewLinkedList(1, 3, 4, 5)
				require.NoError(t, err)

				return ll.AddAt(1, 2), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Add element at index 0",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.AddAt(0, 0), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(0, 1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test Add element at index 3",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll.AddAt(3, 4), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add elements for empty linked list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.AddAt(0, 1), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return error when adding element at invalid index",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.AddAt(4, 5), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewIndexOutOfBoundError(4),
		},
		{
			name: "test return error when adding element of invalid type",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll.AddAt(0, "a"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 5)
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

func TestLinkedListAddFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test add element to start of list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll.AddFirst(1), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add element to start of new list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.AddFirst(1), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add element to start of empty list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				return ll.AddFirst(1), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add first return error when type mismatch",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.AddFirst("a"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add first return error when type mismatch for empty list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				return ll.AddFirst("a"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "int"
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

func TestLinkedListAddLast(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test add element to end of list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll.AddLast(4), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add element to end of new list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.AddLast(1), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add element to start of empty list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				return ll.AddLast(1), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add first return error when type mismatch",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.AddLast("a"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add first return error when type mismatch for empty list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				return ll.AddLast("a"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "int"
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

func TestLinkedListAllAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test add all for new list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.AddAll(1, 2, 3, 4), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add all for integer elements",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.AddAll(3, 4), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add all for string elements",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList("a", "b")
				require.NoError(t, err)

				return ll.AddAll("c", "d"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList("a", "b", "c", "d")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test add all when list is empty",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				return ll.AddAll(1, 2, 3, 4), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to add all elements when type if different",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList("a", "b")
				require.NoError(t, err)

				return ll.AddAll("c", 5), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList("a", "b")
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("type mismatch : all elements must be of same type"),
		},
		{
			name: "test failed to add all elements when type if different for empty list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				return ll.AddAll("c"), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "int"
				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test add all return nil when arguments are empty",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.AddAll(), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
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
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ll.Clear()

				require.Nil(t, ll.first)
				require.Nil(t, ll.last)

				return ll.Size()
			},
		},
		{
			name: "test clear string list",
			result: func() int {
				ll, err := NewLinkedList("a", "b")
				require.NoError(t, err)

				ll.Clear()

				require.Nil(t, ll.first)
				require.Nil(t, ll.last)

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

func TestLinkedListClone(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test clone integer linked list",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.Clone()
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test clone empty linked list",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.Clone()
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
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
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Contains(1)
			},
			expectedResult: true,
		},
		{
			name: "test return false when element is not present",
			actualResult: func() (bool, error) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Contains(0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "test return false when element type mismatch",
			actualResult: func() (bool, error) {
				ll, err := NewLinkedList(1, 2)
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
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.Get(0)
			},
		},
		{
			name: "test Get 0th element from the List",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll.Get(0)
			},
			expectedResult: 1,
		},
		{
			name: "test get 4th element from the List",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(0, 1, 2, 3)
				require.NoError(t, err)

				return ll.Get(3)
			},
			expectedResult: 3,
		},
		{
			name: "test get nil when index is greater than the Size of List",
			actualResult: func() interface{} {
				list, err := NewLinkedList()
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

func TestLinkedListGetFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test get first element",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.GetFirst()
			},
			expectedResult: 1,
		},
		{
			name: "test get first element when list contains one element",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll.GetFirst()
			},
			expectedResult: 1,
		},
		{
			name: "test return error when list is empty",
			actualResult: func() interface{} {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.GetFirst()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestLinkedListGetLast(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test get last element",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.GetLast()
			},
			expectedResult: 4,
		},
		{
			name: "test get last element when list contains one element",
			actualResult: func() interface{} {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll.GetLast()
			},
			expectedResult: 1,
		},
		{
			name: "test return error when list is empty",
			actualResult: func() interface{} {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.GetLast()
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
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return ll.ContainsAll(6, 1, 3)
			},
			expectedResult: true,
		},
		{
			name: "return false when a element is not present",
			actualResult: func() (bool, error) {
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6)
				require.NoError(t, err)

				return ll.ContainsAll(6, 1, 0)
			},
			expectedResult: false,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return false when a element type mismatch",
			actualResult: func() (bool, error) {
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6)
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
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.IndexOf(2)
			},
			expectedResult: 1,
		},
		{
			name: "return -1 when element is not present",
			actualResult: func() (int, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return error when list is empty",
			actualResult: func() (int, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.IndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("list is empty"),
		},
		{
			name: "return type mismatch error when searching for invalid type",
			actualResult: func() (int, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
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
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return false when list is not empty",
			actualResult: func() bool {
				ll, err := NewLinkedList(1)
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
			name: "get last index of the element in list",
			actualResult: func() (int, error) {
				ll, err := NewLinkedList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return ll.LastIndexOf(1)
			},
			expectedResult: 3,
		},
		{
			name: "return -1 when the element in not present",
			actualResult: func() (int, error) {
				ll, err := NewLinkedList(1, 2, 3, 1, 4)
				require.NoError(t, err)

				return ll.LastIndexOf(0)
			},
			expectedResult: -1,
			expectedError:  errors.New("element 0 not found in the list"),
		},
		{
			name: "return type mismatch error when searching for different type",
			actualResult: func() (int, error) {
				ll, err := NewLinkedList(1, 2, 3, 1, 4)
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
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(2)
				require.NoError(t, err)
				require.True(t, ok)

				return ll, nil
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove first element",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(1)
				require.NoError(t, err)
				require.True(t, ok)

				return ll, nil
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove last element",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(4)
				require.NoError(t, err)
				require.True(t, ok)

				return ll, nil
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return false when trying to remove element which is not present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove(0)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("element 0 not found in the list"),
		},
		{
			name: "test return false when trying to Remove element of different type",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.Remove("a")
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
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
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(1)
				assert.Equal(t, 2, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove element at index 0",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(0)
				assert.Equal(t, 1, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove element at index 3",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(3)
				assert.Equal(t, 4, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove element at index 0 when list only contains one element",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				e, err := ll.RemoveAt(0)
				assert.Equal(t, 1, e)
				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "int"
				return ll
			},
		},
		{
			name: "test failed to remove element at invalid index",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				e, err := ll.RemoveAt(4)
				assert.Nil(t, e)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewIndexOutOfBoundError(4),
		},
		{
			name: "test failed to remove element for empty list",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				e, err := ll.RemoveAt(0)
				assert.Nil(t, e)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
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
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(2, 4)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove only elements which are present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(2, 4, 5)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove when list has only one element",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				return &LinkedList{
					typeURL: "int",
				}
			},
		},
		{
			name: "test successfully remove elements at start",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test successfully remove elements at end",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll(4)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove all keeps all element when argument list is empty",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveAll()
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to remove elements due to type mismatch",
			actualResult: func() (List, error) {
				al, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := al.RemoveAll(2, "a")
				require.False(t, ok)

				return al, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test remove all fails when list is empty",
			actualResult: func() (List, error) {
				al, err := NewLinkedList()
				require.NoError(t, err)

				ok, err := al.RemoveAll(1, 2)
				require.False(t, ok)

				return al, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
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

func TestLinkedListRemoveFirst(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, interface{}, error)
		expectedList   func() List
		expectedResult interface{}
		expectedError  error
	}{
		{
			name: "test get and remove first element",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.RemoveFirst()

				return ll, ele, err
			},
			expectedResult: 1,
			expectedList: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get and remove first element when list contains one element",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				ele, err := ll.RemoveFirst()

				return ll, ele, err
			},
			expectedResult: 1,
			expectedList: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "int"
				return ll
			},
		},
		{
			name: "test poll first return error when list is empty",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ele, err := ll.RemoveFirst()

				return ll, ele, err
			},
			expectedList: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l, res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList(), l)
		})
	}
}

func TestLinkedListRemoveLast(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, interface{}, error)
		expectedList   func() List
		expectedResult interface{}
		expectedError  error
	}{
		{
			name: "test get and remove last element",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.RemoveLast()

				return ll, ele, err
			},
			expectedResult: 4,
			expectedList: func() List {
				ll, err := NewLinkedList(1, 2, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get and remove last element when list contains one element",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				ele, err := ll.RemoveLast()

				return ll, ele, err
			},
			expectedResult: 1,
			expectedList: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "int"
				return ll
			},
		},
		{
			name: "test remove last return error when list is empty",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ele, err := ll.RemoveLast()

				return ll, ele, err
			},
			expectedList: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l, res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList(), l)
		})
	}
}

func TestLinkedListRemoveFirstOccurrence(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test remove first occurrence of 1 when multiple occurrence of 1 is present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 1, 3, 1)
				require.NoError(t, err)

				ok, err := ll.RemoveFirstOccurrence(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 1, 3, 1)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove first occurrence of 1 when single occurrence of 1 is present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveFirstOccurrence(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove first occurrence return error when element is not present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveFirstOccurrence(5)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("element 5 not found in the list"),
		},
		{
			name: "test remove first occurrence return error when list is empty",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ok, err := ll.RemoveFirstOccurrence(1)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("list is empty"),
		},
		{
			name: "test remove first occurrence return error when type mismatch",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ok, err := ll.RemoveFirstOccurrence("a")
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
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

func TestLinkedListRemoveLastOccurrence(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, error)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test remove last occurrence of 1 when multiple occurrence of 1 is present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 1, 3, 1)
				require.NoError(t, err)

				ok, err := ll.RemoveLastOccurrence(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 1, 3)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove last occurrence of 1 when single occurrence of 1 is present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveLastOccurrence(1)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test remove last occurrence return error when element is not present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RemoveLastOccurrence(5)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("element 5 not found in the list"),
		},
		{
			name: "test remove last occurrence return error when list is empty",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ok, err := ll.RemoveLastOccurrence(1)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("list is empty"),
		},
		{
			name: "test remove last occurrence return error when type mismatch",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				ok, err := ll.RemoveLastOccurrence("a")
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
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

func TestLinkedListReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, List)
		expectedResult func() List
		expectedError  error
	}{
		{
			name: "test replace a given value with new one",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1)
				require.NoError(t, err)

				return ll.Replace(1, 2), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test replace a given value with new one two",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2, 5, 4)
				require.NoError(t, err)

				return ll.Replace(5, 3), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return error when item is not found in the list",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Replace(5, 3), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("element 5 not found in the list"),
		},
		{
			name: "test return error when old item has different type",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Replace('a', 3), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when new item has different type",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Replace(1, 'a'), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when new and old item has different type",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.Replace('a', 'b'), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when list is empty",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll.Replace(1, 2), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
			expectedError: errors.New("list is empty"),
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
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll.ReplaceAll(testIntIncOperator{}), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test replace all on string List with concat operator",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList("a", "b")
				require.NoError(t, err)

				return ll.ReplaceAll(testStringConcatOperator{}), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList("aa", "ba")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test replace all fails when operator return invalid data",
			actualResult: func() (error, List) {
				ll, err := NewLinkedList(1, 2)
				require.NoError(t, err)

				return ll.ReplaceAll(testInvalidOperator{}), ll
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2)
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
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RetainAll(2, 4)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test retain all from integer list with ignoring element which are not present",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RetainAll(2, 4, 5)
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test retain all from string List",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList("a", "b", "c", "d")
				require.NoError(t, err)

				ok, err := ll.RetainAll("b", "d")
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList("b", "d")
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test retain all removes all when argument list is empty",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList("a", "b", "c", "d")
				require.NoError(t, err)

				ok, err := ll.RetainAll()
				require.True(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ll.typeURL = "string"
				return ll
			},
		},
		{
			name: "test return error when type mismatch",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ok, err := ll.RetainAll(2, "d")
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
		{
			name: "test retain all fails when list is empty",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ok, err := ll.RetainAll(1, 2)
				require.False(t, ok)

				return ll, err
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
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

func TestLinkedListSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (List, interface{}, error)
		expectedResult interface{}
		expectedList   func() List
		expectedError  error
	}{
		{
			name: "test set value at index 3",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.Set(3, 5)

				return ll, ele, err
			},
			expectedResult: 5,
			expectedList: func() List {
				ll, err := NewLinkedList(1, 2, 3, 5)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test set value at index 0",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.Set(0, 2)

				return ll, ele, err
			},
			expectedResult: 2,
			expectedList: func() List {
				ll, err := NewLinkedList(2, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test set value at index 1",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.Set(1, 4)

				return ll, ele, err
			},
			expectedResult: 4,
			expectedList: func() List {
				ll, err := NewLinkedList(1, 4, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to set value due to invalid index",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.Set(5, 3)

				return ll, ele, err
			},
			expectedError: liberror.NewIndexOutOfBoundError(5),
			expectedList: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test set fails when list is empty",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				ele, err := ll.Set(0, 1)

				return ll, ele, err
			},
			expectedError: errors.New("list is empty"),
			expectedList: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test failed to set value due to different element type",
			actualResult: func() (List, interface{}, error) {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				ele, err := ll.Set(1, "a")

				return ll, ele, err
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
			expectedList: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			l, res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult, res)
			assert.Equal(t, testCase.expectedList(), l)
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
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(1, 4)
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist from index 0 to 0",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(0, 0)
			},
			expectedResult: func() List {
				ll, err := NewLinkedList()
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist from index 0 to 4",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(0, 4)
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist from index 4 to 6",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
				require.NoError(t, err)

				return ll.SubList(4, 6)
			},
			expectedResult: func() List {
				ll, err := NewLinkedList(5, 6)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test get sublist return error for invalid start index",
			actualResult: func() (List, error) {
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
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
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
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
				ll, err := NewLinkedList(1, 2, 3, 4, 5, 6, 7)
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
