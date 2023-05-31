package main

// https://github.com/Ocyss
func differenceOfDistinctValues(g [][]int) (ans [][]int) {
	m, n := len(g), len(g[0])
	ans = make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		setBottom := map[int]int{}
		setTop := map[int]int{}
		for i, j := i+1, 1; i < n && j < m; i, j = i+1, j+1 {
			setBottom[g[i][j]]++
		}
		for i, j := i, 0; j < m && i < n; i, j = i+1, j+1 {
			ans[j][i] = abs(len(setTop) - len(setBottom))
			setTop[g[i][j]]++
			if i+1 < m && j+1 < n {
				setBottom[g[i+1][j+1]]--
				if setBottom[g[i+1][j+1]] == 0 {
					delete(setBottom, g[i+1][j+1])
				}
			}
		}

	}

	for j := 1; j < m; j++ {
		setBottom := map[int]int{}
		setTop := map[int]int{}
		for i, j := i+1, 1; i < n && j < m; i, j = i+1, j+1 {
			setBottom[g[i][j]]++
		}
		for i, j := i, 0; j < m && i < n; i, j = i+1, j+1 {
			ans[j][i] = abs(len(setTop) - len(setBottom))
			setTop[g[i][j]]++
			if i+1 < m && j+1 < n {
				setBottom[g[i+1][j+1]]--
				if setBottom[g[i+1][j+1]] == 0 {
					delete(setBottom, g[i+1][j+1])
				}
			}
		}

	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
