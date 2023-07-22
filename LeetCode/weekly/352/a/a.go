package main

// https://github.com/Ocyss
func longestAlternatingSubarray(nums []int, threshold int) (ans int) {
	l := 0
	for l = range nums {
		if nums[l]%2 == 0 && nums[l] <= threshold {
			break
		}
	}
	n := len(nums)
	for r := l + 1; r < n; r++ {
		if nums[r] > threshold {
			l = r
		} else if nums[r-1]%2 == nums[r]%2 {
			l = r
		}
		ans = max(ans, r-l+1)

	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
