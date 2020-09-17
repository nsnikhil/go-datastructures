package unionfind_test

import (
	"github.com/nsnikhil/go-datastructures/unionfind"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnionFindFind(t *testing.T) {
	sz := 6
	uf := unionfind.NewUnionFind(sz)

	for i := 0; i < sz; i++ {
		assert.Equal(t, i, uf.Find(i))
	}

	uf.Union(0, 1)

	assert.Equal(t, 0, uf.Find(0))
	assert.Equal(t, 0, uf.Find(1))

	uf.Union(2, 3)

	assert.Equal(t, 2, uf.Find(2))
	assert.Equal(t, 2, uf.Find(3))

	uf.Union(4, 5)

	assert.Equal(t, 4, uf.Find(4))
	assert.Equal(t, 4, uf.Find(5))

	uf.Union(1, 3)

	assert.Equal(t, 0, uf.Find(2))
	assert.Equal(t, 0, uf.Find(3))

	uf.Union(3, 5)

	assert.Equal(t, 0, uf.Find(4))
	assert.Equal(t, 0, uf.Find(5))
}
