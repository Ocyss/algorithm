package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"time"
)

func start(n, q int, nums, query []int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = nums[i]
	}

	diff := make([]int, n)
	prefixSum := make([]int, n)

	for _, x := range query {
		diff[x-1]++
	}

	prefixSum[0] = diff[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + diff[i]
	}

	for i := 0; i < n; i++ {
		result[i] += 1 << prefixSum[i]
	}

	return result
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) {
		var n, q int
		Fscan(in, &n)
		Fscan(in, &q)
		nums, query := make([]int, n), make([]int, q)
		for i := range nums {
			Fscan(in, &nums[i])
		}
		for i := range query {
			Fscan(in, &query[i])
		}
		arr := start(n, q, nums, query)
		time.Sleep(time.Second)
		for i := 0; i < len(arr); i++ {
			Fprint(out, arr[i])
			if i < len(arr)-1 {
				Fprint(out, " ")
			}
		}
		Fprintln(out)
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
