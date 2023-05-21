package main

import "fmt"

type pair struct {
	x, y int
}

// https://github.com/Ocyss
func maxMoves(g [][]int) (ans int) {
	m, n := len(g), len(g[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}
	var dfs func(int, int, int) int

	dfs = func(x, y, pre int) int {
		p := &cache[x][y]
		if *p != 0 {
			return *p
		}
		fmt.Println(x, y, pre, p)
		ans := 0
		for _, v := range []pair{{-1, 1}, {0, 1}, {1, 1}} {

			i, j := x+v.x, y+v.y
			if i < m && j < n && i >= 0 && j >= 0 && g[i][j] > g[x][y] {
				ans = max(dfs(i, j, g[x][y])+1, ans)
			}

		}
		*p = ans
		return ans
	}
	for i, v := range g {
		ans = max(ans, dfs(i, 0, v[0]))
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
