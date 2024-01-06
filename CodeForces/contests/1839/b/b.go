package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// 1 13
// 1 10
// 1 6
// 2 2
func start(n int, nums []pair) int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i].a == nums[j].a {
			return nums[i].b > nums[j].b
		}
		return nums[i].a < nums[j].a
	})
	x := 0
	ans := 0
	set := map[int]int{} // 分数的灯有几个
	for i := 0; i < n; i++ {

		v := nums[i]
		ans += v.b
		set[v.a]++
		x++
		if x == v.a {
			for i < n && nums[i].a == v.a {
				i++
			}
			i--
		}
		t := x
		x -= set[t]
		delete(set, t)
	}
	return ans
}

type pair struct {
	a, b int
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) {
		var n int
		Fscan(in, &n)
		nums := make([]pair, n)
		for i := range nums {
			Fscan(in, &nums[i].a, &nums[i].b)
		}
		ans := start(n, nums)
		Fprint(out, ans, "\n")
	}

	cases := 1
	Fscan(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		solve(curCase)
	}

	// 如果是多组数据，请务必加上这段保险 —— 已经无法统计，有多少位竞赛选手在漏读数据上损失大把分数（包括我）
	// 此外，如果不小心写成了 scan(v) 而不是 scan(&v)，由于未传入指针，不会读入任何数据，这样本地运行一次就能立马发现错误
	leftData, _ := io.ReadAll(in)
	if s := strings.TrimSpace(string(leftData)); s != "" {
		panic("有未读入的数据：\n" + s)
	}
}

func main() { run(os.Stdin, os.Stdout) }
