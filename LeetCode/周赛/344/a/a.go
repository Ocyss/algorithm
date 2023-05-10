package main

// https://github.com/Ocyss
func distinctDifferenceArray(a []int) (ans []int) {
	n := len(a)
	qz, hz := map[int]int{}, map[int]int{}
	for _, v := range a {
		hz[v]++
	}
	ans = make([]int, n)
	for i, v := range a {
		qz[v]++
		hz[v]--
		if hz[v] == 0 {
			delete(hz, v)
		}
		ans[i] = len(qz) - len(hz)
	}
	return
}
