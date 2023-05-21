package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func run(n, k, ans int, nums []int) int {
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if k <= 0 || i >= n-1 || j < 0 {
			return 0
		}
		return min(dfs(i+2, j, k-1)+nums[i]+nums[i+1], dfs(i, j-1, k-1)+nums[j])
	}
	return ans - dfs(0, n-1, k)
}
func CF1832C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	var l, k int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &l, &k)
		nums := make([]int, l)
		ans := 0
		for j := 0; j < l; j++ {
			fmt.Fscan(in, &nums[j])
			ans += nums[j]
		}
		sort.Ints(nums)
		fmt.Fprintln(out, run(l, k, ans, nums))
	}

}

func main() { CF1832C(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
