package main

import "fmt"

// https://github.com/Ocyss
func maximumOr(a []int, k int) int64 {
	ans := 0
	set := [32]int{}
	for _, v := range a {
		for i := 31; i >= 0; i-- {
			set[i] += v & 1
			v >>= 1
			if v == 0 {
				break
			}
		}
	}
	fmt.Println(set)
	return int64(ans)
}
