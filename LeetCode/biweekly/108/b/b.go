package main

import "sort"

// https://github.com/Ocyss
func relocateMarbles(a []int, moveFrom []int, moveTo []int) (ans []int) {
	set := map[int]struct{}{}
	ans = make([]int, 0, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}
	for i, v := range moveFrom {
		if v == moveTo[i] {
			continue
		}
		set[moveTo[i]] = struct{}{}
		delete(set, v)
	}
	for k, _ := range set {
		ans = append(ans, k)
	}
	sort.Ints(ans)
	return
}
