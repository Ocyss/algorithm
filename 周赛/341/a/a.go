package main

// https://github.com/Ocyss
func rowAndMaximumOnes(a [][]int) (ans []int) {
	r := make([]int, len(a))
	for i, x := range a {
		for _, v := range x {
			if v == 1 {
				r[i]++
			}
		}
	}
	ans = []int{0, 0}
	for i, v := range r {
		if v > ans[1] {
			ans = []int{i, v}
		}
	}
	return
}
