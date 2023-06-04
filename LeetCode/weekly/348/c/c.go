package main

// https://github.com/Ocyss
func matrixSumQueries(n int, qs [][]int) int64 {
	cache := map[[2]int]bool{}
	x, y := n, n
	k := n * n
	ans := 0
	for i := len(qs) - 1; i >= 0 && k >= 0; i-- {
		ty, in, va := qs[i][0], qs[i][1], qs[i][2]
		if cache[[2]int{ty, in}] {
			continue
		}
		cache[[2]int{ty, in}] = true
		if ty == 1 {
			ans += x * va
			y--
		} else {
			ans += y * va
			x--
		}
	}

	return int64(ans)
}
