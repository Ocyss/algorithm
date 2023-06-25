package main

import "sort"

// https://github.com/Ocyss
func findValueOfPartition(a []int) (ans int) {
	n := len(a)
	if n == 2 {
		return abs(a[0] - a[1])
	}
	sort.Ints(a)
	mid := len(a) / 2
	return abs(a[mid+1] - a[mid])
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
