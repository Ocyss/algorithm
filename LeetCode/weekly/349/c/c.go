package main

// https://github.com/Ocyss
func minCost(a []int, x int) int64 {
	ans := 0
	n := len(a)
	for _, v := range a {
		ans += v
	}
	b := make([]int, n)
	copy(b, a)
	var dfs func(int)
	dfs = func(i int) {
		g := -x
		for j := 0; j < n; j++ {
			i := (i + j) % n
			v := b[j] - a[i]
			b[j] = a[i]
			if v > 0 {
				g += v
			}
		}
		if g > 0 {
			ans -= g
			dfs(i + 1)
		}
	}
	dfs(1)
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
