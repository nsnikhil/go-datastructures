package main

import "fmt"

func toGraph(vs [][]int32) map[int32][]int32 {
	graph := make(map[int32][]int32)

	for _, v := range vs {
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	return graph
}

func sort(nd []int32, g map[int32][]int32) []int32 {
	var dfs func(c int32, g map[int32][]int32, vs map[int32]bool)

	dfs = func(c int32, g map[int32][]int32, vs map[int32]bool) {
		vs[c] = true

		if g[c] == nil || len(g[c]) == 0 {
			fmt.Printf("%c ", c)
			return
		}

		for _, e := range g[c] {
			if !vs[e] {
				dfs(e, g, vs)
			}
		}

		fmt.Printf("%c ", c)
	}

	vs := make(map[int32]bool)

	for _, n := range nd {
		if !vs[n] {
			dfs(n, g, vs)
		}
	}

	return nil
}

func main() {
	n := []int32{'a', 'b', 'c', 'd', 'e', 'f'}

	vs := [][]int32{{'a', 'd'}, {'f', 'b'}, {'b', 'd'}, {'f', 'a'}, {'d', 'c'}}

	g := toGraph(vs)

	sort(n, g)
}
