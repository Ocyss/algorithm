package main

import (
	"sort"
)

// https://github.com/Ocyss
func longestValidSubstring(s string, forbidden []string) (ans int) {
	n := len(s)
	qzh := make([][26]int, n+1)
	for i := range s {
		qzh[i+1] = qzh[i]
		qzh[i+1][s[i]-'a']++
	}
	//sort.Slice(forbidden, func(i int, j int) bool {
	//	return len(forbidden[i]) < len(forbidden[j])
	//})
	set := make([][26]int, len(forbidden))
	for i, v := range forbidden {
		temp := [26]int{}
		for j := range v {
			temp[j-'a']++
		}
		set[i] = temp
	}
	return sort.Search(n, func(mid int) bool {
		return false
	})
}
