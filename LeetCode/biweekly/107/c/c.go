package main

type pair struct {
	qz, hz byte
	size   int
}

// https://github.com/Ocyss
func minimizeConcatenatedLength(a []string) (ans int) {
	n := len(a)
	p := make([]pair, n)
	qz, hz := map[byte][]int{}, map[byte][]int{}
	for i, v := range a {
		q, h := v[0], v[len(v)-1]
		qz[q] = append(qz[q], i)
		hz[h] = append(hz[h], i)
		p = append(p, pair{q, h, len(v)})
	}
	set := make([]bool, n)
top:
	for _, v := range p {
		q, h := qz[v.qz], hz[v.hz]
		ans += v.size
		if len(q) != 0 || len(h) != 0 {
			for _, qq := range q {
				if !set[qq] {
					ans += p[qq].size - 1
					set[qq] = true
					continue top
				}
			}
			for _, hh := range h {
				if !set[hh] {
					ans += p[hh].size - 1
					set[hh] = true
					continue top
				}
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
