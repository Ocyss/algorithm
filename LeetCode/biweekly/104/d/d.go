package main

import "fmt"

const mod int = 1e9 + 7

// https://github.com/Ocyss
func sumOfPower(a []int) (ans int) {
	n := len(a)
	spp := map[string]struct{}{}
	var dfs func(int, int, int, map[int]struct{})
	dfs = func(mx, mi, l int, set map[int]struct{}) {
		if _, ok := spp[fmt.Sprint(set)]; ok {
			return
		}
		fmt.Println(l, a[mx], a[mi])
		ans = (ans + a[mx]*a[mx]*a[mi]) % mod
		spp[fmt.Sprint(set)] = struct{}{}
		for i := 0; i < n; i++ {
			if _, ok := set[i]; ok {
				continue
			}
			if a[mx] < a[i] {
				mx = i
			}
			if a[mi] > a[i] {
				mi = i
			}
			set[i] = struct{}{}
			dfs(mx, mi, l+1, set)
			delete(set, i)
		}
	}
	for i := range a {
		dfs(i, i, 1, map[int]struct{}{i: {}})
	}
	ans = (ans%mod + mod) % mod
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type heap []int

func (hp heap) Swap(a, b int) {
	hp[a], hp[b] = hp[b], hp[a]
}

func (hp heap) Less(i, j int) bool {
	return hp[i] > hp[j]
}
