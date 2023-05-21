package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func run(n, m int, nums []int) int {
	if m == 1 {
		return 0
	}
	return sort.Search(nums[n-1]-nums[0], func(mx int) bool {
		f := make([]bool, n+1)
		f[0] = true
		f0 := 0
		for i := m - 1; i < n; i++ {
			for nums[i]-nums[f0] > mx {
				f0++
			}
			for ; f0 <= i-m+1; f0++ {
				if f[f0] {
					f[i+1] = true
					break
				}
			}
		}
		return f[n]
	})
}

func CF883I(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}
	sort.Ints(nums)
	fmt.Fprint(out, run(n, m, nums))
}

func main() { CF883I(os.Stdin, os.Stdout) }
