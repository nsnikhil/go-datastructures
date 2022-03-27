package graph

import (
	"fmt"
	"github.com/nsnikhil/erx"
	"github.com/nsnikhil/go-datastructures/functions/iterator"
	"github.com/nsnikhil/go-datastructures/set"
)

type listGraph[T comparable] struct {
	nodes set.Set[*Node[T]]
}

func NewListGraph[T comparable]() Graph[T] {
	return &listGraph[T]{
		nodes: set.NewHashSet[*Node[T]](),
	}
}

func (lg *listGraph[T]) AddNode(n *Node[T]) {
	if !lg.nodes.Contains(n) {
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
	if !lg.nodes.Contains(curr) {
		return nodeNotFoundError(curr.data, "listGraph.createEdge")
	}

	if !lg.nodes.Contains(next) {
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
	if !lg.nodes.Contains(start) {
		return nodeNotFoundError(start.data, "listGraph.DeleteEdge")
	}

	if !lg.nodes.Contains(end) {
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

func (lg *listGraph[T]) DFSIterator() iterator.Iterator[*Node[T]] {
	return newListGraphIterator(false, lg)
}

func (lg *listGraph[T]) BFSIterator() iterator.Iterator[*Node[T]] {
	return newListGraphIterator(true, lg)
}

type listGraphIterator[T comparable] struct {
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

func newListGraphIterator[T comparable](isBfs bool, graph *listGraph[T]) iterator.Iterator[*Node[T]] {
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
	//var check func(curr *Node[T], pd map[*Node[T]]bool, dn map[*Node[T]]bool) bool
	//check = func(curr *Node[T], pd map[*Node[T]]bool, dn map[*Node[T]]bool) bool {
	//	pd[curr] = true
	//
	//	if len(curr.edges) == 0 {
	//		pd[curr] = false
	//		dn[curr] = true
	//		return false
	//	}
	//
	//	for e := range curr.edges {
	//		nx := e.next
	//
	//		if dn[nx] {
	//			continue
	//		}
	//
	//		if pd[nx] {
	//			return true
	//		}
	//
	//		if check(nx, pd, dn) {
	//			return true
	//		}
	//	}
	//
	//	pd[curr] = false
	//	dn[curr] = true
	//	return false
	//}
	//
	//pd := make(map[*Node[T]]bool)
	//dn := make(map[*Node[T]]bool)
	//
	//it := lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	if dn[v] {
	//		continue
	//	}
	//
	//	if pd[v] {
	//		return true
	//	}
	//
	//	if check(v, pd, dn) {
	//		return true
	//	}
	//}

	return false
}

func (lg *listGraph[T]) HasLoop() bool {
	return false
}

func (lg *listGraph[T]) AreAdjacent(a, b *Node[T]) bool {
	return false
}

func (lg *listGraph[T]) DegreeOfNode(a *Node[T]) int {
	return -1
}

func (lg *listGraph[T]) HasBridge() bool {
	return false
}

func (lg *listGraph[T]) Clone() Graph[T] {
	//	var cl func(curr *Node[T], cache map[*Node[T]]*Node[T]) *Node[T]
	//	cl = func(curr *Node[T], cache map[*Node[T]]*Node[T]) *Node[T] {
	//		if cache[curr] != nil {
	//			return cache[curr]
	//		}
	//
	//		n := NewNode(curr.data)
	//		cache[curr] = n
	//
	//		for e := range curr.edges {
	//			nx := e.next
	//			var ne *edge[T]
	//
	//			if cache[nx] != nil {
	//				ne = newDiEdge[T](cache[nx])
	//			} else {
	//				ne = newDiEdge[T](cl(nx, cache))
	//			}
	//
	//			ne.weight = e.weight
	//			n.addEdge(ne)
	//		}
	//
	//		return n
	//	}
	//
	//	cache := make(map[*Node[T]]*Node[T])
	//	nodes := set.NewHashSet[*Node[T]]()
	//
	//	it := lg.nodes.Iterator()
	//	for it.HasNext() {
	//		v, _ := it.Next()
	//		n := v
	//		t := cache[n]
	//		if cache[n] == nil {
	//			nodes.Add(cl(n, cache))
	//		} else {
	//			nodes.Add(t)
	//		}
	//	}
	//
	//	return &listGraph[T]{
	//		nodes: nodes,
	//	}
	return nil
}

func (lg *listGraph[T]) HasRoute(source, target *Node[T]) bool {
	//	var visit func(curr, target *Node[T], visited map[*Node[T]]bool) bool
	//	visit = func(curr, target *Node[T], visited map[*Node[T]]bool) bool {
	//		visited[curr] = true
	//
	//		if curr == target {
	//			return true
	//		}
	//
	//		found := false
	//		for e := range curr.edges {
	//			nx := e.next
	//
	//			if nx == target {
	//				found = true
	//				break
	//			}
	//
	//			if !visited[nx] && visit(nx, target, visited) {
	//				found = true
	//				break
	//			}
	//		}
	//
	//		return found
	//	}
	//
	//	visited := make(map[*Node[T]]bool)
	//	return visit(source, target, visited)
	return false
}

func (lg *listGraph[T]) IsDirected() bool {
	return false
}

func (lg *listGraph[T]) IsConnected() bool {
	return false
}

func (lg *listGraph[T]) GetConnectedComponents() [][]*Node[T] {
	return koasraju(lg)
}

func koasraju[T comparable](lg *listGraph[T]) [][]*Node[T] {
	return nil
	//var pushToStack func(node *Node[T], visited map[*Node[T]]bool, st *stack.Stack[*Node[T]])
	//pushToStack = func(node *Node[T], visited map[*Node[T]]bool, st *stack.Stack[*Node[T]]) {
	//	visited[node] = true
	//
	//	for edge := range node.edges {
	//		n := edgenext
	//		if !visited[n] {
	//			pushToStack(n, visited, st)
	//		}
	//	}
	//
	//	st.Push(node)
	//}
	//
	//var printComponent func(node *Node[T], visited map[*Node[T]]bool, temp []*Node[T])
	//printComponent = func(node *Node[T], visited map[*Node[T]]bool, temp []*Node[T]) {
	//	visited[node] = true
	//
	//	for edge := range node.edges {
	//		n := edgenext
	//		if !visited[n] {
	//			printComponent(n, visited, temp)
	//		}
	//	}
	//
	//	temp = append(temp, node)
	//}
	//
	//st := stack.NewStack[*Node[T]]()
	//visited := make(map[*Node[T]]bool)
	//
	//it := lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	node := v
	//	if !visited[node] {
	//		pushToStack(node, visited, st)
	//	}
	//}
	//
	//lg.Reverse()
	//
	//visited = make(map[*Node[T]]bool)
	//
	//res := make([][]*Node[T], 0)
	//
	//for !st.Empty() {
	//	n, _ := st.Pop()
	//
	//	if !visited[n] {
	//		temp := make([]*Node[T], 0)
	//		printComponent(n, visited, temp)
	//		res = append(res, temp)
	//	}
	//}
	//
	//return res
}

func (lg *listGraph[T]) ShortestPath(source, target *Node[T]) []*Node[T] {
	return nil
	// unweighted Graph
	// dag
	// no negative weights -> dijkstra
	// dijkstra modifications

	// general case -> bellmen ford

	// all pair shortest path -> floyd(DP)
}

func nonWeightedShortestPath[T comparable](source, target *Node[T], lg *listGraph[T]) {
	//visited := make(map[*Node[T]]bool)
	//q := queue.NewLinkedQueue[*Node[T]]()
	//
	//q.Add(source)
	//
	//for !q.Empty() {
	//	sz := q.Size()
	//
	//	found := false
	//	for i := int64(0); i < sz; i++ {
	//		e, _ := q.Remove()
	//
	//		for edge := range e.edges {
	//			n := edgenext
	//
	//			if n == source {
	//				continue
	//			}
	//
	//			if !visited[n] {
	//				if n.predecessor == nil {
	//					n.predecessor = e
	//				}
	//
	//				if n == target {
	//					found = true
	//					break
	//				}
	//
	//				visited[n] = true
	//				q.Add(n)
	//			}
	//		}
	//
	//		if found {
	//			break
	//		}
	//
	//	}
	//
	//	if found {
	//		break
	//	}
	//}
	//
	//curr := target
	//
	//for curr != nil {
	//	fmt.Printf("%v ", currdata)
	//
	//	curr = curr.predecessor
	//}

}

func dagShortestPath[T comparable](lg *listGraph[T]) {
	//var updateCost func(curr *Node[T])
	//
	//updateCost = func(curr *Node[T]) {
	//	for edge := range curr.edges {
	//		n := edgenext
	//
	//		if n.costToReach > curr.costToReach+edge.weight {
	//			n.costToReach = curr.costToReach + edge.weight
	//		}
	//	}
	//}
	//
	//first := true
	//var firstNode *Node[T]
	//
	//it := lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	node := v
	//	if first {
	//		firstNode = node
	//		node.costToReach = 0
	//		first = false
	//	}
	//
	//	updateCost(node)
	//}
	//
	//if firstNode == nil {
	//	return
	//}
	//
	//it = lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	curr := v
	//	fmt.Printf("%v : %d\n", currdata, curr.costToReach)
	//}

}

type nodeComparator[T comparable] struct {
}

func (nc *nodeComparator[T]) Compare(one *Node[T], two *Node[T]) int {
	return 0
	//return int(one.costToReach - two.costToReach)
}

func dijkstra[T comparable](start *Node[T], lg *listGraph[T]) {
	//var relaxCost func(curr *Node[T], q *queue.PriorityQueue[*Node[T]])
	//relaxCost = func(curr *Node[T], q *queue.PriorityQueue[*Node[T]]) {
	//
	//	for edge := range curr.edges {
	//		n := edgenext
	//
	//		if n.costToReach > curr.costToReach+edge.weight {
	//
	//			//TODO: NEED TO VERIFY IF THIS CHANGE WORKS
	//			cp := n.copy()
	//			cp.costToReach = curr.costToReach + edge.weight
	//			q.Add(cp)
	//
	//		}
	//	}
	//
	//}

	//type nodeWrapper struct {
	//	*Node
	//	costToReach int
	//	predecessor *Node
	//}

	//start.costToReach = 0
	//
	//q := queue.NewPriorityQueue[*Node[T]](false, &nodeComparator[T]{})
	//it := lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	node := v
	//	q.Add(node)
	//}
	//
	//relaxedNodes := set.NewHashSet[*Node[T]]()
	//
	//for !q.Empty() {
	//	n, _ := q.Remove()
	//
	//	relaxedNodes.Add(n)
	//
	//	relaxCost(n, q)
	//}

}

func bellmenFord[T comparable](start *Node[T], lg *listGraph[T]) {
	//start.costToReach = 0
	//
	//edges := make(map[*edge[T]]*Node[T])
	//
	////INEFFICIENT
	//it := lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	curr := v
	//	for edge := range curr.edges {
	//		edges[edge] = curr
	//	}
	//}
	//
	//it = lg.nodes.Iterator()
	//for it.HasNext() {
	//	for edge, source := range edges {
	//		if edge.next.costToReach > source.costToReach+edge.weight {
	//			edge.next.costToReach = source.costToReach + edge.weight
	//		}
	//	}
	//}
	//
	//for edge, source := range edges {
	//	if edge.next.costToReach > source.costToReach+edge.weight {
	//		fmt.Println("negative cycle")
	//		return
	//	}
	//}
	//
	//it = lg.nodes.Iterator()
	//for it.HasNext() {
	//	v, _ := it.Next()
	//	node := v
	//	fmt.Printf("%v %d\n", nodedata, node.costToReach)
	//}

}
