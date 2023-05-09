package main

import (
	"math"
)

// https://github.com/Ocyss
type Graph struct {
	g [][]Pair
}
type Pair struct {
	form, edge int
}

func Constructor(n int, es [][]int) Graph {
	var r Graph
	r.g = make([][]Pair, n)
	for _, e := range es {
		r.g[e[0]] = append(r.g[e[0]], Pair{e[1], e[2]})
	}
	return r
}

func (g Graph) AddEdge(e []int) {
	g.g[e[0]] = append(g.g[e[0]], Pair{e[1], e[2]})
}

func (g Graph) ShortestPath(node1 int, node2 int) int {
	var dfs func(int, int) int
	hash := make([]bool, len(g.g))
	dfs = func(n1, d int) int {
		if n1 == node2 {
			return d
		}
		if hash[n1] {
			return math.MaxInt
		}
		hash[n1] = true
		mi := math.MaxInt
		for _, v := range g.g[node1] {
			mi = min(mi, dfs(v.form, d+v.edge))
		}
		return mi
	}
	r := dfs(node1, 0)
	if r == math.MaxInt {
		return -1
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
