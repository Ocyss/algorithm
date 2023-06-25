package main

import (
	"sort"
)

// https://github.com/Ocyss
func smallestString(s string) (ans string) {
	a := []int{}
	k := -1
	for i, v := range s {
		if v == 'a' {
			a = append(a, i)
		} else if k == -1 {
			k = i
		}
	}
	n := len(s)
	if len(a) == n {
		return rev(s, n-1, n-1)
	}
	if len(a) == 0 {
		return rev(s, k, n-1)
	}
	if a[0] > 0 {
		return rev(s, 0, a[0]-1)
	}
	index := sort.SearchInts(a, k)
	if index == len(a) {
		return rev(s, k, n-1)
	}
	return rev(s, k, a[index]-1)
}

func rev(s string, l, r int) string {
	ss := []byte(s)
	for i := l; i <= r; i++ {
		if ss[i] == 'a' {
			ss[i] = 'z'
		} else {
			ss[i]--
		}
	}
	return string(ss)
}
