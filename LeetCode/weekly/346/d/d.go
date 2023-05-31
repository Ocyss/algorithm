package main

// https://github.com/Ocyss
func modifiedGraphEdges(n int, es [][]int, st int, dt int, tar int) (ans [][]int) {
	set := map[int][]pair{}
	for _, v := range es {
		set[v[0]] = append(set[v[0]], pair{v[1], v[2]})
		set[v[1]] = append(set[v[1]], pair{v[0], v[2]})
	}
	var dfs func(int, int) (int, int)
	vis := make([]bool, n)
	dfs = func(st, dt int) (int, int) {
		if vis[st] {
			return -1, -1
		}
		if st == dt {
			return 0, 0
		}
		for _, v := range set[st] {
			vis[st] = true
			k, s := dfs(v.i, dt)
			vis[st] = false
			if v.v == -1 {
				s++
			} else {
				k += v.v
			}
			if k > tar {
				continue
			}
		}

	}
	return
}

type pair struct {
	i, v int
}
