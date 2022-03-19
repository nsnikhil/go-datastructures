package main

type graph[T comparable] interface {
	addNode(n *node[T])

	createDiEdges(curr *node[T], nodes ...*node[T])
	createWeightedDiEdge(curr, next *node[T], weight int64)
	createBiEdges(curr *node[T], next ...*node[T])
	createWeightedBiEdge(curr, nodes *node[T], weight int64)

	deleteNode(n *node[T])
	deleteEdge(start, end *node[T])

	print()
	dfs()
	bfs()

	//hasLoop() bool
	hasCycle() bool
	//areAdjacent(a, b *node[T) bool
	//degreeOfNode(a *node[T) int
	//hasBridge() bool

	reverse()
	clone() graph[T]

	hasRoute(source, target *node[T]) bool

	//isDirected() bool

	//isConnected() bool

	getConnectedComponents() [][]*node[T]

	shortestPath()
}
