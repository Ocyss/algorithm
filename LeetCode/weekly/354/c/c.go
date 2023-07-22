package main

import "container/heap"

type pair struct {
	V, N int
}
type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].N > h[j].N
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
	old := *h
	val := old[len(old)-1]
	*h = old[:len(old)-1]
	return val
}

// https://github.com/Ocyss
func minimumIndex(a []int) (ans int) {
	var lh, rh hp
	set := map[int]int{}
	for i
	return
}
