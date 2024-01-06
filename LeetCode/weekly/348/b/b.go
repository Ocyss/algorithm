package main

// https://github.com/Ocyss
func semiOrderedPermutation(a []int) (ans int) {
	n := len(a)
	x, y := -1, -1
	for i, v := range a {
		if v == n {
			y = i
		} else if v == 1 {
			x = i
		}
	}
	// fmt.Println(x, y)
	if y < x {
		return n - y + x - 2
	}
	return n - y + x - 1
}
