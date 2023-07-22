package main

type pair struct {
	l, r, len int
	g         [][2]int
}

// https://github.com/Ocyss
func sumImbalanceNumbers(a []int) (ans int) {
	p := []pair{}
	n := len(a)
	for i := 0; i < n-1; i++ {
		j := i + 1
		if abs(a[i]-a[j]) > 1 {
			if a[i] > a[j] {
				p = append(p, pair{i, i + 1, 1, [][2]int{{j, i}}})
			} else {
				p = append(p, pair{i, i + 1, 1, [][2]int{{i, j}}})
			}
			ans++
		}
	}
	for len(p) > 0 {
		q := p
		p = nil
		for i := range q {
			v := q[i]
			if v.l > 0 {
				lv := v
				x := a[v.l-1]
				gtemp := lv.g
				lv.g = nil
				for _, j := range gtemp {
					if a[j[0]] > x && a[j[1]] < x {
						if min(abs(a[j[0]]-x), abs(a[j[1]]-x)) > 1 {
							if abs(a[j[0]]-x) < abs(a[j[1]]-x) {
								lv.g = append(lv.g, [2]int{j[0], lv.l - 1})
							} else {
								lv.g = append(lv.g, [2]int{lv.l - 1, j[1]})
							}
						}
					}
				}
				ans += len(lv.g)
				p = append(p, lv)
			}
			if v.r < n-1 {
				lv := v
				x := a[v.r+1]
				gtemp := lv.g
				lv.g = nil
				for _, j := range gtemp {
					if a[j[0]] > x && a[j[1]] < x {
						if min(abs(a[j[0]]-x), abs(a[j[1]]-x)) > 1 {
							if abs(a[j[0]]-x) < abs(a[j[1]]-x) {
								lv.g = append(lv.g, [2]int{j[0], lv.l - 1})
							} else {
								lv.g = append(lv.g, [2]int{lv.l - 1, j[1]})
							}
						}
					}
				}
				ans += len(lv.g)
				p = append(p, lv)
			}
		}
	}
	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
