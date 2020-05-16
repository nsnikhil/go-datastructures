package set

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/liberror"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewHashSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Set, error)
		expectedResult func() Set
		expectedError  error
	}{
		{
			name: "test create new empty hashset",
			actualResult: func() (Set, error) {
				return NewHashSet()
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap()
				require.NoError(t, err)

				return &HashSet{data: hm}
			},
		},
		{
			name: "test create new hashset with values",
			actualResult: func() (Set, error) {
				return NewHashSet(1, 2)
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap(gmap.NewPair(1, present{}), gmap.NewPair(2, present{}))
				require.NoError(t, err)

				return &HashSet{data: hm}
			},
		},
		{
			name: "test create with same values",
			actualResult: func() (Set, error) {
				return NewHashSet(1, 1, 1, 1)
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap(gmap.NewPair(1, present{}))
				require.NoError(t, err)

				return &HashSet{data: hm}
			},
		},
		{
			name: "test return error when elements are of different types",
			actualResult: func() (Set, error) {
				return NewHashSet(1, 2, 'a')
			},
			expectedResult: func() Set {
				return nil
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
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

func TestHashSetAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set)
		expectedResult func() Set
		expectedError  error
	}{
		{
			name: "test hash set add value to empty hashset",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.Add(1), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test hash set add value to non empty hashset",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(2)
				require.NoError(t, err)

				return hs.Add(1), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(2, 1)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test hash set add return error when element is of different type",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.Add('a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
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

func TestHashSetAddAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set)
		expectedResult func() Set
		expectedError  error
	}{
		{
			name: "test hashset add all values in empty set",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.AddAll(1, 2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test hashset add all values in non empty set",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(3, 4)
				require.NoError(t, err)

				return hs.AddAll(1, 2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(3, 4, 1, 2)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test return error when arguments types are not same",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.AddAll(1, 'a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when arguments types different from existing",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.AddAll('a', 'b'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when argument list is empty",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.AddAll(), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs
			},
			expectedError: errors.New("argument list is empty"),
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

func TestHashSetClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Set
		expectedResult func() Set
	}{
		{
			name: "test clear hash set",
			actualResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				hs.Clear()

				return hs
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap(gmap.NewPair(1, present{}))
				require.NoError(t, err)

				hm.Clear()

				return &HashSet{data: hm}
			},
		},
		{
			name: "test clear hash set two",
			actualResult: func() Set {
				hs, err := NewHashSet('a', 'b')
				require.NoError(t, err)

				hs.Clear()

				return hs
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap(gmap.NewPair('a', present{}), gmap.NewPair('b', present{}))
				require.NoError(t, err)

				hm.Clear()

				return &HashSet{data: hm}
			},
		},
		{
			name: "test clear empty hash set",
			actualResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				hs.Clear()

				return hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), testCase.actualResult())
		})
	}
}

func TestHashSetContains(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when element is present",
			actualResult: func() bool {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.Contains(2)
			},
			expectedResult: true,
		},
		{
			name: "test return false when element is not present",
			actualResult: func() bool {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.Contains(3)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestHashSetContainsAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when all elements are present",
			actualResult: func() bool {
				hs, err := NewHashSet(2, 4, 6)
				require.NoError(t, err)

				return hs.ContainsAll(2, 4, 6)
			},
			expectedResult: true,
		},
		{
			name: "test return false some elements are not present",
			actualResult: func() bool {
				hs, err := NewHashSet(2, 6)
				require.NoError(t, err)

				return hs.ContainsAll(2, 4, 6)
			},
		},
		{
			name: "test return false non of the elements are not present",
			actualResult: func() bool {
				hs, err := NewHashSet(2, 4, 6)
				require.NoError(t, err)

				return hs.ContainsAll(1, 3, 5)
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestHashSetCopy(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Set
		expectedResult func() Set
	}{
		{
			name: "test copy empty hash set",
			actualResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.Copy()
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test copy empty hash set with values",
			actualResult: func() Set {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.Copy()
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), testCase.actualResult())
		})
	}
}

func TestHashSetIsEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "return true when hashset is empty",
			actualResult: func() bool {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return false when hashset is not empty",
			actualResult: func() bool {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.IsEmpty()
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestHashSetSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "return size 0 when hashset is empty",
			actualResult: func() int {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.Size()
			},
			expectedResult: 0,
		},
		{
			name: "return size as 2 when hashset contains 2 elements",
			actualResult: func() int {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.Size()
			},
			expectedResult: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestHashSetRemove(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set)
		expectedError  error
		expectedResult func() Set
	}{
		{
			name: "test hashset remove element",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 2)
				require.NoError(t, err)

				return hs.Remove(2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				_, _ = hs.(*HashSet).data.Get(2)

				return hs
			},
		},
		{
			name: "test hashset remove element two",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.Remove(1), hs
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap(gmap.NewPair(1, present{}))
				require.NoError(t, err)

				_, err = hm.Remove(1)
				require.NoError(t, err)

				hs := &HashSet{data: hm}

				return hs
			},
		},
		{
			name: "test hashset remove return error when hash set is empty",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.Remove(2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
			expectedError: errors.New("set is empty"),
		},
		{
			name: "test hashset remove return error when type is different",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.Remove('a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
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

func TestHashSetRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set)
		expectedError  error
		expectedResult func() Set
	}{
		{
			name: "test hashset remove all elements",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(2, 4)
				require.NoError(t, err)

				return hs.RemoveAll(2, 4), hs
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap(gmap.NewPair(2, present{}), gmap.NewPair(4, present{}))
				require.NoError(t, err)

				_, err = hm.Remove(2)
				require.NoError(t, err)

				_, err = hm.Remove(4)
				require.NoError(t, err)

				hs := &HashSet{data: hm}

				return hs
			},
		},
		{
			name: "test hashset remove all removes some elements",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 2, 3, 4, 5, 6, 7, 8)
				require.NoError(t, err)

				return hs.RemoveAll(6, 3, 8, 2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 4, 5, 7)
				require.NoError(t, err)

				_, _ = hs.(*HashSet).data.Get(6)
				_, _ = hs.(*HashSet).data.Get(3)
				_, _ = hs.(*HashSet).data.Get(8)
				_, _ = hs.(*HashSet).data.Get(2)

				return hs
			},
		},
		{
			name: "test hashset remove all removes only element present in the set",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 3, 4, 6, 7, 9)
				require.NoError(t, err)

				return hs.RemoveAll(2, 4, 6, 8), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 3, 7, 9)
				require.NoError(t, err)

				_, _ = hs.(*HashSet).data.Get(2)
				_, _ = hs.(*HashSet).data.Get(4)
				_, _ = hs.(*HashSet).data.Get(6)
				_, _ = hs.(*HashSet).data.Get(8)

				return hs
			},
		},
		{
			name: "test hash remove all returns error when hash set is empty",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.RemoveAll(2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
			expectedError: errors.New("set is empty"),
		},
		{
			name: "test hash remove all returns error when argument list is empty",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.RemoveAll(), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: errors.New("argument list is empty"),
		},
		{
			name: "test hash remove all returns error when argument types are different",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.RemoveAll(2, 'a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test hash remove all returns error when type is different",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.RemoveAll('a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
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

func TestHashSetRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set)
		expectedError  error
		expectedResult func() Set
	}{
		{
			name: "test hashset retain all elements",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(2, 4)
				require.NoError(t, err)

				return hs.RetainAll(2, 4), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(2, 4)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test hashset retain all retain some elements",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 2, 3, 4, 5, 6, 7, 8)
				require.NoError(t, err)

				return hs.RetainAll(6, 3, 8, 2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(6, 3, 8, 2)
				require.NoError(t, err)

				_, _ = hs.(*HashSet).data.Get(1)
				_, _ = hs.(*HashSet).data.Get(4)
				_, _ = hs.(*HashSet).data.Get(7)
				_, _ = hs.(*HashSet).data.Get(5)

				return hs
			},
		},
		{
			name: "test hashset retain all retain no elements",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1, 2, 3, 4, 5, 6, 7, 8)
				require.NoError(t, err)

				return hs.RetainAll(), hs
			},
			expectedResult: func() Set {
				hm, err := gmap.NewHashMap()
				require.NoError(t, err)

				for i := 1; i <= 8; i++ {
					_, _ = hm.Put(i, present{})
				}

				_, _ = hm.Remove(1)
				_, _ = hm.Remove(6)
				_, _ = hm.Remove(2)
				_, _ = hm.Remove(3)
				_, _ = hm.Remove(4)
				_, _ = hm.Remove(8)
				_, _ = hm.Remove(7)
				_, _ = hm.Remove(5)

				return &HashSet{data: hm}
			},
		},
		{
			name: "test hashset retain all return error when hash set is empty",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs.RetainAll(2), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
			expectedError: errors.New("set is empty"),
		},
		{
			name: "test hashset retain all return error when argument types are different",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.RetainAll(1, 'a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: errors.New("all elements must be of same type"),
		},
		{
			name: "test hashset retain all return error when type is different",
			actualResult: func() (error, Set) {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs.RetainAll('a'), hs
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1)
				require.NoError(t, err)

				return hs
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
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

func TestHashSetIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []interface{}
		expectedResult []interface{}
	}{
		{
			name: "test hash set iterator",
			actualResult: func() []interface{} {
				res := make([]interface{}, 0)

				hs, err := NewHashSet(1, 2, 3, 4)
				require.NoError(t, err)

				it := hs.Iterator()
				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{1, 2, 3, 4},
		},
		{
			name: "test hash set iterator two",
			actualResult: func() []interface{} {
				res := make([]interface{}, 0)

				hs, err := NewHashSet('a', 'b')
				require.NoError(t, err)

				it := hs.Iterator()
				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{'a', 'b'},
		},
		{
			name: "test hash set iterator for empty set",
			actualResult: func() []interface{} {
				res := make([]interface{}, 0)

				hs, err := NewHashSet()
				require.NoError(t, err)

				it := hs.Iterator()
				for it.HasNext() {
					res = append(res, it.Next())
				}

				return res
			},
			expectedResult: []interface{}{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestHashSetUnion(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Set, error)
		expectedResult func() Set
		expectedError  error
	}{
		{
			name: "return union of two sets",
			actualResult: func() (Set, error) {
				hsa, err := NewHashSet(1, 2, 3, 4)
				require.NoError(t, err)

				hsb, err := NewHashSet(5, 6, 7, 8)
				require.NoError(t, err)

				return hsa.Union(hsb)
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 6, 2, 3, 4, 8, 7, 5)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "return union of multiple sets",
			actualResult: func() (Set, error) {
				hsa := make([]Set, 0)

				k := 1
				for i := 0; i < 10; i++ {
					a, err := NewHashSet()
					require.NoError(t, err)

					for j := 0; j < 10; j++ {
						err := a.Add(j + k)
						require.NoError(t, err)
					}
					hsa = append(hsa, a)
					k += 10
				}

				hs, err := NewHashSet()
				require.NoError(t, err)

				for _, s := range hsa {
					th, err := hs.Union(s)
					require.NoError(t, err)

					hs = th
				}

				return hs, nil
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 6, 18, 34, 97, 2, 12, 15, 22, 26, 31, 32, 36, 3, 10, 70, 95, 4, 8, 11, 13, 29, 33, 37, 47, 86, 9, 17, 20, 67, 7, 55, 62, 88, 19, 68, 46, 99, 41, 92, 23, 27, 16, 51, 84, 14, 45, 94, 43, 77, 79, 57, 39, 78, 21, 48, 80, 24, 35, 75, 44, 91, 28, 50, 30, 56, 38, 25, 42, 49, 52, 66, 72, 65, 64, 60, 69, 81, 98, 90, 58, 82, 85, 83, 100, 53, 73, 93, 63, 87, 59, 71, 61, 76, 96, 40, 89, 5, 54, 74)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "return error when two sets elements are different type",
			actualResult: func() (Set, error) {
				hsa, err := NewHashSet(1, 2)
				require.NoError(t, err)

				hsb, err := NewHashSet('a', 'b')
				require.NoError(t, err)

				return hsa.Union(hsb)
			},
			expectedResult: func() Set {
				return nil
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			toMap := func(s Set) map[interface{}]bool {
				tm := make(map[interface{}]bool)
				it := s.Iterator()
				for it.HasNext() {
					tm[it.Next()] = true
				}

				return tm
			}

			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)

			if res != nil {
				expectedMap := toMap(res)

				it := testCase.expectedResult().Iterator()
				for it.HasNext() {
					require.True(t, expectedMap[it.Next()])
				}
			}
		})
	}
}

func TestHashSetIntersection(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Set, error)
		expectedResult func() Set
		expectedError  error
	}{
		{
			name: "test hash set intersection",
			actualResult: func() (Set, error) {
				a, err := NewHashSet(1, 2, 4, 5, 7, 8)
				require.NoError(t, err)

				b, err := NewHashSet(1, 3, 5, 7)
				require.NoError(t, err)

				return a.Intersection(b)
			},
			expectedResult: func() Set {
				hs, err := NewHashSet(1, 7, 5)
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test hash set intersection return empty when one is empty",
			actualResult: func() (Set, error) {
				a, err := NewHashSet(1, 2, 4, 5, 7, 8)
				require.NoError(t, err)

				b, err := NewHashSet()
				require.NoError(t, err)

				return a.Intersection(b)
			},
			expectedResult: func() Set {
				hs, err := NewHashSet()
				require.NoError(t, err)

				return hs
			},
		},
		{
			name: "test hash set intersection return error when type is not same",
			actualResult: func() (Set, error) {
				a, err := NewHashSet(1, 2)
				require.NoError(t, err)

				b, err := NewHashSet('a')
				require.NoError(t, err)

				return a.Intersection(b)
			},
			expectedResult: func() Set {
				return nil
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
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
