package main

// https://github.com/Ocyss
func maximumJumps(a []int, tar int) (ans int) {
	n := len(a)
	if n <= 1 {
		return 0
	}
	i := 0
	ans = -1
	for r := i + 1; r < n; r++ {
		if abs(a[i]-a[r]) <= tar {
			v := maximumJumps(a[r:], tar)
			if v == -1 {
				continue
			}
			ans = max(v+1, ans)
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
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
