package main

import (
	"math"
)

// https://github.com/Ocyss
func maxNonDecreasingLength(a []int, b []int) (ans int) {
	n := len(a)
	cache := make([][]map[int]struct{}, n)
	for i := range cache {
		cache[i] = make([]map[int]struct{}, n)
		for j := range cache[i] {
			cache[i][j] = make(map[int]struct{}, 100)
		}
	}
	var dfs func(int, int, int)
	dfs = func(i, pre, c int) {
		ans = max(ans, c)

		if i < 0 {
			return
		}
		if _, ok := cache[i][c][pre]; ok {
			return
		}
		cache[i][c][pre] = struct{}{}
		if a[i] <= pre {
			dfs(i-1, a[i], c+1)
		}
		if b[i] <= pre {
			dfs(i-1, b[i], c+1)
		}
		dfs(i-1, a[i], 1)
		dfs(i-1, b[i], 1)
	}
	dfs(n-1, math.MaxInt, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
