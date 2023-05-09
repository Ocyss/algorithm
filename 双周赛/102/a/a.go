package main

import "fmt"

// https://github.com/Ocyss
func findColumnWidth(g [][]int) (ans []int) {
	ans = make([]int, len(g[0]))
	for _, x := range g {
		for j, v := range x {
			ans[j] = max(ans[j], len(fmt.Sprintf("%v", v)))
		}
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
