package set

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/internal"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewHashSet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Set[int]
		expectedResult func() Set[int]
		expectedError  error
	}{
		{
			name: "test create new empty hashset",
			actualResult: func() Set[int] {
				return NewHashSet[int]()
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap[int, present]()

				return &HashSet[int]{data: hm}
			},
		},
		{
			name: "test create new hashset with values",
			actualResult: func() Set[int] {
				return NewHashSet[int](1, 2)
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap(gmap.NewPair[int, present](1, present{}), gmap.NewPair[int, present](2, present{}))
				return &HashSet[int]{data: hm}
			},
		},
		{
			name: "test create with same values",
			actualResult: func() Set[int] {
				return NewHashSet[int](1, 1, 1, 1)
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap(gmap.NewPair[int, present](1, present{}))

				return &HashSet[int]{data: hm}
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

func TestHashSetAdd(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Set[int]
		expectedResult func() Set[int]
	}{
		{
			name: "test hash set add value to empty hashset",
			actualResult: func() Set[int] {
				hs := NewHashSet[int]()

				hs.Add(1)
				return hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1)

				return hs
			},
		},
		{
			name: "test hash set add value to non empty hashset",
			actualResult: func() Set[int] {
				hs := NewHashSet[int](2)

				hs.Add(1)
				return hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](2, 1)

				return hs
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

func TestHashSetAddAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Set[int]
		expectedResult func() Set[int]
	}{
		{
			name: "test hashset add all values in empty set",
			actualResult: func() Set[int] {
				hs := NewHashSet[int]()

				hs.AddAll(1, 2)
				return hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 2)

				return hs
			},
		},
		{
			name: "test hashset add all values in non empty set",
			actualResult: func() Set[int] {
				hs := NewHashSet[int](3, 4)

				hs.AddAll(1, 2)
				return hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](3, 4, 1, 2)

				return hs
			},
		},
		{
			name: "test return error when argument list is empty",
			actualResult: func() Set[int] {
				hs := NewHashSet[int](1, 2)

				hs.AddAll()
				return hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 2)

				return hs
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

func TestHashSetClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Set[int]
		expectedResult func() Set[int]
	}{
		{
			name: "test clear hash set",
			actualResult: func() Set[int] {
				hs := NewHashSet[int](1)

				hs.Clear()

				return hs
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap[int, present](gmap.NewPair[int, present](1, present{}))

				hm.Clear()

				return &HashSet[int]{data: hm}
			},
		},
		{
			name: "test clear hash set two",
			actualResult: func() Set[int] {
				hs := NewHashSet[int]('a', 'b')

				hs.Clear()

				return hs
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap[int, present](gmap.NewPair[int, present]('a', present{}), gmap.NewPair[int, present]('b', present{}))

				hm.Clear()

				return &HashSet[int]{data: hm}
			},
		},
		{
			name: "test clear empty hash set",
			actualResult: func() Set[int] {
				hs := NewHashSet[int]()

				hs.Clear()

				return hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int]()

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
				hs := NewHashSet[int](1, 2)

				return hs.Contains(2)
			},
			expectedResult: true,
		},
		{
			name: "test return false when element is not present",
			actualResult: func() bool {
				hs := NewHashSet[int](1, 2)

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
				hs := NewHashSet[int](2, 4, 6)

				return hs.ContainsAll(2, 4, 6)
			},
			expectedResult: true,
		},
		{
			name: "test return false some elements are not present",
			actualResult: func() bool {
				hs := NewHashSet[int](2, 6)

				return hs.ContainsAll(2, 4, 6)
			},
		},
		{
			name: "test return false non of the elements are not present",
			actualResult: func() bool {
				hs := NewHashSet[int](2, 4, 6)

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
		actualResult   func() Set[int]
		expectedResult func() Set[int]
	}{
		{
			name: "test copy empty hash set",
			actualResult: func() Set[int] {
				hs := NewHashSet[int]()

				return hs.Copy()
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int]()

				return hs
			},
		},
		{
			name: "test copy empty hash set with values",
			actualResult: func() Set[int] {
				hs := NewHashSet[int](1, 2)

				return hs.Copy()
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 2)

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
				hs := NewHashSet[int]()

				return hs.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "return false when hashset is not empty",
			actualResult: func() bool {
				hs := NewHashSet[int](1)

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
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "return size 0 when hashset is empty",
			actualResult: func() int64 {
				hs := NewHashSet[int]()

				return hs.Size()
			},
			expectedResult: 0,
		},
		{
			name: "return size as 2 when hashset contains 2 elements",
			actualResult: func() int64 {
				hs := NewHashSet[int](1, 2)

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
		actualResult   func() (error, Set[int])
		expectedError  error
		expectedResult func() Set[int]
	}{
		{
			name: "test hashset remove element",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1, 2)

				return hs.Remove(2), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1)

				_, _ = hs.data.Get(2)

				return hs
			},
		},
		{
			name: "test hashset remove element two",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1)

				return hs.Remove(1), hs
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap[int, present](gmap.NewPair[int, present](1, present{}))

				_, err := hm.Remove(1)
				require.NoError(t, err)

				hs := &HashSet[int]{data: hm}

				return hs
			},
		},
		{
			name: "test hashset remove return error when hash set is empty",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int]()

				return hs.Remove(2), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int]()

				return hs
			},
			expectedError: errors.New("set is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashSetRemoveAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set[int])
		expectedError  error
		expectedResult func() Set[int]
	}{
		{
			name: "test hashset remove all elements",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](2, 4)

				return hs.RemoveAll(2, 4), hs
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap[int, present](gmap.NewPair[int, present](2, present{}), gmap.NewPair[int, present](4, present{}))

				_, err := hm.Remove(2)
				require.NoError(t, err)

				_, err = hm.Remove(4)
				require.NoError(t, err)

				hs := &HashSet[int]{data: hm}

				return hs
			},
		},
		{
			name: "test hashset remove all removes some elements",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1, 2, 3, 4, 5, 6, 7, 8)

				return hs.RemoveAll(6, 3, 8, 2), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 4, 5, 7)

				_, _ = hs.data.Get(6)
				_, _ = hs.data.Get(3)
				_, _ = hs.data.Get(8)
				_, _ = hs.data.Get(2)

				return hs
			},
		},
		{
			name: "test hashset remove all removes only element present in the set",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1, 3, 4, 6, 7, 9)

				return hs.RemoveAll(2, 4, 6, 8), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 3, 7, 9)

				_, _ = hs.data.Get(2)
				_, _ = hs.data.Get(4)
				_, _ = hs.data.Get(6)
				_, _ = hs.data.Get(8)

				return hs
			},
		},
		{
			name: "test hash remove all returns error when hash set is empty",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int]()

				return hs.RemoveAll(2), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int]()

				return hs
			},
			expectedError: errors.New("set is empty"),
		},
		{
			name: "test hash remove all returns error when argument list is empty",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1)

				return hs.RemoveAll(), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1)

				return hs
			},
			expectedError: errors.New("argument list is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashSetRetainAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Set[int])
		expectedError  error
		expectedResult func() Set[int]
	}{
		{
			name: "test hashset retain all elements",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](2, 4)

				return hs.RetainAll(2, 4), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](2, 4)

				return hs
			},
		},
		{
			name: "test hashset retain all retain some elements",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1, 2, 3, 4, 5, 6, 7, 8)

				return hs.RetainAll(6, 3, 8, 2), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](6, 3, 8, 2)

				_, _ = hs.data.Get(1)
				_, _ = hs.data.Get(4)
				_, _ = hs.data.Get(7)
				_, _ = hs.data.Get(5)

				return hs
			},
		},
		{
			name: "test hashset retain all retain no elements",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int](1, 2, 3, 4, 5, 6, 7, 8)

				return hs.RetainAll(), hs
			},
			expectedResult: func() Set[int] {
				hm := gmap.NewHashMap[int, present]()

				for i := 1; i <= 8; i++ {
					_ = hm.Put(i, present{})
				}

				_, _ = hm.Remove(1)
				_, _ = hm.Remove(6)
				_, _ = hm.Remove(2)
				_, _ = hm.Remove(3)
				_, _ = hm.Remove(4)
				_, _ = hm.Remove(8)
				_, _ = hm.Remove(7)
				_, _ = hm.Remove(5)

				return &HashSet[int]{data: hm}
			},
		},
		{
			name: "test hashset retain all return error when hash set is empty",
			actualResult: func() (error, Set[int]) {
				hs := NewHashSet[int]()

				return hs.RetainAll(2), hs
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int]()

				return hs
			},
			expectedError: errors.New("set is empty"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashSetIterator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []int
		expectedResult []int
	}{
		{
			name: "test hash set iterator",
			actualResult: func() []int {
				res := make([]int, 0)

				hs := NewHashSet[int](1, 2, 3, 4)

				it := hs.Iterator()
				for it.HasNext() {
					v, _ := it.Next()
					res = append(res, v)
				}

				return res
			},
			expectedResult: []int{1, 2, 3, 4},
		},
		{
			name: "test hash set iterator two",
			actualResult: func() []int {
				res := make([]int, 0)

				hs := NewHashSet[int]('a', 'b')

				it := hs.Iterator()
				for it.HasNext() {
					v, _ := it.Next()
					res = append(res, v)
				}

				return res
			},
			expectedResult: []int{'a', 'b'},
		},
		{
			name: "test hash set iterator for empty set",
			actualResult: func() []int {
				res := make([]int, 0)

				hs := NewHashSet[int]()

				it := hs.Iterator()
				for it.HasNext() {
					v, _ := it.Next()
					res = append(res, v)
				}

				return res
			},
			expectedResult: []int{},
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
		actualResult   func() (Set[int], error)
		expectedResult func() Set[int]
		expectedError  error
	}{
		{
			name: "return union of two sets",
			actualResult: func() (Set[int], error) {
				hsa := NewHashSet[int](1, 2, 3, 4)

				hsb := NewHashSet[int](5, 6, 7, 8)

				return hsa.Union(hsb)
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 6, 2, 3, 4, 8, 7, 5)

				return hs
			},
		},
		{
			name: "return union of multiple sets",
			actualResult: func() (Set[int], error) {
				hsa := make([]Set[int], 0)

				k := 1
				for i := 0; i < 10; i++ {
					a := NewHashSet[int]()

					for j := 0; j < 10; j++ {
						a.Add(j + k)
					}
					hsa = append(hsa, a)
					k += 10
				}

				hs := NewHashSet[int]()

				for _, s := range hsa {
					th, err := hs.Union(s)
					require.NoError(t, err)

					hs = th.(*HashSet[int])
				}

				return hs, nil
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 6, 18, 34, 97, 2, 12, 15, 22, 26, 31, 32, 36, 3, 10, 70, 95, 4, 8, 11, 13, 29, 33, 37, 47, 86, 9, 17, 20, 67, 7, 55, 62, 88, 19, 68, 46, 99, 41, 92, 23, 27, 16, 51, 84, 14, 45, 94, 43, 77, 79, 57, 39, 78, 21, 48, 80, 24, 35, 75, 44, 91, 28, 50, 30, 56, 38, 25, 42, 49, 52, 66, 72, 65, 64, 60, 69, 81, 98, 90, 58, 82, 85, 83, 100, 53, 73, 93, 63, 87, 59, 71, 61, 76, 96, 40, 89, 5, 54, 74)

				return hs
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			toMap := func(s Set[int]) map[int]bool {
				tm := make(map[int]bool)
				it := s.Iterator()
				for it.HasNext() {
					v, _ := it.Next()
					tm[v] = true
				}

				return tm
			}

			res, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)

			if res != nil {
				expectedMap := toMap(res)

				it := testCase.expectedResult().Iterator()
				for it.HasNext() {
					v, _ := it.Next()
					require.True(t, expectedMap[v])
				}
			}
		})
	}
}

func TestHashSetIntersection(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Set[int], error)
		expectedResult func() Set[int]
		expectedError  error
	}{
		{
			name: "test hash set intersection",
			actualResult: func() (Set[int], error) {
				a := NewHashSet[int](1, 2, 4, 5, 7, 8)

				b := NewHashSet[int](1, 3, 5, 7)

				return a.Intersection(b)
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int](1, 7, 5)

				return hs
			},
		},
		{
			name: "test hash set intersection return empty when one is empty",
			actualResult: func() (Set[int], error) {
				a := NewHashSet[int](1, 2, 4, 5, 7, 8)

				b := NewHashSet[int]()

				return a.Intersection(b)
			},
			expectedResult: func() Set[int] {
				hs := NewHashSet[int]()

				return hs
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}
