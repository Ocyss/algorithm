package main

import "sort"

// https://github.com/Ocyss
func maxSum(a []int) (ans int) {
	set := map[int][]int{}
	for i := range a {
		set[get(a[i])] = append(set[get(a[i])], a[i])
	}
	ans = -1
	for _, v := range set {
		if len(v) > 1 {
			sort.Ints(v)
			ans = max(ans, v[len(v)-1]+v[len(v)-2])
		}
	}
	return
}

func get(n int) (r int) {
	for n > 0 {
		r = max(r, n%10)
		n /= 10
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
