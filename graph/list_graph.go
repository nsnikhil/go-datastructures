package graph

import (
	"errors"
	"fmt"
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/internal"
	"github.com/nsnikhil/go-datastructures/list"
	gmap "github.com/nsnikhil/go-datastructures/map"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/nsnikhil/go-datastructures/stack"
	"math"
)

type listGraph[T any] struct {
	nodes set.Set[*Node[T]]
}

func NewListGraph[T any]() Graph[T] {
	return &listGraph[T]{
		nodes: set.NewHashSet[*Node[T]](),
	}
}

func (lg *listGraph[T]) AddNode(n *Node[T]) {
	if !lg.Contains(n) {
		lg.nodes.Add(n)
	}
}

func (lg *listGraph[T]) CreateDiEdge(curr *Node[T], next *Node[T]) error {
	return lg.createEdge(curr, next, 0)
}

func (lg *listGraph[T]) CreateWeightedDiEdge(curr, next *Node[T], weight int64) error {
	return lg.createEdge(curr, next, weight)
}

func (lg *listGraph[T]) CreateBiEdge(curr *Node[T], next *Node[T]) error {
	if err := lg.createEdge(curr, next, 0); err != nil {
		return err
	}

	return lg.createEdge(next, curr, 0)
}

func (lg *listGraph[T]) CreateWeightedBiEdge(curr, next *Node[T], weight int64) error {
	if err := lg.createEdge(curr, next, weight); err != nil {
		return err
	}

	return lg.createEdge(next, curr, weight)
}

func (lg *listGraph[T]) createEdge(curr, next *Node[T], weight int64) error {
	if !lg.Contains(curr) {
		return nodeNotFoundError(curr.data, "listGraph.createEdge")
	}

	if !lg.Contains(next) {
		return nodeNotFoundError(next.data, "listGraph.createEdge")
	}

	curr.addEdge(newWeightedDiEdge[T](next, weight))
	return nil
}

//TODO: EXPENSIVE IMPLEMENTATION
func (lg *listGraph[T]) DeleteNode(n *Node[T]) error {
	it := lg.nodes.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		if v == n {
			continue
		}

		e, err := v.findEdge(n)
		if err != nil {
			continue
		}

		err = v.removeEdge(e)
		if err != nil {
			return erx.WithArgs(erx.Kind("listGraph.DeleteNode"), err)
		}
	}

	if err := lg.nodes.Remove(n); err != nil {
		return erx.WithArgs(erx.Kind("listGraph.DeleteNode"), fmt.Errorf("failed to remove node from graph: %w", err))
	}

	return nil
}

func (lg *listGraph[T]) DeleteEdge(start, end *Node[T]) error {
	if !lg.Contains(start) {
		return nodeNotFoundError(start.data, "listGraph.DeleteEdge")
	}

	if !lg.Contains(end) {
		return nodeNotFoundError(end.data, "listGraph.DeleteEdge")
	}

	e, err := start.findEdge(end)
	if err != nil {
		return erx.WithArgs(erx.Kind("listGraph.DeleteEdge"), err)
	}

	if err := start.removeEdge(e); err != nil {
		return erx.WithArgs(erx.Kind("listGraph.DeleteEdge"), err)
	}

	return nil
}

func (lg *listGraph[T]) Contains(n *Node[T]) bool {
	return lg.nodes.Contains(n)
}

func (lg *listGraph[T]) DFSIterator() iterator.Iterator[*Node[T]] {
	return newListGraphIterator(false, lg)
}

func (lg *listGraph[T]) BFSIterator() iterator.Iterator[*Node[T]] {
	return newListGraphIterator(true, lg)
}

type listGraphIterator[T any] struct {
	isBfs             bool
	vs                set.Set[*Node[T]]
	nodesIterator     iterator.Iterator[*Node[T]]
	traversalIterator iterator.Iterator[*Node[T]]
}

func (lgi *listGraphIterator[T]) HasNext() bool {
	if lgi.traversalIterator == nil || !lgi.traversalIterator.HasNext() {

		for lgi.nodesIterator.HasNext() {
			n, _ := lgi.nodesIterator.Next()
			if lgi.vs.Contains(n) {
				continue
			}

			if lgi.isBfs {
				lgi.traversalIterator = newNodeBfsIteratorWithVisited[T](n, lgi.vs)
			} else {
				lgi.traversalIterator = newNodeDfsIteratorWithVisited[T](n, lgi.vs)
			}

			break
		}

		return lgi.traversalIterator != nil && lgi.traversalIterator.HasNext()

	}

	return true
}

func (lgi *listGraphIterator[T]) Next() (*Node[T], error) {
	v, err := lgi.traversalIterator.Next()
	if err != nil {
		return nil, emptyIteratorError("listGraphBfsIterator.Next")
	}

	return v, nil
}

func newListGraphIterator[T any](isBfs bool, graph *listGraph[T]) iterator.Iterator[*Node[T]] {
	return &listGraphIterator[T]{
		isBfs:         isBfs,
		vs:            set.NewHashSet[*Node[T]](),
		nodesIterator: graph.nodes.Iterator(),
	}
}

func (lg *listGraph[T]) Reverse() {
	var reverseUtil func(curr *Node[T], vs set.Set[*Node[T]])
	reverseUtil = func(curr *Node[T], vs set.Set[*Node[T]]) {
		vs.Add(curr)

		edges := curr.edges.Copy()
		curr.clearEdges()

		it := edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			if !vs.Contains(e.next) {
				reverseUtil(e.next, vs)
			}

			e.next.addEdge(e)
			e.changeNext(curr)
		}
	}

	vs := set.NewHashSet[*Node[T]]()

	it := lg.nodes.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		if !vs.Contains(v) {
			reverseUtil(v, vs)
		}
	}
}

func (lg *listGraph[T]) HasCycle() bool {

	var check func(curr *Node[T], pd set.Set[*Node[T]], dn set.Set[*Node[T]]) bool

	check = func(curr *Node[T], pd set.Set[*Node[T]], dn set.Set[*Node[T]]) bool {
		pd.Add(curr)

		it := curr.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			nx := e.next

			if dn.Contains(nx) {
				continue
			}

			if pd.Contains(nx) {
				return true
			}

			if check(nx, pd, dn) {
				return true
			}
		}

		pd.Remove(curr)
		dn.Add(curr)
		return false
	}

	pd := set.NewHashSet[*Node[T]]()
	dn := set.NewHashSet[*Node[T]]()

	it := lg.nodes.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		if dn.Contains(v) {
			continue
		}

		if pd.Contains(v) {
			return true
		}

		if check(v, pd, dn) {
			return true
		}
	}

	return false
}

func (lg *listGraph[T]) HasLoop() bool {

	ni := lg.nodes.Iterator()
	for ni.HasNext() {
		n, _ := ni.Next()

		ei := n.edges.Iterator()
		for ei.HasNext() {
			e, _ := ei.Next()
			if n == e.next {
				return true
			}
		}
	}

	return false
}

//TODO: WHAT IF THEIR IS A EDGE FROM B TO A?
func (lg *listGraph[T]) AreAdjacent(a, b *Node[T]) (bool, error) {
	if !lg.Contains(a) {
		return false, nodeNotFoundError(a.data, "listGraph.AreAdjacent")
	}

	if !lg.Contains(b) {
		return false, nodeNotFoundError(b.data, "listGraph.AreAdjacent")
	}

	if _, err := a.findEdge(b); err == nil {
		return true, nil
	}

	return false, edgeNotFoundError(a.data, b.data, "listGraph.AreAdjacent")
}

func (lg *listGraph[T]) InDegreeOfNode(a *Node[T]) (int64, error) {
	if !lg.Contains(a) {
		return internal.Zero, nodeNotFoundError(a.data, "listGraph.InDegreeOfNode")
	}

	var res int64

	ni := lg.nodes.Iterator()
	for ni.HasNext() {
		n, _ := ni.Next()

		ei := n.edges.Iterator()
		for ei.HasNext() {
			e, _ := ei.Next()
			if e.next == a {
				res++
			}
		}
	}

	return res, nil
}

func (lg *listGraph[T]) OutDegreeOfNode(a *Node[T]) (int64, error) {
	if !lg.Contains(a) {
		return internal.Zero, nodeNotFoundError(a.data, "listGraph.OutDegreeOfNode")
	}

	return a.edges.Size(), nil
}

//func (lg *listGraph[T]) HasBridge() bool {
//	return false
//}

func (lg *listGraph[T]) Clone() Graph[T] {

	var cl func(curr *Node[T], cache gmap.Map[*Node[T], *Node[T]]) *Node[T]
	cl = func(curr *Node[T], cache gmap.Map[*Node[T], *Node[T]]) *Node[T] {
		if v, err := cache.Get(curr); v != nil && err == nil {
			return v
		}

		n := NewNode(curr.data)
		cache.Put(curr, n)

		it := curr.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()

			nx := e.next
			var ne *edge[T]

			if v, err := cache.Get(nx); v != nil && err == nil {
				ne = newDiEdge[T](v)
			} else {
				ne = newDiEdge[T](cl(nx, cache))
			}

			ne.weight = e.weight
			n.addEdge(ne)
		}

		return n
	}

	cache := gmap.NewHashMap[*Node[T], *Node[T]]()
	nodes := set.NewHashSet[*Node[T]]()

	it := lg.nodes.Iterator()
	for it.HasNext() {
		n, _ := it.Next()

		t, err := cache.Get(n)
		if err != nil {
			nodes.Add(cl(n, cache))
		} else {
			nodes.Add(t)
		}
	}

	return &listGraph[T]{
		nodes: nodes,
	}
}

func (lg *listGraph[T]) HasRoute(source, target *Node[T]) (bool, error) {
	if !lg.Contains(source) {
		return false, nodeNotFoundError(source.data, "listGraph.HasRoute")
	}

	if !lg.Contains(target) {
		return false, nodeNotFoundError(target.data, "listGraph.HasRoute")
	}

	var visit func(curr, target *Node[T], visited set.Set[*Node[T]]) bool

	visit = func(curr, target *Node[T], visited set.Set[*Node[T]]) bool {
		visited.Add(curr)

		if curr == target {
			return true
		}

		found := false

		it := curr.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			nx := e.next

			if nx == target {
				found = true
				break
			}

			if !visited.Contains(nx) && visit(nx, target, visited) {
				found = true
				break
			}
		}

		return found
	}

	visited := set.NewHashSet[*Node[T]]()

	return visit(source, target, visited), nil
}

//func (lg *listGraph[T]) IsDirected() bool {
//	return false
//}
//
//func (lg *listGraph[T]) IsConnected() bool {
//	return false
//}

func (lg *listGraph[T]) GetConnectedComponents() []list.List[*Node[T]] {
	return koasraju(lg)
}

func koasraju[T any](lg *listGraph[T]) []list.List[*Node[T]] {

	var pushToStack func(node *Node[T], visited set.Set[*Node[T]], st *stack.Stack[*Node[T]])

	pushToStack = func(node *Node[T], visited set.Set[*Node[T]], st *stack.Stack[*Node[T]]) {
		visited.Add(node)

		it := node.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			n := e.next
			if !visited.Contains(n) {
				pushToStack(n, visited, st)
			}
		}

		st.Push(node)
	}

	var printComponent func(node *Node[T], visited set.Set[*Node[T]], temp list.List[*Node[T]])

	printComponent = func(node *Node[T], visited set.Set[*Node[T]], temp list.List[*Node[T]]) {
		visited.Add(node)

		it := node.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			n := e.next
			if !visited.Contains(n) {
				printComponent(n, visited, temp)
			}
		}

		temp.Add(node)
	}

	st := stack.NewStack[*Node[T]]()
	visited := set.NewHashSet[*Node[T]]()

	it := lg.nodes.Iterator()
	for it.HasNext() {
		v, _ := it.Next()
		node := v
		if !visited.Contains(node) {
			pushToStack(node, visited, st)
		}
	}

	//TODO: MUTATING THE GRAPH
	lg.Reverse()

	visited.Clear()

	res := make([]list.List[*Node[T]], 0)

	for !st.Empty() {
		n, _ := st.Pop()

		if !visited.Contains(n) {
			temp := list.NewArrayList[*Node[T]]()
			printComponent(n, visited, temp)
			res = append(res, temp)
		}
	}

	return res
}

func (lg *listGraph[T]) ShortestPath(source, target *Node[T], properties ...Property) (list.List[*Node[T]], error) {
	if !lg.Contains(source) {
		return nil, nodeNotFoundError(source.data, "listGraph.ShortestPath")
	}

	if !lg.Contains(target) {
		return nil, nodeNotFoundError(target.data, "listGraph.ShortestPath")
	}

	toPropertySet := func(properties ...Property) set.Set[Property] {
		res := set.NewHashSet[Property]()
		for _, property := range properties {
			res.Add(property)
		}
		return res
	}

	ps := toPropertySet(properties...)

	//TODO: REFACTOR ALL THE ALGORITHM CONTAINS DUPLICATE CODE TO BACKTRACK NODES

	//UnWeighted Graph -> BFS
	if ps.Contains(UnWeighted) {
		return nonWeightedShortestPath(source, target)
	}

	//DAG -> TOPOLOGICAL SORT
	if ps.Contains(Directed) && ps.Contains(ACyclic) {
		return dagShortestPath(source, target, lg)
	}

	//NON NEGATIVE WEIGHTS -> DIJKSTRA
	if ps.Contains(NonNegativeWeights) {
		return dijkstra(source, target, lg)
	}

	//GENERAL CASE -> BELLMEN FORD
	return bellmenFord(source, target, lg)
}

func nonWeightedShortestPath[T any](source, target *Node[T]) (list.List[*Node[T]], error) {
	vs := set.NewHashSet[*Node[T]]()
	q := queue.NewLinkedQueue[*Node[T]]()

	pm := gmap.NewHashMap[*Node[T], *Node[T]]()

	q.Add(source)
	vs.Add(source)
	pm.Put(source, nil)

	for !q.Empty() {
		sz := q.Size()

		found := false
		for i := int64(0); i < sz; i++ {
			n, _ := q.Remove()

			it := n.edges.Iterator()
			for it.HasNext() {
				e, _ := it.Next()
				nx := e.next

				if !vs.Contains(nx) {
					if !pm.ContainsKey(nx) {
						pm.Put(nx, n)
					}

					if nx == target {
						found = true
						break
					}

					vs.Add(nx)
					q.Add(nx)
				}
			}

			if found {
				break
			}

		}

		if found {
			break
		}
	}

	curr := target
	if !pm.ContainsKey(target) {
		//TODO: REFACTOR OPERATION
		return nil, pathNotFoundError(source.data, target.data, "nonWeightedShortestPath")
	}

	res := list.NewLinkedList[*Node[T]]()

	for pm.ContainsKey(curr) {
		res.AddFirst(curr)

		if curr == source {
			break
		}

		par, err := pm.Get(curr)
		if err != nil {
			//TODO: REFACTOR OPERATION
			return nil, pathNotFoundError(source.data, target.data, "nonWeightedShortestPath")
		}

		curr = par
	}

	return res, nil
}

func dagShortestPath[T any](source, target *Node[T], lg *listGraph[T]) (list.List[*Node[T]], error) {
	cm := gmap.NewHashMap[*Node[T], int64]()
	pm := gmap.NewHashMap[*Node[T], *Node[T]]()

	sortedNodes := lg.topologicalSort()

	it := lg.nodes.Iterator()
	for it.HasNext() {
		n, _ := it.Next()
		pm.Put(n, nil)

		if n == source {
			cm.Put(n, internal.Zero)
		} else {
			cm.Put(n, math.MaxInt64)
		}
	}

	for !sortedNodes.Empty() {
		n, _ := sortedNodes.Pop()

		currCost, _ := cm.Get(n)
		if currCost == math.MaxInt64 {
			continue
		}

		it := n.edges.Iterator()

		for it.HasNext() {
			e, _ := it.Next()
			nx := e.next

			costToReachNext, _ := cm.Get(nx)

			if costToReachNext > currCost+e.weight {
				cm.Put(nx, currCost+e.weight)
				pm.Put(nx, n)
			}
		}
	}

	curr := target
	if !pm.ContainsKey(target) || first(cm.Get(target)) == math.MaxInt64 {
		//TODO: REFACTOR OPERATION
		return nil, pathNotFoundError(source.data, target.data, "dagShortestPath")
	}

	res := list.NewLinkedList[*Node[T]]()

	for pm.ContainsKey(curr) {
		res.AddFirst(curr)

		if curr == source {
			break
		}

		par, err := pm.Get(curr)
		if err != nil {
			//TODO: REFACTOR OPERATION
			return nil, pathNotFoundError(source.data, target.data, "dagShortestPath")
		}

		curr = par
	}

	return res, nil
}

func (lg *listGraph[T]) topologicalSort() *stack.Stack[*Node[T]] {
	var topologicalSortUtil func(n *Node[T], vs set.Set[*Node[T]], st *stack.Stack[*Node[T]])

	topologicalSortUtil = func(n *Node[T], vs set.Set[*Node[T]], st *stack.Stack[*Node[T]]) {
		if n == nil || vs.Contains(n) {
			return
		}

		it := n.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			nx := e.next

			if !vs.Contains(nx) {
				topologicalSortUtil(nx, vs, st)
			}
		}

		st.Push(n)
		vs.Add(n)
	}

	st := stack.NewStack[*Node[T]]()
	vs := set.NewHashSet[*Node[T]]()

	it := lg.nodes.Iterator()
	for it.HasNext() {
		n, _ := it.Next()
		if !vs.Contains(n) {
			topologicalSortUtil(n, vs, st)
		}
	}

	return st
}

type nodeWrapper[T any] struct {
	curr        *Node[T]
	predecessor *Node[T]
	costToReach int64
}

type nodeComparator[T any] struct{}

func (nc *nodeComparator[T]) Compare(one *nodeWrapper[T], two *nodeWrapper[T]) int {
	return int(one.costToReach - two.costToReach)
}

func dijkstra[T any](source, target *Node[T], lg *listGraph[T]) (list.List[*Node[T]], error) {

	var relaxCost func(
		*nodeWrapper[T],
		queue.Queue[*nodeWrapper[T]],
		gmap.Map[*Node[T], int64],
		gmap.Map[*Node[T], *Node[T]],
		set.Set[*Node[T]],
	)

	relaxCost = func(
		currWrapper *nodeWrapper[T],
		q queue.Queue[*nodeWrapper[T]],
		cm gmap.Map[*Node[T], int64],
		pm gmap.Map[*Node[T], *Node[T]],
		relaxedNodes set.Set[*Node[T]],
	) {

		currCost, _ := cm.Get(currWrapper.curr)
		if currCost == math.MaxInt64 {
			return
		}

		it := currWrapper.curr.edges.Iterator()
		for it.HasNext() {
			e, _ := it.Next()
			nx := e.next

			if relaxedNodes.Contains(nx) {
				continue
			}

			costToReach, _ := cm.Get(nx)

			if costToReach > currCost+e.weight {
				newCostToReach := currCost + e.weight
				cm.Put(nx, newCostToReach)
				pm.Put(nx, currWrapper.curr)
				q.Add(&nodeWrapper[T]{curr: nx, costToReach: newCostToReach})
			}
		}
	}

	cm := gmap.NewHashMap[*Node[T], int64]()
	pm := gmap.NewHashMap[*Node[T], *Node[T]]()

	q := queue.NewPriorityQueue[*nodeWrapper[T]](false, &nodeComparator[T]{})

	it := lg.nodes.Iterator()
	for it.HasNext() {
		n, _ := it.Next()
		pm.Put(n, nil)

		if n == source {
			cm.Put(n, 0)
			q.Add(&nodeWrapper[T]{curr: n, costToReach: 0})
		} else {
			cm.Put(n, math.MaxInt64)
			q.Add(&nodeWrapper[T]{curr: n, costToReach: math.MaxInt64})
		}
	}

	relaxedNodes := set.NewHashSet[*Node[T]]()

	for !q.Empty() {
		n, _ := q.Remove()

		relaxCost(n, q, cm, pm, relaxedNodes)

		relaxedNodes.Add(n.curr)
	}

	curr := target
	if !pm.ContainsKey(target) || first(cm.Get(target)) == math.MaxInt64 {
		//TODO: REFACTOR OPERATION
		return nil, pathNotFoundError(source.data, target.data, "dijkstra")
	}

	res := list.NewLinkedList[*Node[T]]()

	for pm.ContainsKey(curr) {
		res.AddFirst(curr)

		if curr == source {
			break
		}

		par, err := pm.Get(curr)
		if err != nil {
			//TODO: REFACTOR OPERATION
			return nil, pathNotFoundError(source.data, target.data, "dijkstra")
		}

		curr = par
	}

	return res, nil
}

func bellmenFord[T any](source, target *Node[T], lg *listGraph[T]) (list.List[*Node[T]], error) {
	cm := gmap.NewHashMap[*Node[T], int64]()
	pm := gmap.NewHashMap[*Node[T], *Node[T]]()

	edges := gmap.NewHashMap[*edge[T], *Node[T]]()

	//INEFFICIENT
	nIt := lg.nodes.Iterator()
	for nIt.HasNext() {
		n, _ := nIt.Next()
		pm.Put(n, nil)

		if n == source {
			cm.Put(n, 0)
		} else {
			cm.Put(n, math.MaxInt64)
		}

		eIt := n.edges.Iterator()
		for eIt.HasNext() {
			e, _ := eIt.Next()
			edges.Put(e, n)
		}
	}

	nIt = lg.nodes.Iterator()
	for nIt.HasNext() {
		_, _ = nIt.Next()

		eIt := edges.Iterator()
		for eIt.HasNext() {
			p, _ := eIt.Next()
			edge := p.First()
			source := p.Second()

			if first(cm.Get(source)) == math.MaxInt64 {
				continue
			}

			if first(cm.Get(edge.next)) > first(cm.Get(source))+edge.weight {
				cm.Put(edge.next, first(cm.Get(source))+edge.weight)
				pm.Put(edge.next, source)
			}

		}
	}

	eIt := edges.Iterator()
	for eIt.HasNext() {
		p, _ := eIt.Next()
		edge := p.First()
		source := p.Second()

		if first(cm.Get(source)) == math.MaxInt64 {
			continue
		}

		if first(cm.Get(edge.next)) > first(cm.Get(source))+edge.weight {
			return nil, erx.WithArgs(erx.Operation("bellmenFord"), errors.New("graph has negative weight cycle"))
		}

	}

	curr := target
	if !pm.ContainsKey(target) || first(cm.Get(target)) == math.MaxInt64 {
		//TODO: REFACTOR OPERATION
		return nil, pathNotFoundError(source.data, target.data, "bellmenFord")
	}

	res := list.NewLinkedList[*Node[T]]()

	for pm.ContainsKey(curr) {

		res.AddFirst(curr)

		if curr == source {
			break
		}

		par, err := pm.Get(curr)
		if err != nil {
			//TODO: REFACTOR OPERATION
			return nil, pathNotFoundError(source.data, target.data, "bellmenFord")
		}

		curr = par
	}

	return res, nil

}

func first[T any, E any](first T, second E) T {
	return first
}
