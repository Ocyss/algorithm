package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func start(n, k int, nums [][2]int) int {
	var mi1, mx1, mi2, mx2 int = 1e6, 0, 1e6, 0
	for _, v := range nums {
		l, r := v[0], v[1]
		if l < mi1 {
			mi1 = l
		}
		if l > mx1 {
			mx1 = l
		}
		if r < mi2 {
			mi2 = r
		}
		if r > mx2 {
			mx2 = r
		}
	}

	return max(mx1-mi1, mx2-mi2)
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) {
		var n, k int
		Fscan(in, &n, &k)
		nums := make([][2]int, n)
		for i := range nums {
			Fscan(in, &nums[i][0], &nums[i][1])
		}
		Fprintln(out, start(n, k, nums))
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
