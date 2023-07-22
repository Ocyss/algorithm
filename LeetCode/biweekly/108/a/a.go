package main

// https://github.com/Ocyss
func alternatingSubarray(nums []int) int {
	l, r := 0, 1
	flag := 1
	pre := nums[0]
	ans := 0
	for ; r < len(nums); r++ {
		if nums[r] != pre+flag {
			ans = max(ans, r-l)
			if nums[r] == pre+1 {
				r--

			}
			l = r
			flag = 1
			pre = nums[l]
		} else {
			if flag == 1 {
				flag = -1
			} else {
				flag = 1
			}
			pre = nums[r]
		}
	}
	ans = max(ans, r-l)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
