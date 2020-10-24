package main

import (
	"fmt"
	"github.com/nsnikhil/go-datastructures/liberr"
	"github.com/nsnikhil/go-datastructures/queue"
	"github.com/nsnikhil/go-datastructures/set"
	"github.com/nsnikhil/go-datastructures/stack"
	"github.com/nsnikhil/go-datastructures/utils"
)

type listGraph struct {
	nodes map[*node]bool
}

func newListGraph() graph {
	return &listGraph{
		nodes: make(map[*node]bool),
	}
}

func (lg *listGraph) addNode(n *node) {
	if !lg.nodes[n] {
		lg.nodes[n] = true
	}
}

func (lg *listGraph) createDiEdges(curr *node, nodes ...*node) {
	for _, node := range nodes {
		lg.createWeightedDiEdge(curr, node, 0)
	}
}

func (lg *listGraph) createWeightedDiEdge(curr, next *node, weight int) {
	curr.addEdges(newWeightedDiEdge(next, weight))
	lg.addNode(curr)
	lg.addNode(next)
}

func (lg *listGraph) createBiEdges(curr *node, nodes ...*node) {
	for _, node := range nodes {
		lg.createWeightedDiEdge(curr, node, 0)
		lg.createWeightedDiEdge(node, curr, 0)
	}
}

func (lg *listGraph) createWeightedBiEdge(curr, next *node, weight int) {
	lg.createWeightedDiEdge(curr, next, weight)
	lg.createWeightedDiEdge(next, curr, weight)
}

func (lg *listGraph) deleteNode(n *node) {
	fmt.Println("UN IMPLEMENTED")
}

func (lg *listGraph) deleteEdge(start, end *node) {
	fmt.Println("UN IMPLEMENTED")
}

func (lg *listGraph) print() {
	for n := range lg.nodes {
		fmt.Printf("%d: ", n.getData())
		for edge := range n.getEdges() {
			fmt.Printf("(%d, %d)", edge.getNext().getData(), edge.getWeight())
		}
		fmt.Println()
	}
}

func (lg *listGraph) dfs() {
	var visitNode func(curr *node, visited map[*node]bool)
	visitNode = func(curr *node, visited map[*node]bool) {
		visited[curr] = true

		for edge := range curr.getEdges() {
			next := edge.getNext()
			if !visited[next] {
				visitNode(next, visited)
			}
		}

		fmt.Printf("%d ", curr.getData())
	}

	visited := make(map[*node]bool)
	for curr := range lg.nodes {
		if !visited[curr] {
			visitNode(curr, visited)
		}
	}
}

func (lg *listGraph) bfs() {
	var visitNode func(curr *node, visited map[*node]bool, q queue.Queue)
	visitNode = func(curr *node, visited map[*node]bool, q queue.Queue) {
		q.Add(curr)
		visited[curr] = true

		for !q.Empty() {
			sz := q.Size()

			for i := 0; i < sz; i++ {
				n, _ := q.Remove()

				fmt.Printf("%d ", n.(*node).getData())

				for edge := range n.(*node).getEdges() {
					next := edge.getNext()
					if !visited[next] {
						visited[next] = true
						q.Add(next)
					}
				}
			}
		}
	}

	q, _ := queue.NewLinkedQueue()
	visited := make(map[*node]bool)
	for curr := range lg.nodes {
		if !visited[curr] {
			visitNode(curr, visited, q)
		}
	}
}

func (lg *listGraph) reverse() {
	var reverseUtil func(curr *node, visited map[*node]bool)
	reverseUtil = func(curr *node, visited map[*node]bool) {
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

	visited := make(map[*node]bool)
	for curr := range lg.nodes {
		if !visited[curr] {
			reverseUtil(curr, visited)
		}
	}
}

func (lg *listGraph) hasCycle() bool {
	var check func(curr *node, pd map[*node]bool, dn map[*node]bool) bool
	check = func(curr *node, pd map[*node]bool, dn map[*node]bool) bool {
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

	pd := make(map[*node]bool)
	dn := make(map[*node]bool)

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

func (lg *listGraph) clone() graph {
	var cl func(curr *node, cache map[*node]*node) *node
	cl = func(curr *node, cache map[*node]*node) *node {
		if cache[curr] != nil {
			return cache[curr]
		}

		n := newNode(curr.getData())
		cache[curr] = n

		for e := range curr.getEdges() {
			nx := e.getNext()
			var ne *edge

			if cache[nx] != nil {
				ne = newDiEdge(cache[nx])
			} else {
				ne = newDiEdge(cl(nx, cache))
			}

			ne.weight = e.getWeight()
			n.addEdge(ne)
		}

		return n
	}

	cache := make(map[*node]*node)
	nodes := make(map[*node]bool, 0)

	for n := range lg.nodes {
		t := cache[n]
		if cache[n] == nil {
			nodes[cl(n, cache)] = true
		} else {
			nodes[t] = true
		}
	}

	return &listGraph{
		nodes: nodes,
	}
}

func (lg *listGraph) hasRoute(source, target *node) bool {
	var visit func(curr, target *node, visited map[*node]bool) bool
	visit = func(curr, target *node, visited map[*node]bool) bool {
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

	visited := make(map[*node]bool)
	return visit(source, target, visited)
}

func (lg *listGraph) getConnectedComponents() [][]*node {
	return koasraju(lg)
}

func koasraju(lg *listGraph) [][]*node {
	var pushToStack func(node *node, visited map[*node]bool, st *stack.Stack)
	pushToStack = func(node *node, visited map[*node]bool, st *stack.Stack) {
		visited[node] = true

		for edge := range node.getEdges() {
			n := edge.getNext()
			if !visited[n] {
				pushToStack(n, visited, st)
			}
		}

		st.Push(node)
	}

	var printComponent func(node *node, visited map[*node]bool, temp []*node)
	printComponent = func(node *node, visited map[*node]bool, temp []*node) {
		visited[node] = true

		for edge := range node.edges {
			n := edge.getNext()
			if !visited[n] {
				printComponent(n, visited, temp)
			}
		}

		temp = append(temp, node)
	}

	st, _ := stack.NewStack()
	visited := make(map[*node]bool)

	for node := range lg.nodes {
		if !visited[node] {
			pushToStack(node, visited, st)
		}
	}

	lg.reverse()

	visited = make(map[*node]bool)

	res := make([][]*node, 0)

	for !st.Empty() {
		n, _ := st.Pop()

		if !visited[n.(*node)] {
			temp := make([]*node, 0)
			printComponent(n.(*node), visited, temp)
			res = append(res, temp)
		}
	}

	return res
}

func (lg *listGraph) shortestPath() {

	// unweighted graph
	// dag
	// no negative weights -> dijkstra
	// dijkstra modifications

	// general case -> bellmen ford

	// all pair shortest path -> floyd(DP)
}

func nonWeightedShortestPath(source, target *node, lg *listGraph) {
	visited := make(map[*node]bool)
	q, _ := queue.NewLinkedQueue()

	q.Add(source)

	for !q.Empty() {
		sz := q.Size()

		found := false
		for i := 0; i < sz; i++ {
			e, _ := q.Remove()

			for edge := range e.(*node).getEdges() {
				n := edge.getNext()

				if n == source {
					continue
				}

				if !visited[n] {
					if n.predecessor == nil {
						n.predecessor = e.(*node)
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
		fmt.Printf("%d ", curr.getData())

		curr = curr.predecessor
	}

}

func dagShortestPath(lg *listGraph) {
	var updateCost func(curr *node)

	updateCost = func(curr *node) {
		for edge := range curr.edges {
			n := edge.getNext()

			if n.costToReach > curr.costToReach+edge.getWeight() {
				n.costToReach = curr.costToReach + edge.getWeight()
			}
		}
	}

	first := true
	var firstNode *node

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

type nodeComparator struct {
}

func (nc *nodeComparator) Compare(one interface{}, two interface{}) (int, error) {
	ot := utils.GetTypeName(one)
	tt := utils.GetTypeName(two)

	if ot != tt {
		return -1, liberr.TypeMismatchError(ot, tt)
	}

	return one.(*node).costToReach - two.(*node).costToReach, nil
}

func dijkstra(start *node, lg *listGraph) {
	var relaxCost func(curr *node, q *queue.PriorityQueue)
	relaxCost = func(curr *node, q *queue.PriorityQueue) {

		for edge := range curr.edges {
			n := edge.getNext()

			if n.costToReach > curr.costToReach+edge.getWeight() {

				uf := func(e interface{}) interface{} {
					e.(*node).costToReach = curr.costToReach + edge.getWeight()
					return e
				}

				q.UpdateFunc(n, uf)
			}
		}

	}

	//type nodeWrapper struct {
	//	*node
	//	costToReach int
	//	predecessor *node
	//}

	start.costToReach = 0

	q, _ := queue.NewPriorityQueue(false, &nodeComparator{})
	for node := range lg.nodes {
		q.Add(node)
	}

	relaxedNodes, _ := set.NewHashSet()

	for !q.Empty() {
		n, _ := q.Remove()

		relaxedNodes.Add(n)

		relaxCost(n.(*node), q)
	}

}

func bellmenFord(start *node, lg *listGraph) {
	start.costToReach = 0

	edges := make(map[*edge]*node)

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
