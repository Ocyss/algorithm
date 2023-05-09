package main

import (
	"fmt"
	"sort"
)

type pair struct {
	x, i int
}

// https://github.com/Ocyss
func getSubarrayBeauty(a []int, k int, x int) (ans []int) {
	ans = make([]int, len(a)-k+1)
	s := make([]pair, k)
	for i := 0; i < k; i++ {
		s[i] = pair{a[i], i}
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].x < s[j].x
	})
	pre := 0
	j := 0
	for i, v := range s {
		if v.i == j {
			pre = i
			break
		}
	}

	for i := range ans {
		if s[x-1].x < 0 {
			ans[i] = s[x-1].x
		} else {
			ans[i] = 0
		}
		fmt.Println(1, ans, s)
		s = append(s[:pre], s[pre+1:]...)
		index := Search(s, a[k])
		fmt.Println(2, index, s)
		r := append([]pair{}, s[index:]...)
		s = append(append(s[:index], pair{a[k], index}), r...)

		fmt.Println(3, index, s)
		for i, v := range s {
			if v.i == j {
				pre = i
				break
			}
		}
		j++
		k++
	}
	return
}

func Search(a []pair, n int) int {
	l, r := 0, len(a)
	for l < r {
		mid := (l + r) >> 1
		if a[mid].x < n {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
