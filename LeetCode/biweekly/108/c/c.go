package main

import (
	"fmt"
)

var set = map[string]struct{}{}

func init() {
	for i := 1; ; i *= 5 {
		a := fmt.Sprintf("%b", i)
		if len(a) > 15 {
			break
		}
		set[a] = struct{}{}
	}
}

// https://github.com/Ocyss
func minimumBeautifulSubstrings(s string) (ans int) {
	if s[0] == '0' {
		return -1
	}
	n := len(s)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = 50
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == n {
			return 0
		}
		c := &cache[i]
		if *c == 50 && s[i] != '0' {
			for r := n; r > i; r-- {
				if _, ok := set[s[i:r]]; ok {
					*c = min(*c, dfs(r)+1)
				}
			}
		}
		return *c
	}
	ans = dfs(0)
	if ans >= 50 {
		return -1
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
