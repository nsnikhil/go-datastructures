package disjointSets

import (
	"github.com/nsnikhil/go-datastructures/internal"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestCreateNewDisjointSet(t *testing.T) {
	testCases := map[string]struct {
		data     []int
		expected func() DisjointSets[int]
	}{
		"test create disjoint set with empty list": {
			data: nil,
			expected: func() DisjointSets[int] {
				return NewDisjointSets[int]()
			},
		},
		"test create disjoint sets with list of arguments": {
			data: []int{1, 2, 3, 4, 5},
			expected: func() DisjointSets[int] {
				sets := gmap.NewHashMap[int, *node[int]]()

				ds := &defaultDisjointSets[int]{
					sets: sets,
				}

				for i := 1; i <= 5; i++ {
					sets.Put(i, newNode[int](i))
				}

				return ds
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expected(), NewDisjointSets[int](testCase.data...))
		})
	}
}

func TestDisjointSetsMakeSet(t *testing.T) {
	ds := NewDisjointSets[int]()

	for i := 0; i < math.MaxInt8; i++ {
		assert.NotPanics(t, func() { ds.MakeSet(i) })
	}
}

func TestDisjointSetsFindSet(t *testing.T) {
	data := internal.SliceGenerator{Size: math.MaxInt8}.Generate()

	ds := NewDisjointSets[int64](data...)

	for i := int64(0); i < math.MaxInt8; i++ {
		v, err := ds.FindSet(i)
		require.NoError(t, err)
		assert.Equal(t, i, v)
	}
}

func TestDisjointSetsUnion(t *testing.T) {
	findSet := func(t *testing.T, uf DisjointSets[int], e int) int {
		res, err := uf.FindSet(e)
		require.NoError(t, err)
		return res
	}

	ds := NewDisjointSets[int](0, 1, 2, 3, 4, 5)

	for i := 0; i < 6; i += 2 {
		require.NoError(t, ds.Union(i, i+1))
		assert.Equal(t, i, findSet(t, ds, i))
		assert.Equal(t, i, findSet(t, ds, i+1))
	}

	require.NoError(t, ds.Union(1, 5))

	assert.Equal(t, 0, findSet(t, ds, 0))
	assert.Equal(t, 0, findSet(t, ds, 1))
	assert.Equal(t, 0, findSet(t, ds, 4))
	assert.Equal(t, 0, findSet(t, ds, 5))

	require.NoError(t, ds.Union(5, 3))

	for i := 0; i < 6; i++ {
		assert.Equal(t, 0, findSet(t, ds, i))
	}
}

func TestDisjointAreInSameSet(t *testing.T) {
	ds := NewDisjointSets[int](0, 1, 2, 3, 4, 5)

	for i := 0; i < 6; i++ {
		v, err := ds.FindSet(i)
		require.NoError(t, err)

		ok, err := ds.AreInSameSet(i, v)
		require.NoError(t, err)
		assert.True(t, ok)
	}

	for i := 1; i < 6; i++ {
		ok, err := ds.AreInSameSet(0, i)
		require.NoError(t, err)
		assert.False(t, ok)
	}

	for i := 0; i < 5; i++ {
		require.NoError(t, ds.Union(i, i+1))
	}

	for i := 0; i < 6; i++ {
		ok, err := ds.AreInSameSet(0, i)
		require.NoError(t, err)
		assert.True(t, ok)
	}
}

func TestDisjointUniqueSetsCount(t *testing.T) {
	ds := NewDisjointSets[int](0, 1, 2, 3, 4, 5)

	assert.Equal(t, int64(6), ds.SetsCount())

	for i := 0; i < 6; i += 2 {
		require.NoError(t, ds.Union(i, i+1))
	}

	assert.Equal(t, int64(3), ds.SetsCount())

	require.NoError(t, ds.Union(1, 5))

	assert.Equal(t, int64(2), ds.SetsCount())

	require.NoError(t, ds.Union(5, 3))

	assert.Equal(t, int64(1), ds.SetsCount())
}

func TestDisjointGetAllSets(t *testing.T) {
	ds := NewDisjointSets[int](0, 1, 2, 3, 4, 5)

	assert.Equal(t, 6, len(ds.GetAllSets()))

	for i := 0; i < 6; i += 2 {
		require.NoError(t, ds.Union(i, i+1))
	}

	assert.Equal(t, 3, len(ds.GetAllSets()))

	require.NoError(t, ds.Union(1, 5))

	assert.Equal(t, 2, len(ds.GetAllSets()))

	require.NoError(t, ds.Union(5, 3))

	assert.Equal(t, 1, len(ds.GetAllSets()))
}

func TestDisjointClear(t *testing.T) {
	testCases := map[string]struct {
		data     []int
		expected func() DisjointSets[int]
	}{
		"test clear empty set": {
			data: nil,
		},
		"test clear set with data": {
			data: []int{1, 2, 3, 4, 5},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ds := NewDisjointSets[int](testCase.data...)
			ds.Clear()
			assert.Equal(t, int64(0), ds.SetsCount())
		})
	}
}
