package main

import (
	"strconv"
)

var set = make([]int, 1001)

func init() {
	for i := 1; i <= 1000; i++ {
		set[i] = tool(i, i*i)
	}
}

// https://github.com/Ocyss
func punishmentNumber(n int) (ans int) {
	for i := 1; i <= n; i++ {
		ans += set[i]
	}
	return
}

func tool(a, b int) int {
	var dfs func(string, int, int) int
	dfs = func(b string, a, ans int) int {
		if a == 0 {
			return ans
		}
		for i := 0; i < len(b); i++ {
			atoi, _ := strconv.Atoi(b[:i])
			if atoi > a {
				break
			}
			if v := dfs(b[i:], a-atoi, ans+atoi); v > 0 {
				return v
			}
		}
		return 0
	}
	if dfs(strconv.Itoa(b), a, 0) == a {
		return b
	}
	return 0
}
