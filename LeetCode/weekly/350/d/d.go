package main

// https://github.com/Ocyss
func paintWalls(cost []int, time []int) (ans int) {
	n := len(cost)
	//f := make([][]int, n+1)
	//for i := range f {
	//	f[i] = make([]int, n+1)
	//	f[i][n] = 1e8
	//}
	//
	//for i := 0; i < n; i++ {
	//	for j := n - 1; j >= 0; j-- {
	//		f[i][j] = min(f[i+1][j], f[i+1][j-time[i]-1]+cost[i])
	//	}
	//}
	//fmt.Println(f)
	//return f[0][n]
	var dfs func(int, int) int
	dfs = func(i, need int) int {
		if i >= n {
			if need > 0 {
				return 1e7
			}
			return 0
		}
		if cache[i][need] == 0 {
			cache[i][need] = min(dfs(i+1, need), dfs(i+1, need-time[i]-1)+cost[i])
		}
		return cache[i][need]
	}
	return dfs(0, n)
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
