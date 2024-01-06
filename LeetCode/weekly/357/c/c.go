package main

var (
	n   int
	set [][2]int
)

// https://github.com/Ocyss
func maximumSafenessFactor(g [][]int) (ans int) {
	n = len(g)
	if g[0][0] == 1 || g[n-1][n-1] == 1 {
		return
	}
	ans = 99999
	set = make([][2]int, 0, 20)
	for i, c := range g {
		for j, v := range c {
			if v == 1 {
				set = append(set, [2]int{i, j})
				ans = min(ans, i+j)
				ans = min(ans, abs(n-1-i)+abs(n-1-j))
			}
		}
	}
	dfs(0, 0, g, &ans)
	return
}

func dfs(i, j int, g [][]int, ans *int) int {
	if i < 0 || j < 0 || i >= n || j >= n || g[i][j] == 1 {
		return 0
	}
	res := 0
	res = max(res, dfs(i+1, j, g, ans))
	res = max(res, dfs(i, j+1, g, ans))
	res = max(res, dfs(i, j-1, g, ans))
	res = max(res, dfs(i-1, j, g, ans))
	r := abs()
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
