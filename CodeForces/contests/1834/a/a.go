package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func start(a, b int) int {
	// B 必须是偶数
	ans := 0
	for a < b {
		a++
		b--
		ans++
	}
	if b%2 == 1 {
		ans++
	}
	return ans
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) {
		var n, v int
		Fscan(in, &n)
		a, b := 0, 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v == 1 {
				a++
			} else {
				b++
			}
		}
		Fprintln(out, start(a, b))
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
