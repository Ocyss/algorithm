package main

import "math/bits"

// https://github.com/Ocyss
const mod int = 1e9 + 7

func specialPerm(a []int) (ans int) {
	n := len(a)
	var dfs func(uint, int) int
	cache := make(map[int]map[uint]int)
	dfs = func(set uint, pre int) int {
		if bits.OnesCount(set) == n {
			return 1
		}
		if cache[pre] == nil {
			cache[pre] = make(map[uint]int)
		} else if cache[pre][set] > 0 {
			return cache[pre][set]
		}
		r := 0
		for i := 0; i < n; i++ {
			if (set>>i)&1 == 0 {
				if check(a[i], pre) {
					r += (dfs(set|(1<<i), a[i])) % mod
				}
			}
		}
		cache[pre][set] = r
		return r
	}
	ans = dfs(0, 1)
	ans = (ans%mod + mod) % mod
	return
}

func check(a, b int) bool {
	return a%b == 0 || b%a == 0
}
