package main

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/nsnikhil/go-datastructures/stack"
)

type listGraph[T comparable] struct {
	nodes map[*node[T]]bool
}

func newListGraph[T comparable]() graph[T] {
	return &listGraph[T]{
		nodes: make(map[*node[T]]bool),
	}
}

func (lg *listGraph[T]) addNode(n *node[T]) {
	if !lg.nodes[n] {
		lg.nodes[n] = true
	}
}

func (lg *listGraph[T]) createDiEdges(curr *node[T], nodes ...*node[T]) {
	for _, node := range nodes {
		lg.createWeightedDiEdge(curr, node, 0)
	}
}

func (lg *listGraph[T]) createWeightedDiEdge(curr, next *node[T], weight int64) {
	curr.addEdges(newWeightedDiEdge[T](next, weight))
	lg.addNode(curr)
	lg.addNode(next)
}

func (lg *listGraph[T]) createBiEdges(curr *node[T], nodes ...*node[T]) {
	for _, node := range nodes {
		lg.createWeightedDiEdge(curr, node, 0)
		lg.createWeightedDiEdge(node, curr, 0)
	}
}

func (lg *listGraph[T]) createWeightedBiEdge(curr, next *node[T], weight int64) {
	lg.createWeightedDiEdge(curr, next, weight)
	lg.createWeightedDiEdge(next, curr, weight)
}

func (lg *listGraph[T]) deleteNode(n *node[T]) {
	fmt.Println("UN IMPLEMENTED")
}

func (lg *listGraph[T]) deleteEdge(start, end *node[T]) {
	fmt.Println("UN IMPLEMENTED")
}

func (lg *listGraph[T]) print() {
	for n := range lg.nodes {
		fmt.Printf("%v: ", n.getData())
		for edge := range n.getEdges() {
			fmt.Printf("(%v, %d)", edge.getNext().getData(), edge.getWeight())
		}
		fmt.Println()
	}
}

func (lg *listGraph[T]) dfs() {
	var visitNode func(curr *node[T], visited map[*node[T]]bool)
	visitNode = func(curr *node[T], visited map[*node[T]]bool) {
		visited[curr] = true

		for edge := range curr.getEdges() {
			next := edge.getNext()
			if !visited[next] {
				visitNode(next, visited)
			}
		}

		fmt.Printf("%v ", curr.getData())
	}

	visited := make(map[*node[T]]bool)
	for curr := range lg.nodes {
		if !visited[curr] {
			visitNode(curr, visited)
		}
	}
}

func (lg *listGraph[T]) bfs() {
	var visitNode func(curr *node[T], visited map[*node[T]]bool, q queue.Queue[*node[T]])
	visitNode = func(curr *node[T], visited map[*node[T]]bool, q queue.Queue[*node[T]]) {
		q.Add(curr)
		visited[curr] = true

		for !q.Empty() {
			sz := q.Size()

			for i := int64(0); i < sz; i++ {
				n, _ := q.Remove()

				fmt.Printf("%d ", n.getData())

				for edge := range n.getEdges() {
					next := edge.getNext()
					if !visited[next] {
						visited[next] = true
						q.Add(next)
					}
				}
			}
		}
	}

	q := queue.NewLinkedQueue[*node[T]]()
	visited := make(map[*node[T]]bool)
	for curr := range lg.nodes {
		if !visited[curr] {
			visitNode(curr, visited, q)
		}
	}
}

func (lg *listGraph[T]) reverse() {
	var reverseUtil func(curr *node[T], visited map[*node[T]]bool)
	reverseUtil = func(curr *node[T], visited map[*node[T]]bool) {
		visited[curr] = true

		edges := curr.getEdges()
		curr.clearEdges()

		for edge := range edges {
			n := edge.getNext()
			if !visited[n] {
				reverseUtil(n, visited)
			}

			edge.changeNext(curr)
			n.addEdges(edge)
		}
	}

	visited := make(map[*node[T]]bool)
	for curr := range lg.nodes {
		if !visited[curr] {
			reverseUtil(curr, visited)
		}
	}
}

func (lg *listGraph[T]) hasCycle() bool {
	var check func(curr *node[T], pd map[*node[T]]bool, dn map[*node[T]]bool) bool
	check = func(curr *node[T], pd map[*node[T]]bool, dn map[*node[T]]bool) bool {
		pd[curr] = true

		if len(curr.getEdges()) == 0 {
			pd[curr] = false
			dn[curr] = true
			return false
		}

		for e := range curr.getEdges() {
			nx := e.getNext()

			if dn[nx] {
				continue
			}

			if pd[nx] {
				return true
			}

			if check(nx, pd, dn) {
				return true
			}
		}

		pd[curr] = false
		dn[curr] = true
		return false
	}

	pd := make(map[*node[T]]bool)
	dn := make(map[*node[T]]bool)

	for n := range lg.nodes {
		if dn[n] {
			continue
		}

		if pd[n] {
			return true
		}

		if check(n, pd, dn) {
			return true
		}
	}

	return false
}

func (lg *listGraph[T]) clone() graph[T] {
	var cl func(curr *node[T], cache map[*node[T]]*node[T]) *node[T]
	cl = func(curr *node[T], cache map[*node[T]]*node[T]) *node[T] {
		if cache[curr] != nil {
			return cache[curr]
		}

		n := newNode(curr.getData())
		cache[curr] = n

		for e := range curr.getEdges() {
			nx := e.getNext()
			var ne *edge[T]

			if cache[nx] != nil {
				ne = newDiEdge[T](cache[nx])
			} else {
				ne = newDiEdge[T](cl(nx, cache))
			}

			ne.weight = e.getWeight()
			n.addEdge(ne)
		}

		return n
	}

	cache := make(map[*node[T]]*node[T])
	nodes := make(map[*node[T]]bool, 0)

	for n := range lg.nodes {
		t := cache[n]
		if cache[n] == nil {
			nodes[cl(n, cache)] = true
		} else {
			nodes[t] = true
		}
	}

	return &listGraph[T]{
		nodes: nodes,
	}
}

func (lg *listGraph[T]) hasRoute(source, target *node[T]) bool {
	var visit func(curr, target *node[T], visited map[*node[T]]bool) bool
	visit = func(curr, target *node[T], visited map[*node[T]]bool) bool {
		visited[curr] = true

		if curr == target {
			return true
		}

		found := false
		for e := range curr.getEdges() {
			nx := e.getNext()

			if nx == target {
				found = true
				break
			}

			if !visited[nx] && visit(nx, target, visited) {
				found = true
				break
			}
		}

		return found
	}

	visited := make(map[*node[T]]bool)
	return visit(source, target, visited)
}

func (lg *listGraph[T]) getConnectedComponents() [][]*node[T] {
	return koasraju(lg)
}

func koasraju[T comparable](lg *listGraph[T]) [][]*node[T] {
	var pushToStack func(node *node[T], visited map[*node[T]]bool, st *stack.Stack[*node[T]])
	pushToStack = func(node *node[T], visited map[*node[T]]bool, st *stack.Stack[*node[T]]) {
		visited[node] = true

		for edge := range node.getEdges() {
			n := edge.getNext()
			if !visited[n] {
				pushToStack(n, visited, st)
			}
		}

		st.Push(node)
	}

	var printComponent func(node *node[T], visited map[*node[T]]bool, temp []*node[T])
	printComponent = func(node *node[T], visited map[*node[T]]bool, temp []*node[T]) {
		visited[node] = true

		for edge := range node.edges {
			n := edge.getNext()
			if !visited[n] {
				printComponent(n, visited, temp)
			}
		}

		temp = append(temp, node)
	}

	st := stack.NewStack[*node[T]]()
	visited := make(map[*node[T]]bool)

	for node := range lg.nodes {
		if !visited[node] {
			pushToStack(node, visited, st)
		}
	}

	lg.reverse()

	visited = make(map[*node[T]]bool)

	res := make([][]*node[T], 0)

	for !st.Empty() {
		n, _ := st.Pop()

		if !visited[n] {
			temp := make([]*node[T], 0)
			printComponent(n, visited, temp)
			res = append(res, temp)
		}
	}

	return res
}

func (lg *listGraph[T]) shortestPath() {

	// unweighted graph
	// dag
	// no negative weights -> dijkstra
	// dijkstra modifications

	// general case -> bellmen ford

	// all pair shortest path -> floyd(DP)
}

func nonWeightedShortestPath[T comparable](source, target *node[T], lg *listGraph[T]) {
	visited := make(map[*node[T]]bool)
	q := queue.NewLinkedQueue[*node[T]]()

	q.Add(source)

	for !q.Empty() {
		sz := q.Size()

		found := false
		for i := int64(0); i < sz; i++ {
			e, _ := q.Remove()

			for edge := range e.getEdges() {
				n := edge.getNext()

				if n == source {
					continue
				}

				if !visited[n] {
					if n.predecessor == nil {
						n.predecessor = e
					}

					if n == target {
						found = true
						break
					}

					visited[n] = true
					q.Add(n)
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

	for curr != nil {
		fmt.Printf("%v ", curr.getData())

		curr = curr.predecessor
	}

}

func dagShortestPath[T comparable](lg *listGraph[T]) {
	var updateCost func(curr *node[T])

	updateCost = func(curr *node[T]) {
		for edge := range curr.edges {
			n := edge.getNext()

			if n.costToReach > curr.costToReach+edge.getWeight() {
				n.costToReach = curr.costToReach + edge.getWeight()
			}
		}
	}

	first := true
	var firstNode *node[T]

	for node := range lg.nodes {
		if first {
			firstNode = node
			node.costToReach = 0
			first = false
		}

		updateCost(node)
	}

	if firstNode == nil {
		return
	}

	for curr := range lg.nodes {
		fmt.Printf("%d : %d\n", curr.getData(), curr.costToReach)
	}

}

type nodeComparator[T comparable] struct {
}

func (nc *nodeComparator[T]) Compare(one *node[T], two *node[T]) int {
	return int(one.costToReach - two.costToReach)
}

func dijkstra[T comparable](start *node[T], lg *listGraph[T]) {
	var relaxCost func(curr *node[T], q *queue.PriorityQueue[*node[T]])
	relaxCost = func(curr *node[T], q *queue.PriorityQueue[*node[T]]) {

		for edge := range curr.edges {
			n := edge.getNext()

			if n.costToReach > curr.costToReach+edge.getWeight() {

				//TODO: NEED TO VERIFY IF THIS CHANGE WORKS
				cp := n.copy()
				cp.costToReach = curr.costToReach + edge.getWeight()
				q.Add(cp)

			}
		}

	}

	//type nodeWrapper struct {
	//	*node
	//	costToReach int
	//	predecessor *node
	//}

	start.costToReach = 0

	q := queue.NewPriorityQueue[*node[T]](false, &nodeComparator[T]{})
	for node := range lg.nodes {
		q.Add(node)
	}

	relaxedNodes := set.NewHashSet[*node[T]]()

	for !q.Empty() {
		n, _ := q.Remove()

		relaxedNodes.Add(n)

		relaxCost(n, q)
	}

}

func bellmenFord[T comparable](start *node[T], lg *listGraph[T]) {
	start.costToReach = 0

	edges := make(map[*edge[T]]*node[T])

	//INEFFICIENT
	for curr := range lg.nodes {
		for edge := range curr.edges {
			edges[edge] = curr
		}
	}

	for range lg.nodes {
		for edge, source := range edges {
			if edge.next.costToReach > source.costToReach+edge.getWeight() {
				edge.next.costToReach = source.costToReach + edge.getWeight()
			}
		}
	}

	for edge, source := range edges {
		if edge.next.costToReach > source.costToReach+edge.getWeight() {
			fmt.Println("negative cycle")
			return
		}
	}

	for node := range lg.nodes {
		fmt.Printf("%d %d\n", node.getData(), node.costToReach)
	}
}
