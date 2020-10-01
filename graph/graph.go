package main

type graph interface {
	addNode(n *node)

	createDiEdges(curr *node, nodes ...*node)
	createWeightedDiEdge(curr, next *node, weight int)
	createBiEdges(curr *node, next ...*node)
	createWeightedBiEdge(curr, nodes *node, weight int)

	deleteNode(n *node)
	deleteEdge(start, end *node)

	print()
	dfs()
	bfs()

	//hasLoop() bool
	hasCycle() bool
	//areAdjacent(a, b *node) bool
	//degreeOfNode(a *node) int
	//hasBridge() bool

	reverse()
	clone() graph

	hasRoute(source, target *node) bool

	//isDirected() bool

	//isConnected() bool

	getConnectedComponents() [][]*node

	shortestPath()
}
