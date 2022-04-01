package graph

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"sort"
	"strings"
	"testing"
)

func TestCreateNewListGraph(t *testing.T) {
	actual := NewListGraph[int]()

	expected := &listGraph[int]{
		nodes: set.NewHashSet[*Node[int]](),
	}

	assert.Equal(t, expected, actual)
}

func TestListGraphAddNode(t *testing.T) {
	g := NewListGraph[int]()

	for i := 0; i < math.MaxInt8; i++ {
		n := NewNode[int](i)
		g.AddNode(n)
		assert.True(t, g.(*listGraph[int]).nodes.Contains(n))
	}
}

func TestListGraphCreateDiEdgesSuccess(t *testing.T) {
	g := NewListGraph[int]()

	nodes := make([]*Node[int], 10)
	for i := 0; i < 10; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	for i, ni := range nodes {
		for j, nj := range nodes {
			if i != j {
				assert.NoError(t, g.CreateDiEdge(ni, nj))
			}
		}
	}
}

func TestListGraphCreateDiEdgesFailure(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)
	b := NewNode[int](2)

	err := g.CreateDiEdge(a, b)
	internal.AssertErrorEquals(t, errors.New("node 1 not found in the graph"), err)

	g.AddNode(a)

	err = g.CreateDiEdge(a, b)
	internal.AssertErrorEquals(t, errors.New("node 2 not found in the graph"), err)
}

func TestListGraphCreateWeightedDiEdgeSuccess(t *testing.T) {
	g := NewListGraph[int]()

	nodes := make([]*Node[int], 10)
	for i := 0; i < 10; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	for i, ni := range nodes {
		for j, nj := range nodes {
			if i != j {
				assert.NoError(t, g.CreateWeightedDiEdge(ni, nj, int64(i+1)))
			}
		}
	}
}

func TestListGraphCreateWeightedDiEdgeFailure(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)
	b := NewNode[int](2)

	err := g.CreateWeightedDiEdge(a, b, 1)
	internal.AssertErrorEquals(t, errors.New("node 1 not found in the graph"), err)

	g.AddNode(a)

	err = g.CreateWeightedDiEdge(a, b, 1)
	internal.AssertErrorEquals(t, errors.New("node 2 not found in the graph"), err)
}

func TestListGraphCreateBiEdgesSuccess(t *testing.T) {
	g := NewListGraph[int]()

	nodes := make([]*Node[int], 10)
	for i := 0; i < 10; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	for i, ni := range nodes {
		for j, nj := range nodes {
			if i != j {
				assert.NoError(t, g.CreateBiEdge(ni, nj))
			}
		}
	}
}

func TestListGraphCreateBiEdgesFailure(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)
	b := NewNode[int](2)

	err := g.CreateBiEdge(a, b)
	internal.AssertErrorEquals(t, errors.New("node 1 not found in the graph"), err)

	g.AddNode(a)

	err = g.CreateBiEdge(a, b)
	internal.AssertErrorEquals(t, errors.New("node 2 not found in the graph"), err)
}

func TestListGraphCreateWeightedBiEdgeSuccess(t *testing.T) {
	g := NewListGraph[int]()

	nodes := make([]*Node[int], 10)
	for i := 0; i < 10; i++ {
		nodes[i] = NewNode[int](i)
		g.AddNode(nodes[i])
	}

	for i, ni := range nodes {
		for j, nj := range nodes {
			if i != j {
				assert.NoError(t, g.CreateWeightedBiEdge(ni, nj, int64(i+1)))
			}
		}
	}
}

func TestListGraphCreateWeightedBiEdgeFailure(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)
	b := NewNode[int](2)

	err := g.CreateWeightedBiEdge(a, b, 1)
	internal.AssertErrorEquals(t, errors.New("node 1 not found in the graph"), err)

	g.AddNode(a)

	err = g.CreateWeightedBiEdge(a, b, 1)
	internal.AssertErrorEquals(t, errors.New("node 2 not found in the graph"), err)
}

func TestListGraphDeleteNodeSuccessWhenNoEdgesArePresent(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)
	g.AddNode(a)

	b := NewNode[int](2)
	g.AddNode(b)

	c := NewNode[int](3)
	g.AddNode(c)

	d := NewNode[int](4)
	g.AddNode(d)

	e := NewNode[int](5)
	g.AddNode(e)

	require.NoError(t, g.CreateDiEdge(a, b))
	require.NoError(t, g.CreateDiEdge(b, a))
	require.NoError(t, g.CreateDiEdge(c, a))
	require.NoError(t, g.CreateDiEdge(d, a))
	require.NoError(t, g.CreateDiEdge(e, a))

	assert.NoError(t, g.DeleteNode(a))
	assert.False(t, g.(*listGraph[int]).nodes.Contains(a))

	ed, err := b.findEdge(a)
	assert.Nil(t, ed)
	internal.AssertErrorEquals(t, errors.New("edge 2 to 1 not found in the graph"), err)

	ed, err = c.findEdge(a)
	assert.Nil(t, ed)
	internal.AssertErrorEquals(t, errors.New("edge 3 to 1 not found in the graph"), err)

	ed, err = d.findEdge(a)
	assert.Nil(t, ed)
	internal.AssertErrorEquals(t, errors.New("edge 4 to 1 not found in the graph"), err)

	ed, err = e.findEdge(a)
	assert.Nil(t, ed)
	internal.AssertErrorEquals(t, errors.New("edge 5 to 1 not found in the graph"), err)
}

func TestListGraphDeleteNodeSuccessWithMultipleEdges(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)

	g.AddNode(a)
	assert.True(t, g.(*listGraph[int]).nodes.Contains(a))

	err := g.DeleteNode(a)
	assert.NoError(t, err)
	assert.False(t, g.(*listGraph[int]).nodes.Contains(a))
}

func TestListGraphDeleteNodeFailure(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)

	assert.False(t, g.(*listGraph[int]).nodes.Contains(a))

	err := g.DeleteNode(a)
	internal.AssertErrorEquals(t, errors.New("failed to remove node from graph: set is empty"), err)

	b := NewNode[int](1)
	g.AddNode(b)

	err = g.DeleteNode(a)
	//TODO: REFACTOR THIS CHECK
	assert.True(t, strings.Contains(err.Error(), "failed to remove node from graph"))
}

func TestListGraphDeleteEdgeSuccess(t *testing.T) {
	g := NewListGraph[int]()

	a := NewNode[int](1)
	g.AddNode(a)

	b := NewNode[int](2)
	g.AddNode(b)

	require.NoError(t, g.CreateDiEdge(a, b))

	assert.NoError(t, g.DeleteEdge(a, b))

	ed, err := a.findEdge(b)
	assert.Nil(t, ed)
	internal.AssertErrorEquals(t, errors.New("edge 1 to 2 not found in the graph"), err)
}
func TestListGraphDeleteEdgeFailure(t *testing.T) {
	testCases := map[string]struct {
		actualError   func() error
		expectedError error
	}{
		"should return error when first node is not present in the graph": {
			actualError: func() error {
				g := NewListGraph[int]()
				return g.DeleteEdge(NewNode[int](1), NewNode[int](2))
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error when second node is not present in the graph": {
			actualError: func() error {
				g := NewListGraph[int]()

				a := NewNode[int](1)
				g.AddNode(a)

				return g.DeleteEdge(a, NewNode[int](2))
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
		"should return error no edge exists between the nodes": {
			actualError: func() error {
				g := NewListGraph[int]()

				a := NewNode[int](1)
				g.AddNode(a)

				b := NewNode[int](2)
				g.AddNode(b)

				return g.DeleteEdge(a, b)
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

func TestListGraphContains(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() bool
		expectedResult bool
	}{
		"should return true when node is present in the graphs": {
			actualResult: func() bool {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				return g.Contains(a)
			},
			expectedResult: true,
		},
		"should return false when node is not present in the graphs": {
			actualResult: func() bool {
				a := NewNode[int](1)

				g := NewListGraph[int]()

				return g.Contains(a)
			},
			expectedResult: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestListGraphBfsIteratorHasNext(t *testing.T) {
	testCases := map[string]struct {
		actualRes   func() bool
		expectedRes bool
	}{
		"should return false when iterator is empty": {
			actualRes: func() bool {
				return NewListGraph[int]().BFSIterator().HasNext()
			},
			expectedRes: false,
		},
		"should return true when iterator is not empty": {
			actualRes: func() bool {
				g := NewListGraph[int]()
				g.AddNode(NewNode[int](1))

				return g.BFSIterator().HasNext()
			},
			expectedRes: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedRes, testCase.actualRes())
		})
	}
}

func TestListGraphBfsIteratorNext(t *testing.T) {
	g, _ := graphTwentyFour()

	var temp []int

	it := g.BFSIterator()
	for it.HasNext() {
		v, err := it.Next()
		require.NoError(t, err)

		temp = append(temp, v.data)
	}

	sort.Ints(temp)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, temp)
}

func TestListGraphDfsIteratorHasNext(t *testing.T) {
	testCases := map[string]struct {
		actualRes   func() bool
		expectedRes bool
	}{
		"should return false when iterator is empty": {
			actualRes: func() bool {
				return NewListGraph[int]().DFSIterator().HasNext()
			},
			expectedRes: false,
		},
		"should return true when iterator is not empty": {
			actualRes: func() bool {
				g := NewListGraph[int]()
				g.AddNode(NewNode[int](1))

				return g.DFSIterator().HasNext()
			},
			expectedRes: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedRes, testCase.actualRes())
		})
	}
}

func TestListGraphDfsIteratorNext(t *testing.T) {
	g, _ := graphTwentyFour()

	var temp []int

	it := g.DFSIterator()
	for it.HasNext() {
		v, err := it.Next()
		require.NoError(t, err)

		temp = append(temp, v.data)
	}

	sort.Ints(temp)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, temp)
}

func TestListGraphReverse(t *testing.T) {
	testCases := map[string]struct {
		actualRes   func() Graph[int]
		expectedRes func() Graph[int]
	}{
		"should not fail reverse when graph is empty": {
			actualRes: func() Graph[int] {
				g := NewListGraph[int]()
				g.Reverse()

				return g
			},
			expectedRes: func() Graph[int] {
				return NewListGraph[int]()
			},
		},
		"should reverse graph with one node": {
			actualRes: func() Graph[int] {
				g := NewListGraph[int]()
				g.AddNode(NewNode[int](1))
				g.Reverse()

				return g
			},
			expectedRes: func() Graph[int] {
				g := NewListGraph[int]()
				g.AddNode(NewNode[int](1))

				return g
			},
		},
		"should reverse graph with multiple nodes": {
			actualRes: func() Graph[int] {
				g, _ := graphOne()
				g.Reverse()

				return g
			},
			expectedRes: func() Graph[int] {
				g, _ := graphOneReverse()

				return g
			},
		},
		"should reverse undirected graph produce same graph": {
			actualRes: func() Graph[int] {
				g, _ := graphNine()
				g.Reverse()

				return g
			},
			expectedRes: func() Graph[int] {
				g, _ := graphNine()

				return g
			},
		},
		"should reverse disconnected graph": {
			actualRes: func() Graph[int] {
				g, _ := graphTwentyFive()
				g.Reverse()

				return g
			},
			expectedRes: func() Graph[int] {
				g, _ := graphTwentyFiveReverse()

				return g
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			eg := simplifyGraph(testCase.expectedRes().(*listGraph[int]))
			ag := simplifyGraph(testCase.actualRes().(*listGraph[int]))
			assert.True(t, internal.AreMapsSame[int, []int](eg, ag, intSliceComparator{}))
		})
	}
}

func TestListGraphHasCycle(t *testing.T) {
	cyclicGraphs := getGraphs[int](cyclic)

	for _, cyclicGraph := range cyclicGraphs {
		assert.True(t, cyclicGraph.HasCycle())
	}

	aCyclicGraphs := getGraphs[int](ACyclic)

	for _, aCyclicGraph := range aCyclicGraphs {
		assert.False(t, aCyclicGraph.HasCycle())
	}
}

func TestListGraphHasLoop(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() bool
		expectedResult bool
	}{
		"should return false when loop does not exists": {
			actualResult: func() bool {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				return g.HasLoop()
			},
			expectedResult: false,
		},
		"should return true when loop exists": {
			actualResult: func() bool {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				require.NoError(t, g.CreateDiEdge(a, a))

				return g.HasLoop()
			},
			expectedResult: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestListGraphAreAdjacent(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		"should return false when there is no path between the nodes": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				return g.AreAdjacent(a, b)
			},
			expectedResult: false,
			expectedError:  errors.New("edge 1 to 2 not found in the graph"),
		},
		"should return false when there is path but not adjacent": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)
				c := NewNode[int](3)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)
				g.AddNode(c)

				require.NoError(t, g.CreateDiEdge(a, c))
				require.NoError(t, g.CreateDiEdge(c, b))

				return g.AreAdjacent(a, b)
			},
			expectedResult: false,
			expectedError:  errors.New("edge 1 to 2 not found in the graph"),
		},
		"should return false when there is a reverse edge": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				require.NoError(t, g.CreateDiEdge(b, a))

				return g.AreAdjacent(a, b)
			},
			expectedResult: false,
			expectedError:  errors.New("edge 1 to 2 not found in the graph"),
		},
		"should return true when two nodes not adjacent": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				require.NoError(t, g.CreateDiEdge(a, b))

				return g.AreAdjacent(a, b)
			},
			expectedResult: true,
		},
		"should return error when first node does not exist in the graph": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()

				return g.AreAdjacent(a, b)
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error when second node does not exist in the graph": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)

				return g.AreAdjacent(a, b)
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, testCase.expectedResult, res)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphInDegreeOfNode(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() (int64, error)
		expectedResult int64
		expectedError  error
	}{
		"should return 0 when node has no in-degree": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				return g.InDegreeOfNode(a)
			},
			expectedResult: 0,
		},
		"should return 1 for self loop": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				require.NoError(t, g.CreateDiEdge(a, a))

				return g.InDegreeOfNode(a)
			},
			expectedResult: 1,
		},
		"should return 0 for di direction edge": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				require.NoError(t, g.CreateDiEdge(a, b))

				return g.InDegreeOfNode(a)
			},
			expectedResult: 0,
		},
		"should return 1 for bi direction edge": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				require.NoError(t, g.CreateBiEdge(a, b))

				return g.InDegreeOfNode(a)
			},
			expectedResult: 1,
		},
		"should return 4 when 4 nodes point to the curr node": {
			actualResult: func() (int64, error) {
				g, _ := graphTwentyFour()

				n := getNodeWithVal(g, 2)

				return g.InDegreeOfNode(n)
			},
			expectedResult: 4,
		},
		"should return 2 when 2 nodes point to the curr node": {
			actualResult: func() (int64, error) {
				g, _ := graphTwentyFive()

				n := getNodeWithVal(g, 2)

				return g.InDegreeOfNode(n)
			},
			expectedResult: 2,
		},
		"should return error when node is not present in the graph": {
			actualResult: func() (int64, error) {
				g := NewListGraph[int]()

				return g.InDegreeOfNode(NewNode[int](1))
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphOutDegreeOfNode(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() (int64, error)
		expectedResult int64
		expectedError  error
	}{
		"should return 0 when node has no out-degree": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				return g.OutDegreeOfNode(a)
			},
			expectedResult: 0,
		},
		"should return 1 for self loop": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)

				g := NewListGraph[int]()
				g.AddNode(a)

				require.NoError(t, g.CreateDiEdge(a, a))

				return g.OutDegreeOfNode(a)
			},
			expectedResult: 1,
		},
		"should return 1 for bi direction edge": {
			actualResult: func() (int64, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				require.NoError(t, g.CreateBiEdge(b, a))

				return g.OutDegreeOfNode(a)
			},
			expectedResult: 1,
		},
		"should return 4 when 4 edges comes out the curr node": {
			actualResult: func() (int64, error) {
				g, _ := graphTwentyFour()

				n := getNodeWithVal(g, 2)

				return g.InDegreeOfNode(n)
			},
			expectedResult: 4,
		},
		"should return 2 when 2 edges comes out the curr node": {
			actualResult: func() (int64, error) {
				g, _ := graphTwentyFive()

				n := getNodeWithVal(g, 2)

				return g.InDegreeOfNode(n)
			},
			expectedResult: 2,
		},
		"should return error when node is not present in the graph": {
			actualResult: func() (int64, error) {
				g := NewListGraph[int]()

				return g.OutDegreeOfNode(NewNode[int](1))
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, res)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphHasBridge(t *testing.T) {

}

func TestListGraphClone(t *testing.T) {
	for _, graph := range getAllGraphs() {
		assert.Equal(t, graph, graph.Clone())
	}
}

func TestListGraphHasRoute(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() (bool, error)
		expectedResult bool
		expectedError  error
	}{
		"should return false when no route exist": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				return g.HasRoute(a, b)
			},
		},
		"should return true when direct route exists": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)
				g.AddNode(b)

				require.NoError(t, g.CreateDiEdge(a, b))

				return g.HasRoute(a, b)
			},
			expectedResult: true,
		},
		"should return true when indirect route exists": {
			actualResult: func() (bool, error) {
				g, _ := graphFour()

				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)

				require.NoError(t, g.CreateDiEdge(a, b))

				return g.HasRoute(a, b)
			},
			expectedResult: true,
		},
		"should return error when source not does not exist in graph": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()

				return g.HasRoute(a, b)
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error when target not does not exist in graph": {
			actualResult: func() (bool, error) {
				a := NewNode[int](1)
				b := NewNode[int](2)

				g := NewListGraph[int]()
				g.AddNode(a)

				return g.HasRoute(a, b)
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			ok, err := testCase.actualResult()

			assert.Equal(t, testCase.expectedResult, ok)
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphIsDirected(t *testing.T) {

}

func TestListGraphIsConnected(t *testing.T) {

}

func TestListGraphGetConnectedComponents(t *testing.T) {
	testCases := map[string]struct {
		actualResult   func() []list.List[*Node[int]]
		expectedResult func() []list.List[*Node[int]]
	}{
		"get empty list when graph is empty": {
			actualResult: func() []list.List[*Node[int]] {
				return NewListGraph[int]().GetConnectedComponents()
			},
			expectedResult: func() []list.List[*Node[int]] {
				return []list.List[*Node[int]]{}
			},
		},
		"get individual nodes when graph has not edges": {
			actualResult: func() []list.List[*Node[int]] {
				g := NewListGraph[int]()

				for i := 0; i < 10; i++ {
					g.AddNode(NewNode[int](i))
				}

				return g.GetConnectedComponents()
			},
			expectedResult: func() []list.List[*Node[int]] {
				var res []list.List[*Node[int]]
				for i := 9; i >= 0; i-- {
					res = append(res, list.NewArrayList[*Node[int]](NewNode[int](i)))
				}
				return res
			},
		},
		"get connected component scenario one": {
			actualResult: func() []list.List[*Node[int]] {
				g, _ := graphFour()
				return g.GetConnectedComponents()
			},
			expectedResult: func() []list.List[*Node[int]] {

				addData := func(res *[]list.List[*Node[int]], elements ...int) {
					temp := list.NewArrayList[*Node[int]]()
					for _, element := range elements {
						temp.Add(NewNode[int](element))
					}
					*res = append(*res, temp)
				}

				var res []list.List[*Node[int]]
				addData(&res, 0)
				addData(&res, 5)
				addData(&res, 3, 2, 1)
				addData(&res, 4)
				return res
			},
		},
		"get connected component scenario two": {
			actualResult: func() []list.List[*Node[int]] {
				g, _ := graphTwentyTwo()
				return g.GetConnectedComponents()
			},
			expectedResult: func() []list.List[*Node[int]] {

				addData := func(res *[]list.List[*Node[int]], elements ...int) {
					temp := list.NewArrayList[*Node[int]]()
					for _, element := range elements {
						temp.Add(NewNode[int](element))
					}
					*res = append(*res, temp)
				}

				var res []list.List[*Node[int]]
				addData(&res, 0)
				addData(&res, 3)
				addData(&res, 4)
				addData(&res, 5)
				addData(&res, 1)
				addData(&res, 2)
				return res
			},
		},
		"get connected component scenario three": {
			actualResult: func() []list.List[*Node[int]] {
				g, _ := graphTwentyFour()
				return g.GetConnectedComponents()
			},
			expectedResult: func() []list.List[*Node[int]] {

				addData := func(res *[]list.List[*Node[int]], elements ...int) {
					temp := list.NewArrayList[*Node[int]]()
					for _, element := range elements {
						temp.Add(NewNode[int](element))
					}
					*res = append(*res, temp)
				}

				var res []list.List[*Node[int]]
				addData(&res, 12, 14, 13, 11, 10)
				addData(&res, 9, 8)
				addData(&res, 3, 5, 6, 7, 4, 2, 1, 0)
				return res
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.True(t, areComponentsEqual(testCase.expectedResult(), testCase.actualResult()))
		})
	}
}

func TestListGraphUnWeightedGraphShortestPath(t *testing.T) {
	toList := func(data ...int) list.List[*Node[int]] {
		res := list.NewLinkedList[*Node[int]]()
		for _, e := range data {
			res.Add(NewNode[int](e))
		}
		return res
	}

	testCases := map[string]struct {
		actualResult   func() (list.List[*Node[int]], error)
		expectedResult list.List[*Node[int]]
		expectedError  error
	}{
		"should return shortest path for unweighted graph scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphOne()
				a := getNodeWithVal(g, 6)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedResult: toList(6, 5, 4),
		},
		"should return shortest path for unweighted graph scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphOne()
				a := getNodeWithVal(g, 6)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(b, a, UnWeighted)
			},
			expectedResult: toList(4, 6),
		},
		"should return shortest path for unweighted graph scenario three": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphFour()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedResult: toList(0, 1, 3, 2),
		},
		"should return shortest path for unweighted graph scenario four": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFour()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 6)
				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedResult: toList(5, 1, 2, 6),
		},
		"should return error when no path exists between source and target scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFour()
				a := getNodeWithVal(g, 7)
				b := getNodeWithVal(g, 11)
				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedError: errors.New("path 7 to 11 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphFour()
				a := getNodeWithVal(g, 2)
				b := getNodeWithVal(g, 0)
				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedError: errors.New("path 2 to 0 not found in the graph"),
		},
		"should return error source vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error target vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				g.AddNode(a)

				return g.ShortestPath(a, b, UnWeighted)
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.True(t, isListEqual(testCase.expectedResult, res))
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphDirectedAcyclicGraphShortestPath(t *testing.T) {
	toList := func(data ...int) list.List[*Node[int]] {
		res := list.NewLinkedList[*Node[int]]()
		for _, e := range data {
			res.Add(NewNode[int](e))
		}
		return res
	}

	testCases := map[string]struct {
		actualResult   func() (list.List[*Node[int]], error)
		expectedResult list.List[*Node[int]]
		expectedError  error
	}{
		"should return shortest path for dag scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphNineTeen()
				a := getNodeWithVal(g, 4)
				b := getNodeWithVal(g, 5)
				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedResult: toList(4, 6, 7, 5),
		},
		"should return shortest path for dag scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphNineTeen()
				a := getNodeWithVal(g, 7)
				b := getNodeWithVal(g, 5)
				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedResult: toList(7, 5),
		},
		"should return shortest path for dag scenario three": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyOne()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedResult: toList(0, 1, 3, 2),
		},
		"should return shortest path for dag scenario four": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyTwo()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedResult: toList(0, 3, 2),
		},
		"should return error when no path exists between source and target scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyTwo()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedError: errors.New("path 5 to 4 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyOne()
				a := getNodeWithVal(g, 3)
				b := getNodeWithVal(g, 1)
				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedError: errors.New("path 3 to 1 not found in the graph"),
		},
		"should return error source vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error target vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				g.AddNode(a)

				return g.ShortestPath(a, b, Directed, ACyclic)
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.True(t, isListEqual(testCase.expectedResult, res))
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphNonNegativeWeightGraphShortestPath(t *testing.T) {
	toList := func(data ...int) list.List[*Node[int]] {
		res := list.NewLinkedList[*Node[int]]()
		for _, e := range data {
			res.Add(NewNode[int](e))
		}
		return res
	}

	testCases := map[string]struct {
		actualResult   func() (list.List[*Node[int]], error)
		expectedResult list.List[*Node[int]]
		expectedError  error
	}{
		"should return shortest path for non negative graph scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphOne()
				a := getNodeWithVal(g, 6)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(6, 5, 4),
		},
		"should return shortest path for non negative graph scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphOne()
				a := getNodeWithVal(g, 6)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(b, a, NonNegativeWeights)
			},
			expectedResult: toList(4, 6),
		},
		"should return shortest path for non negative graph scenario three": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphFour()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(0, 1, 3, 2),
		},
		"should return shortest path for non negative graph scenario four": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFour()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 6)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(5, 1, 2, 6),
		},
		"should return shortest path for non negative graph scenario five": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphNineTeen()
				a := getNodeWithVal(g, 4)
				b := getNodeWithVal(g, 5)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(4, 6, 7, 5),
		},
		"should return shortest path for non negative graph scenario six": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphNineTeen()
				a := getNodeWithVal(g, 7)
				b := getNodeWithVal(g, 5)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(7, 5),
		},
		"should return shortest path for non negative graph scenario seven": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyOne()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(0, 1, 3, 2),
		},
		"should return shortest path for non negative graph scenario eight": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyTwo()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(0, 3, 2),
		},
		"should return shortest path for non negative graph scenario nine": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphEighteen()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 1)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(5, 3, 2, 1),
		},
		"should return shortest path for non negative graph scenario ten": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFive()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedResult: toList(0, 1, 2, 6, 7, 4),
		},
		"should return error when no path exists between source and target scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyTwo()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedError: errors.New("path 5 to 4 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyOne()
				a := getNodeWithVal(g, 3)
				b := getNodeWithVal(g, 1)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedError: errors.New("path 3 to 1 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario three": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFour()
				a := getNodeWithVal(g, 7)
				b := getNodeWithVal(g, 11)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedError: errors.New("path 7 to 11 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario four": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphFour()
				a := getNodeWithVal(g, 2)
				b := getNodeWithVal(g, 0)
				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedError: errors.New("path 2 to 0 not found in the graph"),
		},
		"should return error source vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error target vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				g.AddNode(a)

				return g.ShortestPath(a, b, NonNegativeWeights)
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.True(t, isListEqual(testCase.expectedResult, res))
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func TestListGraphGraphShortestPath(t *testing.T) {
	toList := func(data ...int) list.List[*Node[int]] {
		res := list.NewLinkedList[*Node[int]]()
		for _, e := range data {
			res.Add(NewNode[int](e))
		}
		return res
	}

	testCases := map[string]struct {
		actualResult   func() (list.List[*Node[int]], error)
		expectedResult list.List[*Node[int]]
		expectedError  error
	}{
		"should return shortest path for graph scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphOne()
				a := getNodeWithVal(g, 6)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(6, 5, 4),
		},
		"should return shortest path for graph scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphOne()
				a := getNodeWithVal(g, 6)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(b, a)
			},
			expectedResult: toList(4, 6),
		},
		"should return shortest path for graph scenario three": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphFour()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(0, 1, 3, 2),
		},
		"should return shortest path for graph scenario four": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFour()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 6)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(5, 1, 2, 6),
		},
		"should return shortest path for graph scenario five": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphNineTeen()
				a := getNodeWithVal(g, 4)
				b := getNodeWithVal(g, 5)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(4, 6, 7, 5),
		},
		"should return shortest path for graph scenario six": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphNineTeen()
				a := getNodeWithVal(g, 7)
				b := getNodeWithVal(g, 5)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(7, 5),
		},
		"should return shortest path for graph scenario seven": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyOne()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(0, 1, 3, 2),
		},
		"should return shortest path for graph scenario eight": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyTwo()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 2)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(0, 3, 2),
		},
		"should return shortest path for graph scenario nine": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphEighteen()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 1)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(5, 3, 2, 1),
		},
		"should return shortest path for graph scenario ten": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFive()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b)
			},
			expectedResult: toList(0, 1, 2, 6, 7, 4),
		},
		"should return error when no path exists between source and target scenario one": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyTwo()
				a := getNodeWithVal(g, 5)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("path 5 to 4 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario two": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyOne()
				a := getNodeWithVal(g, 3)
				b := getNodeWithVal(g, 1)
				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("path 3 to 1 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario three": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyFour()
				a := getNodeWithVal(g, 7)
				b := getNodeWithVal(g, 11)
				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("path 7 to 11 not found in the graph"),
		},
		"should return error when no path exists between source and target scenario four": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphFour()
				a := getNodeWithVal(g, 2)
				b := getNodeWithVal(g, 0)
				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("path 2 to 0 not found in the graph"),
		},
		"should return error when graph has negative weight cycle": {
			actualResult: func() (list.List[*Node[int]], error) {
				g, _ := graphTwentyThree()
				a := getNodeWithVal(g, 0)
				b := getNodeWithVal(g, 4)
				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("graph has negative weight cycle"),
		},
		"should return error source vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("node 1 not found in the graph"),
		},
		"should return error target vertex is not present in the graph": {
			actualResult: func() (list.List[*Node[int]], error) {
				g := NewListGraph[int]()
				a := NewNode[int](1)
				b := NewNode[int](2)

				g.AddNode(a)

				return g.ShortestPath(a, b)
			},
			expectedError: errors.New("node 2 not found in the graph"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.True(t, isListEqual(testCase.expectedResult, res))
			internal.AssertErrorEquals(t, testCase.expectedError, err)
		})
	}
}

func isListEqual(a, b list.List[*Node[int]]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Size() != b.Size() {
		printList(a)
		printList(b)
		return false
	}

	aIt := a.Iterator()
	bIt := b.Iterator()

	for aIt.HasNext() && bIt.HasNext() {
		av, _ := aIt.Next()
		bv, _ := bIt.Next()
		if av.data != bv.data {
			printList(a)
			printList(b)
			return false
		}
	}

	return true
}

func printList(l list.List[*Node[int]]) {
	it := l.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}
