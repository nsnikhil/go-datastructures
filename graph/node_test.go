package graph

import (
	"errors"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestCreateNewNode(t *testing.T) {
	for i := 0; i < math.MaxInt8; i++ {
		assert.Equal(t, &Node[int]{data: i, edges: set.NewHashSet[*edge[int]]()}, NewNode[int](i))
	}
}

func TestNodeAddEdge(t *testing.T) {
	a := NewNode[int](1)
	b := NewNode[int](2)

	e := newDiEdge[int](b)

	a.addEdge(e)

	assert.True(t, a.edges.Contains(e))
}

func TestNodeRemoveEdge(t *testing.T) {
	testCases := map[string]struct {
		actualError   func() error
		expectedError error
	}{
		"should remove edge successfully": {
			actualError: func() error {
				a := NewNode[int](1)
				b := NewNode[int](2)

				e := newDiEdge[int](b)

				a.addEdge(e)

				return a.removeEdge(e)
			},
		},
		"should return error when edge does not exist": {
			actualError: func() error {
				a := NewNode[int](1)
				b := NewNode[int](2)

				e := newDiEdge[int](b)

				return a.removeEdge(e)
			},
			expectedError: errors.New("edge 1 to 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			internal.AssertErrorEquals(t, testCase.expectedError, testCase.actualError())
		})
	}
}

func TestNodeClearEdges(t *testing.T) {
	a := NewNode[int](1)
	b := NewNode[int](2)

	e := newDiEdge[int](b)

	a.addEdge(e)

	assert.True(t, a.edges.Contains(e))

	a.clearEdges()
	assert.False(t, a.edges.Contains(e))
	assert.Equal(t, int64(0), a.edges.Size())
}

func TestNodeFindEdge(t *testing.T) {
	testCases := map[string]struct {
		res           func() (*edge[int], error)
		expectedRes   *edge[int]
		expectedError error
	}{
		"should find edge successfully": {
			res: func() (*edge[int], error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				e := newDiEdge[int](b)

				a.addEdge(e)

				return a.findEdge(b)
			},
			expectedRes: newDiEdge[int](NewNode[int](2)),
		},
		"should return error when edge is not connected to the node": {
			res: func() (*edge[int], error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				return a.findEdge(b)
			},
			expectedError: errors.New("edge 1 to 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.res()
			assert.Equal(t, testCase.expectedRes, res)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestNodeCopy(t *testing.T) {
	testCases := map[string]struct {
		res         func() *Node[int]
		expectedRes func() *Node[int]
	}{
		"should copy node with no edges": {
			res: func() *Node[int] {
				return NewNode[int](1).copy()
			},
			expectedRes: func() *Node[int] { return NewNode[int](1) },
		},
		"should copy node with edges": {
			res: func() *Node[int] {
				a := NewNode[int](1)
				b := NewNode[int](2)
				c := NewNode[int](3)
				d := NewNode[int](4)
				e := NewNode[int](5)

				be := newDiEdge[int](b)
				ce := newDiEdge[int](c)
				de := newDiEdge[int](d)
				ee := newDiEdge[int](e)

				a.addEdge(be)
				b.addEdge(ce)
				c.addEdge(de)
				d.addEdge(ee)

				return a.copy()
			},
			expectedRes: func() *Node[int] {
				a := NewNode[int](1)
				b := NewNode[int](2)
				c := NewNode[int](3)
				d := NewNode[int](4)
				e := NewNode[int](5)

				be := newDiEdge[int](b)
				ce := newDiEdge[int](c)
				de := newDiEdge[int](d)
				ee := newDiEdge[int](e)

				a.addEdge(be)
				b.addEdge(ce)
				c.addEdge(de)
				d.addEdge(ee)

				return a
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.True(t, areNodeEqual(testCase.expectedRes(), testCase.res()))
		})
	}
}

func TestNodeBfsIteratorHasNext(t *testing.T) {
	testCases := map[string]struct {
		res         func() bool
		expectedRes bool
	}{
		"should return true when iterator has element": {
			res: func() bool {
				n := NewNode[int](1)

				it := n.bfsIterator()

				return it.HasNext()
			},
			expectedRes: true,
		},
		"should return false when iterator is exhausted": {
			res: func() bool {
				n := NewNode[int](1)

				it := n.bfsIterator()

				it.HasNext()
				it.Next()

				return it.HasNext()
			},
			expectedRes: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedRes, testCase.res())
		})
	}
}

func TestNodeBfsIteratorNext(t *testing.T) {
	testCases := map[string]struct {
		res           func() ([]int, error)
		expectedRes   []int
		expectedError error
	}{
		"test bfs iterator for node with no edges": {
			res: func() ([]int, error) {
				a := NewNode[int](1)

				v, err := a.bfsIterator().Next()

				return []int{v.data}, err
			},
			expectedRes: []int{1},
		},
		"test bfs iterator for node with multiple edges": {
			res: func() ([]int, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)
				c := NewNode[int](3)
				d := NewNode[int](4)
				e := NewNode[int](5)

				ae := newDiEdge[int](a)
				be := newDiEdge[int](b)
				ce := newDiEdge[int](c)
				de := newDiEdge[int](d)
				ee := newDiEdge[int](e)

				a.addEdge(be)
				b.addEdge(ce)
				b.addEdge(de)
				d.addEdge(ee)
				e.addEdge(ae)

				it := a.bfsIterator()

				var res []int
				for it.HasNext() {
					v, err := it.Next()
					require.NoError(t, err)

					res = append(res, v.data)
				}

				return res, nil
			},
			expectedRes: []int{1, 2, 3, 4, 5},
		},
		"test bfs iterator for node with one edge": {
			res: func() ([]int, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				e := newDiEdge[int](b)

				a.addEdge(e)

				it := a.bfsIterator()

				var res []int
				for it.HasNext() {
					v, err := it.Next()
					require.NoError(t, err)

					res = append(res, v.data)
				}

				return res, nil
			},
			expectedRes: []int{1, 2},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.res()
			internal.AssertSliceEquals[int](t, testCase.expectedRes, res)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestNodeDfsIteratorHasNext(t *testing.T) {
	testCases := map[string]struct {
		res         func() bool
		expectedRes bool
	}{
		"should return true when iterator has element": {
			res: func() bool {
				n := NewNode[int](1)

				it := n.dfsIterator()

				return it.HasNext()
			},
			expectedRes: true,
		},
		"should return false when iterator is exhausted": {
			res: func() bool {
				n := NewNode[int](1)

				it := n.dfsIterator()

				it.HasNext()
				it.Next()

				return it.HasNext()
			},
			expectedRes: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedRes, testCase.res())
		})
	}
}

func TestNodeDfsIteratorNext(t *testing.T) {
	testCases := map[string]struct {
		res           func() ([]int, error)
		expectedRes   []int
		expectedError error
	}{
		"test dfs iterator for node with no edges": {
			res: func() ([]int, error) {
				a := NewNode[int](1)

				v, err := a.dfsIterator().Next()

				return []int{v.data}, err
			},
			expectedRes: []int{1},
		},
		"test dfs iterator for node with multiple edges": {
			res: func() ([]int, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)
				c := NewNode[int](3)
				d := NewNode[int](4)
				e := NewNode[int](5)

				ae := newDiEdge[int](a)
				be := newDiEdge[int](b)
				ce := newDiEdge[int](c)
				de := newDiEdge[int](d)
				ee := newDiEdge[int](e)

				a.addEdge(be)
				b.addEdge(ce)
				b.addEdge(de)
				d.addEdge(ee)
				e.addEdge(ae)

				it := a.dfsIterator()

				var res []int
				for it.HasNext() {
					v, err := it.Next()
					require.NoError(t, err)

					res = append(res, v.data)
				}

				return res, nil
			},
			expectedRes: []int{1, 2, 3, 4, 5},
		},
		"test dfs iterator for node with one edge": {
			res: func() ([]int, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				e := newDiEdge[int](b)

				a.addEdge(e)

				it := a.dfsIterator()

				var res []int
				for it.HasNext() {
					v, err := it.Next()
					require.NoError(t, err)

					res = append(res, v.data)
				}

				return res, nil
			},
			expectedRes: []int{1, 2},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.res()
			internal.AssertSliceEquals[int](t, testCase.expectedRes, res)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}
