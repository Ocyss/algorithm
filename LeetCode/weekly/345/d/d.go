package main

// https://github.com/Ocyss
func countCompleteComponents(n int, es [][]int) (ans int) {
	ljt := make(map[int][]int)
	for _, v := range es {
		ljt[v[0]] = append(ljt[v[0]], v[1])
		ljt[v[1]] = append(ljt[v[1]], v[0])
	}
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}

		lj := ljt[i]
		if len(lj) == 0 || len(lj) == 1 && len(ljt[lj[0]]) == 1 {
			vis[i] = true
			ans += 1
		} else {
			ans += dfs(i, vis, ljt)
		}
	}
	return
}

func dfs(i int, vis []bool, ljt map[int][]int) int {
	if vis[i] {
		return 0
	}
	vis[i] = true
	for _, v := range ljt[i] {
		dfs(v, vis, ljt)
	}

}
