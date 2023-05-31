package main

import "sort"

// https://github.com/Ocyss
func matrixSum(a [][]int) (ans int) {
	mx := 0
	for _, v := range a {
		sort.Slice(v, func(i, j int) bool {
			return v[i] > v[j]
		})
		mx = max(mx, len(v))
	}
	i := 0
	for ; mx > 0; mx-- {
		f := 0
		for _, v := range a {
			if i < len(v) && v[i] > f {
				f = v[i]
			}
		}
		ans += f
		i++
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
