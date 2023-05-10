package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func run(n, m int, nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	ans := sort.Search(n, func(mid int) bool {
		mid++
		m := m
		for i, v := range nums {
			v -= i / mid
			if v <= 0 {
				break
			}
			m -= v
			if m <= 0 {
				return true
			}
		}
		return false
	}) + 1
	if ans > n {
		return -1
	}
	return ans
}

func CF1118D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	fmt.Fscan(in, &n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}
	fmt.Fprint(_w, run(n, m, nums))
}

func main() { CF1118D2(os.Stdin, os.Stdout) }
