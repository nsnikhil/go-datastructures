package unionfind_test

import (
	"github.com/nsnikhil/go-datastructures/unionfind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnionFindFind(t *testing.T) {

	findSet := func(t *testing.T, uf unionfind.UnionFind[int], e int) int {
		res, err := uf.Find(e)
		require.NoError(t, err)
		return res
	}

	uf := unionfind.NewUnionFind[int]()

	for i := 0; i < 6; i++ {
		uf.Add(i)
	}

	for i := 0; i < 6; i++ {
		v, err := uf.Find(i)
		require.NoError(t, err)
		assert.Equal(t, i, v)
	}

	require.NoError(t, uf.Union(0, 1))

	assert.Equal(t, 0, findSet(t, uf, 0))
	assert.Equal(t, 0, findSet(t, uf, 1))

	require.NoError(t, uf.Union(2, 3))

	assert.Equal(t, 2, findSet(t, uf, 2))
	assert.Equal(t, 2, findSet(t, uf, 3))

	require.NoError(t, uf.Union(4, 5))

	assert.Equal(t, 4, findSet(t, uf, 4))
	assert.Equal(t, 4, findSet(t, uf, 5))

	require.NoError(t, uf.Union(1, 5))

	assert.Equal(t, 0, findSet(t, uf, 0))
	assert.Equal(t, 0, findSet(t, uf, 1))
	assert.Equal(t, 0, findSet(t, uf, 4))
	assert.Equal(t, 0, findSet(t, uf, 5))

	require.NoError(t, uf.Union(5, 3))

	for i := 0; i < 6; i++ {
		assert.Equal(t, 0, findSet(t, uf, i))
	}

	for i := 0; i < 6; i++ {
		ok, err := uf.AreInSameSet(0, i)
		require.NoError(t, err)
		assert.True(t, ok)
	}
}
