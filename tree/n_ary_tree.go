package tree

type node struct {
	data  interface{}
	links []*node
}

type nAryTree struct {
	root *node
}
