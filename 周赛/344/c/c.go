package main

// https://github.com/Ocyss
func colorTheArray(n int, qs [][]int) (ans []int) {
	nums := make([]int, n)
	ans = make([]int, len(qs))
	if n == 1 {
		return
	}
	a := make(map[int]int, n)
	for i, v := range qs {
		index, color := v[0], v[1]
		nums[index] = color
		if v, ok := a[index-1]; ok && v != color {
			delete(a, index-1)
		}
		if v, ok := a[index]; ok && v != color {
			delete(a, index)
		}
		if index > 0 && nums[index] == nums[index-1] {
			a[index-1] = color
		}
		if index < n-1 && nums[index] == nums[index+1] {
			a[index] = color
		}
		ans[i] = len(a)
	}
	return
}
