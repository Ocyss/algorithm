package main

// https://github.com/Ocyss
func findPrefixScore(a []int) (ans []int64) {
	ans = make([]int64, len(a))
	mx := a[0]
	for i, v := range a {
		if v > mx {
			mx = v
		}
		ans[i] = int64(v + mx)
		if i > 0 {
			ans[i] += ans[i-1]
		}

	}
	return
}
