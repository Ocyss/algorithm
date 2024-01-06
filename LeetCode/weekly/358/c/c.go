package main

import (
	"math"
	"sort"
)

type pair struct {
	i, v int
}

// https://github.com/Ocyss
func minAbsoluteDifference(a []int, x int) (ans int) {
	if x == 0 {
		return 0
	}
	n := len(a)
	if n-x < 1000 {
		return f2(a, x)
	}
	nums := make([]pair, 0, n)
	for i, v := range a {
		nums = append(nums, pair{i, v})
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i].v == nums[j].v {
			return nums[i].i > nums[j].i
		}
		return nums[i].v < nums[j].v
	})
	set := map[int]int{}
	for i, v := range nums {
		set[v.i] = i
	}
	ans = math.MaxInt
	// fmt.Println(nums)
	for i := 0; i < n-x; i++ {
		index := set[i]
		// fmt.Println(i, index)
		f := func(j int) bool {
			if abs(i-nums[j].i) >= x {
				ans = min(ans, abs(a[i]-nums[j].v))
				return true
			} else if abs(a[i]-nums[j].v) > ans {
				return true
			}
			return false
		}
		for j := 1; j < n; j++ {
			a, b := false, false
			if !a && index+j < n && f(index+j) {
				a = true
			}
			if !b && index-j >= 0 && f(index-j) {
				b = true
			}
			if a == b && a == true {
				break
			}
		}
		if ans == 0 {
			return
		}
	}
	return
}

func f2(a []int, x int) (ans int) {
	ans = math.MaxInt
	for i := 0; i < len(a)-x; i++ {
		for j := i + x; j < len(a); j++ {
			ans = min(ans, abs(a[i]-a[j]))
			if ans == 0 {
				return
			}
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
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
