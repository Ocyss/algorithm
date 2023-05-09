package main

// https://github.com/Ocyss
func maxDivScore(a []int, divisors []int) (ans int) {
	hash := make(map[int]int, len(a))
	mx := 0
	for _, v := range a {
		hash[v]++
		if v > mx {
			mx = v
		}
	}
	x := 0
	for _, v := range divisors {
		y := 0
		for i := v; i <= mx; i += v {
			y += hash[i]
		}
		if y > x || (y == x && ans > v) || ans == 0 {
			ans = v
			x = y
		}
	}
	return
}
