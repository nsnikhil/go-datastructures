package gmap

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/liberror"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestCreateNewHashMap(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Map, error)
		expectedResult func() Map
		expectedError  error
	}{
		{
			name: "test create empty hashmap",
			actualResult: func() (Map, error) {
				return NewHashMap()
			},
			expectedResult: func() Map {

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "na", valueTypeURL: "na"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 0, countMap: make(map[int]bool), uniqueCount: 0},
					data:    make([]*list.LinkedList, 16),
					h:       sha3.New512(),
				}
			},
		},
		{
			name: "test create hashmap with args",
			actualResult: func() (Map, error) {
				return NewHashMap(NewPair(1, 'a'), NewPair(2, 'b'), NewPair(3, 'c'), NewPair(4, 'd'))
			},
			expectedResult: func() Map {
				nwl := make([]*list.LinkedList, 16)

				hs := sha3.New512()

				data := []struct {
					k interface{}
					v interface{}
				}{
					{1, 'a'}, {2, 'b'}, {3, 'c'}, {4, 'd'},
				}

				for _, e := range data {
					idx, err := indexOf(&hs, e.k, 16)
					require.NoError(t, err)

					ll := nwl[idx]

					if ll == nil {
						tll, err := list.NewLinkedList()
						require.NoError(t, err)

						nwl[idx] = tll
					}

					require.NoError(t, nwl[idx].AddLast(NewPair(e.k, e.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 4, countMap: map[int]bool{0: true, 2: true, 3: true, 4: true}, uniqueCount: 4},
					h:       hs,
					data:    nwl,
				}

			},
		},
		{
			name: "test create hashmap with args of same keys",
			actualResult: func() (Map, error) {
				return NewHashMap(NewPair(1, 'a'), NewPair(1, 'b'), NewPair(2, 'c'), NewPair(2, 'd'))
			},
			expectedResult: func() Map {
				nwl := make([]*list.LinkedList, 16)

				hs := sha3.New512()

				data := []struct {
					k interface{}
					v interface{}
				}{
					{1, 'b'}, {2, 'd'},
				}

				for _, e := range data {
					idx, err := indexOf(&hs, e.k, 16)
					require.NoError(t, err)

					ll := nwl[idx]

					if ll == nil {
						tll, err := list.NewLinkedList()
						require.NoError(t, err)

						nwl[idx] = tll
					}

					require.NoError(t, nwl[idx].AddLast(NewPair(e.k, e.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 2, countMap: map[int]bool{0: true, 2: true}, uniqueCount: 2},
					h:       hs,
					data:    nwl,
				}

			},
		},
		{
			name: "test create hashmap with args more than default capacity",
			actualResult: func() (Map, error) {

				res := make([]*Pair, 26)
				for i := 0; i < 26; i++ {
					res[i] = NewPair(i+1, int32(i+97))
				}

				return NewHashMap(res...)
			},
			expectedResult: func() Map {
				nwl := make([]*list.LinkedList, 32)

				hs := sha3.New512()

				type dt struct {
					k interface{}
					v interface{}
				}

				data := make([]dt, 26)

				for i := 0; i < 26; i++ {
					data[i] = dt{i + 1, int32(i + 97)}
				}

				c := 16
				tm := make(map[int]bool)
				uc := 0
				lf := 0.75
				incF := 2

				for _, e := range data {
					if uc >= int(float64(c)*lf) {
						c *= incF
					}

					idx, err := indexOf(&hs, e.k, float64(c))
					require.NoError(t, err)

					if !tm[idx] {
						tm[idx] = true
						uc++
					}

					if nwl[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						nwl[idx] = ll
					}

					require.NoError(t, nwl[idx].AddLast(NewPair(e.k, e.v)))

				}
				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: c, upperLoadFactor: lf, lowerLoadFactor: 0.40, scalingFactor: incF},
					counter: &counter{elementCount: 26, countMap: tm, uniqueCount: uc},
					h:       hs,
					data:    nwl,
				}

			},
		},
		{
			name: "test return error when value have different types",
			actualResult: func() (Map, error) {
				return NewHashMap(NewPair(1, 'a'), NewPair('b', 'c'))
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
			expectedResult: func() Map {
				return nil
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

func TestHashMapPut(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, Map)
		expectedError   error
		expectedElement interface{}
		expectedResult  func() Map
	}{
		{
			name: "test put value in empty map",
			actualResult: func() (interface{}, error, Map) {
				m, err := NewHashMap()
				require.NoError(t, err)

				e, err := m.Put(1, 'a')
				return e, err, m
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				idx, err := indexOf(&hs, 1, 16)
				require.NoError(t, err)
				if d[idx] == nil {
					ll, err := list.NewLinkedList()
					require.NoError(t, err)

					d[idx] = ll
				}

				require.NoError(t, d[idx].AddLast(NewPair(1, 'a')))

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
		},
		{
			name: "test put value in non empty map",
			actualResult: func() (interface{}, error, Map) {
				m, err := NewHashMap(NewPair(1, 'a'))
				require.NoError(t, err)

				e, err := m.Put(2, 'b')
				return e, err, m
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{NewPair(1, 'a'), NewPair(2, 'b')}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 2, countMap: map[int]bool{0: true, 2: true}, uniqueCount: 2},
					h:       hs,
					data:    d,
				}
			},
		},
		{
			name: "test put value for existing key",
			actualResult: func() (interface{}, error, Map) {
				m, err := NewHashMap(NewPair(1, 'a'))
				require.NoError(t, err)

				e, err := m.Put(1, 'b')
				return e, err, m
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{NewPair(1, 'b')}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedElement: 'a',
		},
		{
			name: "test return error when key type is different",
			actualResult: func() (interface{}, error, Map) {
				m, err := NewHashMap(NewPair(1, 'a'))
				require.NoError(t, err)

				e, err := m.Put('b', 'd')
				return e, err, m
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				idx, err := indexOf(&hs, 1, 16)
				require.NoError(t, err)
				if d[idx] == nil {
					ll, err := list.NewLinkedList()
					require.NoError(t, err)

					d[idx] = ll
				}

				require.NoError(t, d[idx].AddLast(NewPair(1, 'a')))

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "test return error when value type is different",
			actualResult: func() (interface{}, error, Map) {
				m, err := NewHashMap(NewPair(1, 'a'))
				require.NoError(t, err)

				e, err := m.Put(2, 4)
				return e, err, m
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				idx, err := indexOf(&hs, 1, 16)
				require.NoError(t, err)
				if d[idx] == nil {
					ll, err := list.NewLinkedList()
					require.NoError(t, err)

					d[idx] = ll
				}

				require.NoError(t, d[idx].AddLast(NewPair(1, 'a')))

				_, err = indexOf(&hs, 2, 16)
				require.NoError(t, err)

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedError: liberror.NewTypeMismatchError("int32", "int"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapPutAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map)
		expectedError  error
		expectedResult func() Map
	}{
		{
			name: "put all values in empty map",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.PutAll(NewPair(1, 'a'), NewPair(2, 'b')), hm
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{NewPair(1, 'a'), NewPair(2, 'b')}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 2, countMap: map[int]bool{0: true, 2: true}, uniqueCount: 2},
					h:       hs,
					data:    d,
				}
			},
		},
		{
			name: "put all values in non empty map",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 'a'), NewPair(2, 'b'))
				require.NoError(t, err)

				return hm.PutAll(NewPair(3, 'c'), NewPair(4, 'd')), hm
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{
					NewPair(1, 'a'), NewPair(2, 'b'), NewPair(3, 'c'), NewPair(4, 'd'),
				}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 4, countMap: map[int]bool{0: true, 2: true, 3: true, 4: true}, uniqueCount: 4},
					h:       hs,
					data:    d,
				}
			},
		},
		{
			name: "put all values replace existing value for same key",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 'a'), NewPair(2, 'b'), NewPair(3, 'c'))
				require.NoError(t, err)

				return hm.PutAll(NewPair(1, 'd'), NewPair(3, 'e')), hm
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{
					NewPair(1, 'd'), NewPair(2, 'b'), NewPair(3, 'e'),
				}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 3, countMap: map[int]bool{0: true, 2: true, 3: true}, uniqueCount: 3},
					h:       hs,
					data:    d,
				}
			},
		},
		{
			name: "return error when key type mismatch",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.PutAll(NewPair(1, 'a'), NewPair('b', 'd')), hm
			},
			expectedResult: func() Map {
				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "na", valueTypeURL: "na"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 0, countMap: map[int]bool{}, uniqueCount: 0},
					h:       sha3.New512(),
					data:    make([]*list.LinkedList, 16),
				}
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "return error when value type mismatch",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.PutAll(NewPair(1, 'a'), NewPair(2, 3)), hm
			},
			expectedResult: func() Map {
				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "na", valueTypeURL: "na"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 0, countMap: map[int]bool{}, uniqueCount: 0},
					h:       sha3.New512(),
					data:    make([]*list.LinkedList, 16),
				}
			},
			expectedError: liberror.NewTypeMismatchError("int32", "int"),
		},
		{
			name: "return error when key type mismatch with key present",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 'a'))
				require.NoError(t, err)

				return hm.PutAll(NewPair('a', 'b')), hm
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{NewPair(1, 'a')}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedError: liberror.NewTypeMismatchError("int", "int32"),
		},
		{
			name: "return error when value type mismatch with value present",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 'a'))
				require.NoError(t, err)

				return hm.PutAll(NewPair(2, 3)), hm
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{NewPair(1, 'a')}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)
					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int32"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedError: liberror.NewTypeMismatchError("int32", "int"),
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

func TestHashMapIterator(t *testing.T) {
	type testType struct{ i int }

	testCases := []struct {
		name           string
		actualResult   func() []*Pair
		expectedResult func() []*Pair
	}{
		{
			name: "iterate map to get 26 elements",
			actualResult: func() []*Pair {
				res := make([]*Pair, 26)
				for i := 0; i < 26; i++ {
					res[i] = NewPair(i+1, int32(i+97))
				}

				hm, err := NewHashMap(res...)
				require.NoError(t, err)

				res = make([]*Pair, 26)

				i := 0

				it := hm.Iterator()
				for it.HasNext() {
					res[i] = it.Next().(*Pair)
					i++
				}

				return res
			},
			expectedResult: func() []*Pair {
				res := make([]*Pair, 26)
				for i := 0; i < 26; i++ {
					res[i] = NewPair(i+1, int32(i+97))
				}

				return res
			},
		},
		{
			name: "iterate map to get 8 elements",
			actualResult: func() []*Pair {
				res := make([]*Pair, 8)
				for i := 0; i < 8; i++ {
					res[i] = NewPair(i, testType{i})
				}

				hm, err := NewHashMap(res...)
				require.NoError(t, err)

				res = make([]*Pair, 8)

				i := 0
				it := hm.Iterator()
				for it.HasNext() {
					res[i] = it.Next().(*Pair)
					i++
				}

				return res
			},
			expectedResult: func() []*Pair {
				res := make([]*Pair, 8)
				for i := 0; i < 8; i++ {
					res[i] = NewPair(i, testType{i})
				}

				return res
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			toMap := func(p []*Pair) map[interface{}]interface{} {
				res := make(map[interface{}]interface{})
				for _, p := range p {
					res[p.k] = p.v
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
		actualResult   func() (interface{}, error)
		expectedResult interface{}
		expectedError  error
	}{
		{
			name: "test hash map get value",
			actualResult: func() (interface{}, error) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				return hm.Get(2)
			},
			expectedResult: 4,
		},
		{
			name: "test hash map get value two",
			actualResult: func() (interface{}, error) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				err = hm.PutAll(NewPair('a', "abc"), NewPair('d', "def"))
				require.NoError(t, err)

				return hm.Get('a')
			},
			expectedResult: "abc",
		},
		{
			name: "test hash map get value after replace",
			actualResult: func() (interface{}, error) {
				hm, err := NewHashMap(NewPair(1, 'a'), NewPair(2, 'b'))
				require.NoError(t, err)

				_, err = hm.Put(2, 'c')
				require.NoError(t, err)

				return hm.Get(2)
			},
			expectedResult: 'c',
		},
		{
			name: "test return error when element is not found",
			actualResult: func() (interface{}, error) {
				hm, err := NewHashMap(NewPair('a', "abc"))
				require.NoError(t, err)

				return hm.Get(1)
			},
			expectedError: liberror.NewTypeMismatchError("int32", "int"),
		},
		{
			name: "test return error indexOf has bucket but element is not present",
			actualResult: func() (interface{}, error) {
				hm, err := NewHashMap(NewPair(3, 'a'))
				require.NoError(t, err)

				return hm.Get(10)
			},
			expectedError: errors.New("no value found against the key: 10"),
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

func TestHashMapGetOrDefault(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() interface{}
		expectedResult interface{}
	}{
		{
			name: "test return value when found in map",
			actualResult: func() interface{} {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				return hm.GetOrDefault(2, -1)
			},
			expectedResult: 4,
		},
		{
			name: "test return value when found in map two",
			actualResult: func() interface{} {
				hm, err := NewHashMap()
				require.NoError(t, err)

				err = hm.PutAll(NewPair('a', "abc"), NewPair('d', "def"))
				require.NoError(t, err)

				return hm.GetOrDefault('a', '/')
			},
			expectedResult: "abc",
		},
		{
			name: "test return value after replace",
			actualResult: func() interface{} {
				hm, err := NewHashMap(NewPair('a', "def"))
				require.NoError(t, err)

				err = hm.PutAll(NewPair('a', "abc"), NewPair('d', "efg"))
				require.NoError(t, err)

				return hm.GetOrDefault('a', '/')
			},
			expectedResult: "abc",
		},
		{
			name: "test return default value when value is not found",
			actualResult: func() interface{} {
				hm, err := NewHashMap(NewPair('a', "abc"))
				require.NoError(t, err)

				return hm.GetOrDefault(1, 'd')
			},
			expectedResult: 'd',
		},
		{
			name: "test return error indexOf has bucket but element is not present",
			actualResult: func() interface{} {
				hm, err := NewHashMap(NewPair(3, 'a'))
				require.NoError(t, err)

				return hm.GetOrDefault(10, 'b')
			},
			expectedResult: 'b',
		},
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
		actualResult    func() (interface{}, error, Map)
		expectedResult  func() Map
		expectedElement interface{}
		expectedError   error
	}{
		{
			name: "test remove value against the given key",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				e, err := hm.Remove(2)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, _ = hm.Get(2)

				return hm
			},
			expectedElement: 4,
		},
		{
			name: "test remove value reduce capacity",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				for i := 0; i < 22; i++ {
					_, err := hm.Put(i, int32(i+97))
					require.NoError(t, err)
				}

				var e interface{}
				for i := 0; i < 11; i++ {
					e, err = hm.Remove(i)
				}

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap()
				require.NoError(t, err)

				for i := 11; i < 22; i++ {
					_, err := hm.Put(i, int32(i+97))
					require.NoError(t, err)
				}

				_, _ = hm.Get(10)

				return hm
			},
			expectedElement: int32(107),
		},
		{
			name: "test failed to remove element when key is not present",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				e, err := hm.Remove(4)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("no value found against the key: 4"),
		},
		{
			name: "test failed to remove element when key is not present two",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(3, 'c'))
				require.NoError(t, err)

				e, err := hm.Remove(10)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(3, 'c'))
				require.NoError(t, err)

				_, _ = hm.Get(10)

				return hm
			},
			expectedError: errors.New("no value found against the key: 10"),
		},
		{
			name: "test failed to remove element when bucket dose not contain the key",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(3, 4))
				require.NoError(t, err)

				e, err := hm.Remove(10)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(3, 4))
				require.NoError(t, err)

				_, _ = hm.Get(10)

				return hm
			},
			expectedError: errors.New("no value found against the key: 10"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapRemoveWithVal(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (interface{}, error, Map)
		expectedResult  func() Map
		expectedElement interface{}
		expectedError   error
	}{
		{
			name: "test remove value against the given key and value",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				e, err := hm.RemoveWithVal(2, 4)

				return e, err, hm
			},
			expectedResult: func() Map {
				d := make([]*list.LinkedList, 16)
				hs := sha3.New512()

				values := []*Pair{NewPair(1, 2), NewPair(2, 4)}

				for _, value := range values {
					idx, err := indexOf(&hs, value.k, 16)
					require.NoError(t, err)

					if d[idx] == nil {
						ll, err := list.NewLinkedList()
						require.NoError(t, err)

						d[idx] = ll
					}

					require.NoError(t, d[idx].AddLast(NewPair(value.k, value.v)))
				}

				d[2] = nil

				return &HashMap{
					typeURL: &typeURL{keyTypeURL: "int", valueTypeURL: "int"},
					factors: &factors{capacity: 16, upperLoadFactor: 0.75, lowerLoadFactor: 0.40, scalingFactor: 2},
					counter: &counter{elementCount: 1, countMap: map[int]bool{0: true}, uniqueCount: 1},
					h:       hs,
					data:    d,
				}
			},
			expectedElement: 4,
		},
		{
			name: "test failed to remove element when not present",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				e, err := hm.RemoveWithVal(4, 5)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("no value found against the key: 4"),
		},
		{
			name: "test failed when key dose not match",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(3, 8))
				require.NoError(t, err)

				e, err := hm.RemoveWithVal(4, 2)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(3, 8))
				require.NoError(t, err)

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("no value found against the key: 4"),
		},
		{
			name: "test failed when value dose not match",
			actualResult: func() (interface{}, error, Map) {
				hm, err := NewHashMap(NewPair(3, 8))
				require.NoError(t, err)

				e, err := hm.RemoveWithVal(3, 4)

				return e, err, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(3, 8))
				require.NoError(t, err)

				_, _ = hm.Get(3)

				return hm
			},
			expectedError: errors.New("value mismatch: expected 4, got 8"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e, err, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedElement, e)
			assert.Equal(t, testCase.expectedResult(), res)
		})
	}
}

func TestHashMapReplace(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map)
		expectedError  error
		expectedResult func() Map
	}{
		{
			name: "test hashmap replace value of a key",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				return hm.Replace(1, 4), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 4), NewPair(2, 4))
				require.NoError(t, err)

				_, _ = hm.Get(1)

				return hm
			},
		},
		{
			name: "test hashmap replace return error when key is not present",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				return hm.Replace(2, 4), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, _ = hm.Get(2)

				return hm
			},
			expectedError: errors.New("no value found against the key: 2"),
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

func TestHashMapReplaceWithVal(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map)
		expectedError  error
		expectedResult func() Map
	}{
		{
			name: "test hashmap replace with val given key value pair and new value",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				return hm.ReplaceWithVal(1, 2, 4), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 4), NewPair(2, 4))
				require.NoError(t, err)

				_, _ = hm.Get(1)

				return hm
			},
		},
		{
			name: "test hashmap replace with val return error when key is not present",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				return hm.ReplaceWithVal(2, 4, 5), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, _ = hm.Get(2)

				return hm
			},
			expectedError: errors.New("no value found against the key: 2"),
		},
		{
			name: "test hashmap replace with val return error when value mismatch",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				return hm.ReplaceWithVal(1, 4, 5), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, _ = hm.Get(1)

				return hm
			},
			expectedError: errors.New("value mismatch: expected 2, got 4"),
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

type doubleIt struct{}

func (dt doubleIt) Apply(t interface{}, u interface{}) interface{} {
	return t.(int) * u.(int)
}

type appendIt struct{}

func (at appendIt) Apply(t interface{}, u interface{}) interface{} {
	return t.(int32) + u.(int32)
}

type invalidOp struct{}

func (iv invalidOp) Apply(t interface{}, u interface{}) interface{} {
	return fmt.Sprintf("%v%v", t, u)
}

func TestHashMapReplaceAll(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (error, Map)
		expectedError  error
		expectedResult func() Map
	}{
		{
			name: "test replace all multiply key and val",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(3, 6), NewPair(4, 7))
				require.NoError(t, err)

				return hm.ReplaceAll(doubleIt{}), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(3, 18), NewPair(4, 28))
				require.NoError(t, err)

				return hm
			},
		},
		{
			name: "test replace all append key and val",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair('0', 'A'), NewPair('1', 'B'))
				require.NoError(t, err)

				return hm.ReplaceAll(appendIt{}), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair('0', 'q'), NewPair('1', 's'))
				require.NoError(t, err)

				return hm
			},
		},
		{
			name: "test replace all return error when function return different type value",
			actualResult: func() (error, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				return hm.ReplaceAll(invalidOp{}), hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				return hm
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

func TestHashMapCompute(t *testing.T) {
	testCases := []struct {
		name            string
		actualResult    func() (error, interface{}, Map)
		expectedError   error
		expectedElement interface{}
		expectedResult  func() Map
	}{
		{
			name: "test hashmap compute for integer",
			actualResult: func() (error, interface{}, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				nv, err := hm.Compute(2, doubleIt{})

				return err, nv, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 8))
				require.NoError(t, err)

				_, _ = hm.Get(2)

				return hm
			},
			expectedElement: 8,
		},
		{
			name: "test hashmap compute for char",
			actualResult: func() (error, interface{}, Map) {
				hm, err := NewHashMap(NewPair('0', 'A'), NewPair('1', 'B'))
				require.NoError(t, err)

				nv, err := hm.Compute('0', appendIt{})

				return err, nv, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair('0', 'q'), NewPair('1', 'B'))
				require.NoError(t, err)

				_, _ = hm.Get('0')

				return hm
			},
			expectedElement: 'q',
		},
		{
			name: "test hashmap compute return error when key is not found",
			actualResult: func() (error, interface{}, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 8))
				require.NoError(t, err)

				nv, err := hm.Compute(4, doubleIt{})

				return err, nv, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 8))
				require.NoError(t, err)

				_, _ = hm.Get(4)

				return hm
			},
			expectedError: errors.New("no value found against the key: 4"),
		},
		{
			name: "test hashmap compute return error bucket does not contain the key",
			actualResult: func() (error, interface{}, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

				nv, err := hm.Compute(10, doubleIt{})

				return err, nv, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

				_, _ = hm.Get(10)

				return hm
			},
			expectedError: errors.New("no value found against the key: 10"),
		},
		{
			name: "test hashmap compute return function return type is different",
			actualResult: func() (error, interface{}, Map) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

				nv, err := hm.Compute(3, invalidOp{})

				return err, nv, hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

				_, _ = hm.Get(3)

				return hm
			},
			expectedError: liberror.NewTypeMismatchError("int", "string"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err, e, res := testCase.actualResult()

			assert.Equal(t, testCase.expectedError, err)
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
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

				return hm.ContainsKey(1)
			},
			expectedResult: true,
		},
		{
			name: "test contain key returns false when key is not present",
			actualResult: func() bool {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

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
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

				return hm.ContainsValue(8)
			},
			expectedResult: true,
		},
		{
			name: "test contain value returns false when value is not present",
			actualResult: func() bool {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(3, 8))
				require.NoError(t, err)

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
		actualResult   func() int
		expectedResult int
	}{
		{
			name: "test return size 0 of empty hash map",
			actualResult: func() int {
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return size 1 of empty hash map with 1 element",
			actualResult: func() int {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				return hm.Size()
			},
			expectedResult: 1,
		},
		{
			name: "test return size 0 after deleting the pair",
			actualResult: func() int {
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				_, err = hm.Remove(1)
				require.NoError(t, err)

				return hm.Size()
			},
			expectedResult: 0,
		},
		{
			name: "test return size 2 after delete and put operations",
			actualResult: func() int {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				_, err = hm.Remove(1)
				require.NoError(t, err)

				err = hm.PutAll(NewPair(3, 6), NewPair(4, 8))
				require.NoError(t, err)

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
		actualResult   func() (list.List, error)
		expectedResult func() list.List
		expectedError  error
	}{
		{
			name: "test return hash map keys as list",
			actualResult: func() (list.List, error) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				err = hm.PutAll(NewPair(3, 6), NewPair(4, 8))
				require.NoError(t, err)

				return hm.Keys()
			},
			expectedResult: func() list.List {
				ll, err := list.NewLinkedList(1, 2, 3, 4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return hash map keys as list two",
			actualResult: func() (list.List, error) {
				hm, err := NewHashMap(NewPair(2, 4))
				require.NoError(t, err)

				return hm.Keys()
			},
			expectedResult: func() list.List {
				ll, err := list.NewLinkedList(2)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return empty list when hashmap is empty",
			actualResult: func() (list.List, error) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.Keys()
			},
			expectedResult: func() list.List {
				ll, err := list.NewLinkedList()
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

func TestHashMapValues(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (list.List, error)
		expectedResult func() list.List
		expectedError  error
	}{
		{
			name: "test return hash map values as list",
			actualResult: func() (list.List, error) {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				err = hm.PutAll(NewPair(3, 6), NewPair(4, 8))
				require.NoError(t, err)

				return hm.Values()
			},
			expectedResult: func() list.List {
				ll, err := list.NewLinkedList(2, 4, 6, 8)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return hash map values as list two",
			actualResult: func() (list.List, error) {
				hm, err := NewHashMap(NewPair(2, 4))
				require.NoError(t, err)

				return hm.Values()
			},
			expectedResult: func() list.List {
				ll, err := list.NewLinkedList(4)
				require.NoError(t, err)

				return ll
			},
		},
		{
			name: "test return empty list when hashmap is empty",
			actualResult: func() (list.List, error) {
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.Values()
			},
			expectedResult: func() list.List {
				ll, err := list.NewLinkedList()
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

func TestHashMapClear(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() Map
		expectedResult func() Map
	}{
		{
			name: "test hash map clear",
			actualResult: func() Map {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				hm.Clear()

				return hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap()
				require.NoError(t, err)

				hm.(*HashMap).keyTypeURL = "int"
				hm.(*HashMap).valueTypeURL = "int"

				return hm
			},
		},
		{
			name: "test hash map clear empty map",
			actualResult: func() Map {
				hm, err := NewHashMap()
				require.NoError(t, err)

				hm.Clear()

				return hm
			},
			expectedResult: func() Map {
				hm, err := NewHashMap()
				require.NoError(t, err)

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
				hm, err := NewHashMap()
				require.NoError(t, err)

				return hm.IsEmpty()
			},
			expectedResult: true,
		},
		{
			name: "test return true when hashmap is empty after few operations",
			actualResult: func() bool {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				_, err = hm.Remove(2)
				require.NoError(t, err)

				_, err = hm.Put(3, 6)
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
				hm, err := NewHashMap(NewPair(1, 2))
				require.NoError(t, err)

				return hm.IsEmpty()
			},
			expectedResult: false,
		},
		{
			name: "test return false when hashmap is not empty after few operations",
			actualResult: func() bool {
				hm, err := NewHashMap(NewPair(1, 2), NewPair(2, 4))
				require.NoError(t, err)

				_, err = hm.Remove(2)
				require.NoError(t, err)

				_, err = hm.Put(3, 6)
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
