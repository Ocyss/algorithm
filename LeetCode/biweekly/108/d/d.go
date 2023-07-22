package main

var pr = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

// https://github.com/Ocyss
func countBlackBlocks(m int, n int, coordinates [][]int) (ans []int64) {
	ans = make([]int64, 5)
	set := map[[2]int]struct{}{}
	for _, v := range coordinates {
		set[[2]int{v[0], v[1]}] = struct{}{}
	}
	for _, v := range coordinates {
		h := [8]int{}
		for i, p := range pr {
			x, y := p[0]+v[0], p[1]+v[1]
			if x >= 0 && x < m && y < n && y >= 0 {
				if _, ok := set[[2]int{x, y}]; ok {
					h[i] = 1
				}
			} else {
				h[i] = -1
			}
		}
		if h[1] != -1 {
			if h[0] != -1 && h[3] != -1 {
				ans[h[1]+h[0]+h[3]]++
			}
			if h[2] != -1 && h[4] != -1 {
				ans[h[1]+h[2]+h[4]]++
			}
		}
		if h[3] != -1 {
			if h[0] != -1 && h[3] != -1 {
				ans[h[1]+h[0]+h[3]]++
			}
			if h[2] != -1 && h[4] != -1 {
				ans[h[1]+h[2]+h[4]]++
			}
		}
		if h[4] != -1 {
			if h[0] != -1 && h[3] != -1 {
				ans[h[1]+h[0]+h[3]]++
			}
			if h[2] != -1 && h[4] != -1 {
				ans[h[1]+h[2]+h[4]]++
			}
		}
		if h[1] != -1 {
			if h[0] != -1 && h[3] != -1 {
				ans[h[1]+h[0]+h[3]]++
			}
			if h[2] != -1 && h[4] != -1 {
				ans[h[1]+h[2]+h[4]]++
			}
		}

	}
	return
}
