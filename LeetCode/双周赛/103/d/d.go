package main

import (
	"fmt"
	"sort"
)

type pair struct {
	i, v int
}

// https://github.com/Ocyss
func countOperationsToEmptyArray(a []int) int64 {
	ans := 0
	n := len(a)
	var t []pair
	for i, v := range a {
		t = append(t, pair{i, v})
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].v < t[j].v
	})
	py := 0
	for _, v := range t {
		fmt.Println(v, py)
		ans += v.i + py
		py = (py + v.i) % n
		n--
	}
	return int64(ans)
}
