package main

// https://github.com/Ocyss
func addMinimum(s string) (ans int) {
	var dfs func(int) int
	n := len(s)
	dfs = func(start int) int {
		if start >= n {
			return 0
		}
		if start < n-1 && s[start] < s[start+1] {
			if start < n-2 && s[start+1] < s[start+2] {
				return dfs(start + 3)
			}
			return dfs(start+2) + 1
		}
		return dfs(start+1) + 2
	}
	ans = dfs(0)
	return
}
