package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func start(n int, nums []int) (int, [][2]int) {
	sort.Ints(nums)
	anss := [][2]int{}
	ans := 0
	for i := 0; i < n; i++ {
		if i+1 < n {
			ans += nums[i+1] - nums[i]
			ans += nums[i+1+n] - nums[i+n]
		}
		anss = append(anss, [2]int{nums[i], nums[i+n]})
	}
	return ans, anss
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) {
		var n int
		Fscan(in, &n)
		nums := make([]int, 2*n)
		for i := range nums {
			Fscan(in, &nums[i])
		}
		ans, anss := start(n, nums)
		Fprintln(out, ans)
		for _, v := range anss {
			Fprintln(out, v[0], v[1])
		}
	}

	cases := 1
	Fscan(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		solve(curCase)
	}

	// 如果是多组数据，请务必加上这段保险 —— 已经无法统计，有多少位竞赛选手在漏读数据上损失大把分数（包括我）
	// 此外，如果不小心写成了 scan(v) 而不是 scan(&v)，由于未传入指针，不会读入任何数据，这样本地运行一次就能立马发现错误
	//leftData, _ := io.ReadAll(in)
	//if s := strings.TrimSpace(string(leftData)); s != "" {
	//	panic("有未读入的数据：\n" + s)
	//}
}

func main() { run(os.Stdin, os.Stdout) }
