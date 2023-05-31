package main

// https://github.com/Ocyss
func findMaxFish(g [][]int) (ans int) {
	for i, c := range g {
		for j, v := range c {
			if v > 0 {
				ans = max(ans, dfs(g, i, j))
			}
		}
	}
	return
}
func dfs(g [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(g) || j >= len(g[0]) || g[i][j] == 0 {
		return 0
	}
	temp := g[i][j]
	g[i][j] = 0
	return temp + dfs(g, i+1, j) + dfs(g, i-1, j) + dfs(g, i, j+1) + dfs(g, i, j-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
