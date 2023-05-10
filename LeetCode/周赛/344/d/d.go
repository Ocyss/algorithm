package main

import (
	"math"
)

// https://github.com/Ocyss
func minIncrements(n int, cost []int) (ans int) {
	deelp := int(math.Log2(float64(n + 1)))
	cache := make([]int, len(cost))
	var dfs func(int, int) int
	dfs = func(start, d int) int {
		if d == deelp {
			return 0
		}
		l, r := dfs(start*2, d+1), dfs(start*2+1, d+1)
		cache[start-1] = max(l, r) + cost[start-1]
		return cache[start-1]
	}
	mx := dfs(1, 0)
	var bfs func(int, int, int)
	bfs = func(i, d, p int) {
		if d == deelp {
			return
		}
		x := mx - p - cache[i-1]
		ans += x
		cost[i-1] += x
		bfs(i*2, d+1, p+cost[i-1])
		bfs(i*2+1, d+1, p+cost[i-1])
	}
	bfs(1, 0, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
