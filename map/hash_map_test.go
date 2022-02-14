package gmap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
	"strconv"
	"testing"
)

func TestCreateNewHashMap(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Map[int64, int64]
		expectedResult func() Map[int64, int64]
		expectedError  error
	}{
		{
			name: "test create empty hashmap",
			actualResult: func() Map[int64, int64] {
				return NewHashMap[int64, int64]()
			},
			expectedResult: func() Map[int64, int64] {
				return &HashMap[int64, int64]{
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 0, countMap: make(map[int64]bool), uniqueCount: 0},
					data:    make([]*list.LinkedList[*Pair[int64, int64]], 16),
					h:       sha3.New512(),
				}
			},
		},
		{
			name: "test create hashmap with args",
			actualResult: func() Map[int64, int64] {
				return NewHashMap[int64, int64](sliceToPair(internal.SliceGenerator{Size: 10}.Generate())...)
			},
			expectedResult: func() Map[int64, int64] {
				nwl := make([]*list.LinkedList[*Pair[int64, int64]], 16)

				hs := sha3.New512()

				data := internal.SliceGenerator{Size: 10}.Generate()

				cp := make(map[int64]bool)
				cpd := []int64{0, 1, 2, 3, 4, 5, 7, 9, 13}
				for _, k := range cpd {
					cp[k] = true
				}

				for _, e := range data {
					idx, err := indexOf(&hs, e, 16)
					require.NoError(t, err)

					ll := nwl[idx]

					if ll == nil {
						tll := list.NewLinkedList[*Pair[int64, int64]]()

						nwl[idx] = tll
					}

					nwl[idx].AddLast(NewPair[int64, int64](e, e))

				}

				return &HashMap[int64, int64]{
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 10, countMap: cp, uniqueCount: 9},
					h:       hs,
					data:    nwl,
				}

			},
		},
		{
			name: "test create hashmap with args of same keys",
			actualResult: func() Map[int64, int64] {
				res := NewHashMap[int64, int64](
					append(
						sliceToPair(internal.SliceGenerator{Size: 24}.Generate()),
						sliceToPair(internal.SliceGenerator{Size: 24}.Generate())...,
					)...,
				)

				return res
			},
			expectedResult: func() Map[int64, int64] {
				nwl := make([]*list.LinkedList[*Pair[int64, int64]], 32)

				hs := sha3.New512()

				data := internal.SliceGenerator{Size: 24}.Generate()

				cp := make(map[int64]bool)
				cpd := []int64{0, 1, 2, 3, 4, 5, 6, 9, 13, 14, 15, 17, 18, 19, 20, 21, 23, 29}
				for _, k := range cpd {
					cp[k] = true
				}

				var c int64 = 32
				tm := make(map[int64]bool)
				uc := 0
				var lf float64 = 0.75
				var incF float64 = 2

				for _, e := range data {
					if uc >= int(float64(c)*lf) {
						c = int64(float64(c) * incF)
					}

					idx, err := indexOf(&hs, e, c)
					require.NoError(t, err)

					if !tm[idx] {
						tm[idx] = true
						uc++
					}

					if nwl[idx] == nil {
						ll := list.NewLinkedList[*Pair[int64, int64]]()

						nwl[idx] = ll
					}

					nwl[idx].AddLast(NewPair[int64, int64](e, e))

				}

				return &HashMap[int64, int64]{
					factors: &factors{capacity: 32, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 24, countMap: cp, uniqueCount: 18},
					h:       hs,
					data:    nwl,
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			e := testCase.expectedResult()
			assert.Equal(t, e, res)
		})
	}
}

func TestHashMapPut(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (rune, Map[int, rune])
		expectedElement rune
		expectedResult  func() Map[int, rune]
	}{
		{
			name: "test put value in empty map",
			actualResult: func() (rune, Map[int, rune]) {
				m := NewHashMap[int, rune]()

				e := m.Put(1, 'a')
				return e, m
			},
			expectedResult: func() Map[int, rune] {
				res := NewHashMap[int, rune](NewPair[int, rune](1, 'a'))

				return res
			},
		},
		{
			name: "test put value in non empty map",
			actualResult: func() (rune, Map[int, rune]) {
				m := NewHashMap(NewPair[int, rune](1, 'a'))

				e := m.Put(2, 'b')
				return e, m
			},
			expectedResult: func() Map[int, rune] {
				res := NewHashMap[int, rune](NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'))

				return res
			},
		},
		{
			name: "test put value for existing key",
			actualResult: func() (rune, Map[int, rune]) {
				m := NewHashMap(NewPair[int, rune](1, 'a'))

				e := m.Put(1, 'b')
				return e, m
			},
			expectedResult: func() Map[int, rune] {
				res := NewHashMap[int, rune](NewPair[int, rune](1, 'b'))

				return res
			},
			expectedElement: 'a',
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapPutAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Map[int, rune]
		expectedError  error
		expectedResult func() Map[int, rune]
	}{
		{
			name: "put all values in empty map",
			actualResult: func() Map[int, rune] {
				hm := NewHashMap[int, rune]()

				hm.PutAll(NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'))

				return hm
			},
			expectedResult: func() Map[int, rune] {
				res := NewHashMap[int, rune](NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'))

				return res
			},
		},
		{
			name: "put all values in non empty map",
			actualResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'))

				hm.PutAll(NewPair[int, rune](3, 'c'), NewPair[int, rune](4, 'd'))

				return hm
			},
			expectedResult: func() Map[int, rune] {
				return NewHashMap[int, rune](
					NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'),
					NewPair[int, rune](3, 'c'), NewPair[int, rune](4, 'd'),
				)
			},
		},
		{
			name: "put all values replace existing value for same key",
			actualResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'), NewPair[int, rune](3, 'c'))

				hm.PutAll(NewPair[int, rune](1, 'd'), NewPair[int, rune](3, 'e'))

				return hm
			},
			expectedResult: func() Map[int, rune] {
				res := NewHashMap[int, rune](
					NewPair[int, rune](1, 'd'), NewPair[int, rune](2, 'b'), NewPair[int, rune](3, 'e'),
				)

				return res
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

func TestHashMapIterator(t *testing.T) {

	testCases := []struct {
		name           string
		actualResult   func() []*Pair[int64, int64]
		expectedResult func() []*Pair[int64, int64]
	}{
		{
			name: "iterate map to get all elements",
			actualResult: func() []*Pair[int64, int64] {
				hm := NewHashMap[int64, int64](sliceToPair(internal.SliceGenerator{Size: 26}.Generate())...)

				res := make([]*Pair[int64, int64], 26)
				i := 0

				it := hm.Iterator()
				for it.HasNext() {
					res[i], _ = it.Next()
					i++
				}

				return res
			},
			expectedResult: func() []*Pair[int64, int64] {
				return sliceToPair(internal.SliceGenerator{Size: 26}.Generate())
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			toMap := func(p []*Pair[int64, int64]) map[int64]int64 {
				res := make(map[int64]int64)
				for _, p := range p {
					res[p.first] = p.second
				}
				return res
			}

			eRes := toMap(testCase.expectedResult())
			res := toMap(testCase.actualResult())

			assert.Equal(t, len(eRes), len(res))

			for k, v := range res {
				assert.Equal(t, eRes[k], v)
			}
		})
	}
}

func TestHashMapGet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (rune, error)
		expectedResult rune
		expectedError  error
	}{
		{
			name: "test hash map get value",
			actualResult: func() (rune, error) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				return hm.Get(2)
			},
			expectedResult: 4,
		},
		{
			name: "test hash map get value after replace",
			actualResult: func() (rune, error) {
				hm := NewHashMap(NewPair[int, rune](1, 'a'), NewPair[int, rune](2, 'b'))

				hm.Put(2, 'c')

				return hm.Get(2)
			},
			expectedResult: 'c',
		},
		{
			name: "test return error indexOf has bucket but element is not present",
			actualResult: func() (rune, error) {
				hm := NewHashMap(NewPair[int, rune](3, 'a'))

				return hm.Get(10)
			},
			expectedError: errors.New("key 10 not found in the map"),
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

func TestHashMapGetOrDefault(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() rune
		expectedResult rune
	}{
		{
			name: "test return value when found in map",
			actualResult: func() rune {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				return hm.GetOrDefault(2, -1)
			},
			expectedResult: 4,
		},
		{
			name: "test return default value when value is not found",
			actualResult: func() rune {
				hm := NewHashMap(NewPair[int, rune](1, 'a'))

				return hm.GetOrDefault(2, 'd')
			},
			expectedResult: 'd',
		},
		//TODO: WHAT IS THE TEST BELOW?
		//{
		//	name: "test return error indexOf has bucket but element is not present",
		//	actualResult: func() rune {
		//		hm := NewHashMap(NewPair[int, rune](3, 'a'))
		//
		//		return hm.GetOrDefault(10, 'b')
		//	},
		//	expectedResult: 'b',
		//},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestHashMapRemove(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (rune, error, Map[int, rune])
		expectedResult  func() Map[int, rune]
		expectedElement rune
		expectedError   error
	}{
		{
			name: "test remove value against the given key",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				e, err := hm.Remove(2)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, _ = hm.Get(2)

				return hm
			},
			expectedElement: 4,
		},
		{
			name: "test remove value reduce capacity",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap[int, rune]()

				for i := 0; i < 22; i++ {
					hm.Put(i, int32(i+97))
				}

				var e rune
				var err error
				for i := 0; i < 11; i++ {
					e, err = hm.Remove(i)
				}

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap[int, rune]()

				for i := 11; i < 22; i++ {
					hm.Put(i, int32(i+97))
				}

				_, _ = hm.Get(10)

				return hm
			},
			expectedElement: int32(107),
		},
		{
			name: "test failed to remove element when key is not present",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				e, err := hm.Remove(4)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("key 4 not found in the map"),
		},
		{
			name: "test failed to remove element when key is not present two",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](3, 'c'))

				e, err := hm.Remove(10)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](3, 'c'))

				_, _ = hm.Get(10)

				return hm
			},
			expectedError: errors.New("key 10 not found in the map"),
		},
		{
			name: "test failed to remove element when bucket dose not contain the key",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](3, 4))

				e, err := hm.Remove(10)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](3, 4))

				_, _ = hm.Get(10)

				return hm
			},
			expectedError: errors.New("key 10 not found in the map"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapRemoveWithVal(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (rune, error, Map[int, rune])
		expectedResult  func() Map[int, rune]
		expectedElement rune
		expectedError   error
	}{
		{
			name: "test remove value against the given key and value",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				e, err := hm.RemoveWithVal(2, 4)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				d := make([]*list.LinkedList[*Pair[int, rune]], 16)
				hs := sha3.New512()

				values := []*Pair[int, rune]{NewPair[int, rune](1, 2), NewPair[int, rune](2, 4)}

				for _, value := range values {
					idx, err := indexOf(&hs, value.first, 16)
					require.NoError(t, err)

					if d[idx] == nil {
						ll := list.NewLinkedList[*Pair[int, rune]]()

						d[idx] = ll
					}

					d[idx].AddLast(NewPair[int, rune](value.first, value.second))
				}

				d[2] = nil

				return &HashMap[int, rune]{
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int64]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedElement: 4,
		},
		{
			name: "test failed to remove element when not present",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				e, err := hm.RemoveWithVal(4, 5)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("key 4 not found in the map"),
		},
		{
			name: "test failed when value dose not match",
			actualResult: func() (rune, error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](3, 8))

				e, err := hm.RemoveWithVal(3, 4)

				return e, err, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](3, 8))

				_, _ = hm.Get(3)

				return hm
			},
			expectedError: errors.New("value mismatch: expected 4, got 8"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map[int, rune])
		expectedError  error
		expectedResult func() Map[int, rune]
	}{
		{
			name: "test hashmap replace value of a key",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				return hm.Replace(1, 4), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 4), NewPair[int, rune](2, 4))

				_, _ = hm.Get(1)

				return hm
			},
		},
		{
			name: "test hashmap replace return error when key is not present",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				return hm.Replace(2, 4), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, _ = hm.Get(2)

				return hm
			},
			expectedError: errors.New("key 2 not found in the map"),
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

func TestHashMapReplaceWithVal(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map[int, rune])
		expectedError  error
		expectedResult func() Map[int, rune]
	}{
		{
			name: "test hashmap replace with val given key value pair and new value",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				return hm.ReplaceWithVal(1, 2, 4), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 4), NewPair[int, rune](2, 4))

				_, _ = hm.Get(1)

				return hm
			},
		},
		{
			name: "test hashmap replace with val return error when key is not present",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				return hm.ReplaceWithVal(2, 4, 5), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, _ = hm.Get(2)

				return hm
			},
			expectedError: errors.New("key 2 not found in the map"),
		},
		{
			name: "test hashmap replace with val return error when value mismatch",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				return hm.ReplaceWithVal(1, 4, 5), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, _ = hm.Get(1)

				return hm
			},
			expectedError: errors.New("value mismatch: expected 2, got 4"),
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

type doubleIt struct{}

func (dt doubleIt) Apply(t int, u rune) rune {
	return int32(t * int(u))
}

type appendIt struct{}

func (at appendIt) Apply(t int, u rune) rune {
	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", t, u))
	return int32(res)
}

func TestHashMapReplaceAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map[int, rune])
		expectedError  error
		expectedResult func() Map[int, rune]
	}{
		{
			name: "test replace all multiply key and val",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](3, 6), NewPair[int, rune](4, 7))

				return hm.ReplaceAll(doubleIt{}), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](3, 18), NewPair[int, rune](4, 28))

				return hm
			},
		},
		{
			name: "test replace all append key and val",
			actualResult: func() (error, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](3, 6), NewPair[int, rune](4, 7))

				return hm.ReplaceAll(appendIt{}), hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](3, 36), NewPair[int, rune](4, 47))

				return hm
			},
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

func TestHashMapCompute(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (error, rune, Map[int, rune])
		expectedError   error
		expectedElement rune
		expectedResult  func() Map[int, rune]
	}{
		{
			name: "test hashmap compute for integer",
			actualResult: func() (error, rune, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				nv, err := hm.Compute(2, doubleIt{})

				return err, nv, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 8))

				_, _ = hm.Get(2)

				return hm
			},
			expectedElement: 8,
		},
		{
			name: "test hashmap compute for char",
			actualResult: func() (error, rune, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](3, 6), NewPair[int, rune](4, 7))

				nv, err := hm.Compute(3, appendIt{})

				return err, nv, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](3, 36), NewPair[int, rune](4, 7))

				_, _ = hm.Get(3)

				return hm
			},
			expectedElement: 36,
		},
		{
			name: "test hashmap compute return error when key is not found",
			actualResult: func() (error, rune, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 8))

				nv, err := hm.Compute(4, doubleIt{})

				return err, nv, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 8))

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("key 4 not found in the map"),
		},
		{
			name: "test hashmap compute return error bucket does not contain the key",
			actualResult: func() (error, rune, Map[int, rune]) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](3, 8))

				nv, err := hm.Compute(10, doubleIt{})

				return err, nv, hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](3, 8))

				_, _ = hm.Get(10)

				return hm
			},
			expectedError: errors.New("key 10 not found in the map"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, e, res := testCase.actualResult()

			internal.AssertErrorEquals(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapContainsKey(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test contain key returns true when key is present",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](3, 8))

				return hm.ContainsKey(1)
			},
			expectedResult: true,
		},
		{
			name: "test contain key returns false when key is not present",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](3, 8))

				return hm.ContainsKey(4)
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

func TestHashMapContainsValue(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test contain value returns true when value is present",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](3, 8))

				return hm.ContainsValue(8)
			},
			expectedResult: true,
		},
		{
			name: "test contain value returns false when value is not present",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](3, 8))

				return hm.ContainsValue(4)
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

func TestHashMapSize(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() int64
		expectedResult int64
	}{
		{
			name: "test return size 0 of empty hash map",
			actualResult: func() int64 {
				hm := NewHashMap[int, rune]()

				return hm.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return size 1 of empty hash map with 1 element",
			actualResult: func() int64 {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				return hm.Size()
			},
			expectedResult: 1,
		},
		{
			name: "test return size 0 after deleting the pair",
			actualResult: func() int64 {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				_, err := hm.Remove(1)
				require.NoError(t, err)

				return hm.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return size 2 after delete and put operations",
			actualResult: func() int64 {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				_, err := hm.Remove(1)
				require.NoError(t, err)

				hm.PutAll(NewPair[int, rune](3, 6), NewPair[int, rune](4, 8))

				_, err = hm.Remove(4)
				require.NoError(t, err)

				return hm.Size()
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

func TestHashMapKeys(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (list.List[int], error)
		expectedResult func() list.List[int]
		expectedError  error
	}{
		{
			name: "test return hash map keys as list",
			actualResult: func() (list.List[int], error) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				hm.PutAll(NewPair[int, rune](3, 6), NewPair[int, rune](4, 8))

				return hm.Keys()
			},
			expectedResult: func() list.List[int] {
				return list.NewArrayList[int](1, 2, 3, 4)
			},
		},
		{
			name: "test return hash map keys as list two",
			actualResult: func() (list.List[int], error) {
				hm := NewHashMap(NewPair[int, rune](2, 4))

				return hm.Keys()
			},
			expectedResult: func() list.List[int] {
				return list.NewArrayList[int](2)
			},
		},
		{
			name: "test return empty list when hashmap is empty",
			actualResult: func() (list.List[int], error) {
				hm := NewHashMap[int, rune]()

				return hm.Keys()
			},
			expectedResult: func() list.List[int] {
				return list.NewArrayList[int]()
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

func TestHashMapValues(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (list.List[rune], error)
		expectedResult func() list.List[rune]
		expectedError  error
	}{
		{
			name: "test return hash map values as list",
			actualResult: func() (list.List[rune], error) {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				hm.PutAll(NewPair[int, rune](3, 6), NewPair[int, rune](4, 8))

				return hm.Values()
			},
			expectedResult: func() list.List[rune] {
				return list.NewArrayList[rune](2, 4, 6, 8)
			},
		},
		{
			name: "test return hash map values as list two",
			actualResult: func() (list.List[rune], error) {
				hm := NewHashMap(NewPair[int, rune](2, 4))

				return hm.Values()
			},
			expectedResult: func() list.List[rune] {
				return list.NewArrayList[rune](4)
			},
		},
		{
			name: "test return empty list when hashmap is empty",
			actualResult: func() (list.List[rune], error) {
				hm := NewHashMap[int, rune]()

				return hm.Values()
			},
			expectedResult: func() list.List[rune] {
				return list.NewArrayList[rune]()
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

func TestHashMapClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Map[int, rune]
		expectedResult func() Map[int, rune]
	}{
		{
			name: "test hash map clear",
			actualResult: func() Map[int, rune] {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				hm.Clear()

				return hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap[int, rune]()

				return hm
			},
		},
		{
			name: "test hash map clear empty map",
			actualResult: func() Map[int, rune] {
				hm := NewHashMap[int, rune]()

				hm.Clear()

				return hm
			},
			expectedResult: func() Map[int, rune] {
				hm := NewHashMap[int, rune]()

				return hm
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult(), testCase.actualResult())
		})
	}
}

func TestHashMapIsEmpty(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when hashmap is empty",
			actualResult: func() bool {
				hm := NewHashMap[int, rune]()

				return hm.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "test return true when hashmap is empty after few operations",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				_, err := hm.Remove(2)
				require.NoError(t, err)

				hm.Put(3, 6)
				require.NoError(t, err)

				_, err = hm.Remove(1)
				require.NoError(t, err)

				_, err = hm.Remove(3)
				require.NoError(t, err)

				return hm.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "test return false when hashmap is not empty",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2))

				return hm.IsEmpty()
			},
			expectedResult: false,
		},
		{
			name: "test return false when hashmap is not empty after few operations",
			actualResult: func() bool {
				hm := NewHashMap(NewPair[int, rune](1, 2), NewPair[int, rune](2, 4))

				_, err := hm.Remove(2)
				require.NoError(t, err)

				hm.Put(3, 6)
				require.NoError(t, err)

				_, err = hm.Remove(3)
				require.NoError(t, err)

				return hm.IsEmpty()
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

func sliceToPair(data []int64) []*Pair[int64, int64] {
	sz := len(data)
	res := make([]*Pair[int64, int64], sz)
	for i := 0; i < sz; i++ {
		res[i] = NewPair[int64, int64](data[i], data[i])
	}
	return res
}
