package graph

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCreateNewDiEdge(t *testing.T) {
	for i := 0; i < math.MaxInt8; i++ {
		n := NewNode[int](i)
		assert.Equal(t, &edge[int]{next: n}, newDiEdge[int](n))
	}
}

func TestCreateNewWeightedDiEdge(t *testing.T) {
	for i := int64(0); i < math.MaxInt8; i++ {
		n := NewNode[int64](i)
		assert.Equal(t, &edge[int64]{next: n, weight: i}, newWeightedDiEdge[int64](n, i))
	}
}

func TestEdgeChangeWeight(t *testing.T) {
	n := NewNode[int](1)

	e := newWeightedDiEdge[int](n, 10)
	e.changeWeight(20)

	assert.Equal(t, newWeightedDiEdge[int](n, 20), e)
}

func TestEdgeChangeNext(t *testing.T) {
	a := NewNode[int](1)
	b := NewNode[int](2)

	e := newDiEdge[int](a)
	e.changeNext(b)

	assert.Equal(t, newDiEdge[int](b), e)
}

func TestEdgeCopy(t *testing.T) {
	a := NewNode[int](1)

	e := newDiEdge[int](a)

	assert.Equal(t, e, e.copy())
}
