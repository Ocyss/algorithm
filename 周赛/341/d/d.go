package main

// https://github.com/Ocyss
func minimumTotalPrice(n int, es [][]int, price []int, trips [][]int) (ans int) {
	g := make([][]int, n)
	for _, v := range es {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	price = dz(g, n, price)
	var dfs func(int, int, []int) int
	dfs = func(start, end int, p []int) int {
		if start == end {
			return 0
		}
		var mi int = 1e8
		p[start] = 0
		for _, v := range g[start] {
			if p[v] == -1 {
				mi = min(mi, dfs(v, end, p)+price[v])
			}
		}
		p[start] = mi
		return mi
	}
	for _, t := range trips {
		p := make([]int, n)
		for i := range p {
			p[i] = -1
		}
		ans += dfs(t[0], t[1], p) + price[t[0]]
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func dz(g [][]int, n int, price []int) []int {
	var dfs func(int, bool) int
	b := make([]bool, n)
	dfs = func(start int, z bool) int {
		if start >= n {
			return 0
		}
		ans := 0
		for _, e := range g[start] {
			if b[e] {
				continue
			}
			b[e] = true
			if z {
				ans += dfs(e, false) + price[e]/2
			} else {
				ans += dfs(e, true)
			}
		}
		return ans
	}
	l := dfs(0, true)
	b = make([]bool, n)
	r := dfs(0, false)

	var dfs2 func(int, bool)
	dfs2 = func(start int, z bool) {
		if start >= n {
			return
		}
		for _, e := range g[start] {
			if b[e] {
				continue
			}
			b[e] = true
			if z {
				price[e] /= 2
				dfs2(e, false)
			} else {
				dfs2(e, true)
			}
		}
	}
	b = make([]bool, n)
	if l > r {
		dfs2(0, true)
	} else {
		dfs2(0, false)
	}
	return price
}
