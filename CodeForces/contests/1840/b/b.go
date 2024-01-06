package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

func start(n, k int) int {
	res := 1
	p := int(math.Pow(2, float64(k-1)))
	var dfs func(int, int, bool)
	dfs = func(i int, n int, flag bool) {
		if n < 0 {
			return
		}
		if i > p {
			if flag {
				res++
			}
			return
		}
		res++
		if i <= n {
			dfs(i*2, n-i, true)
			if flag {
				dfs(i*2, n, false)
			}
		}
	}
	Println("p:", p)
	dfs(1, n, true)
	return res
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) {
		var n, k int
		Fscan(in, &n, &k)
		Fprintln(out, start(n, k))
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
