package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

var cache = map[string]bool{}

func start(n int, nums []int, ss []string) int {
	ans := 0
	check := func(s1, s2 string) bool {
		var a, b uint8 = 0, 0
		s := s1 + s2
		if v, ok := cache[s]; ok {
			return v
		}
		n := len(s) / 2
		for i := 0; i < n; i++ {
			a += s[i] - '0'
			b += s[i+n] - '0'
		}
		cache[s] = a == b
		if len(s1) == len(s2) {
			cache[s2+s1] = a == b
		}
		return cache[s]
	}
	for i := 0; i < n; i++ {
		x := len(ss[i]) % 2
		j := 0
		for ; j < n; j++ {
			if len(ss[j])%2 == x && check(ss[i], ss[j]) {
				ans++
				// Printf("i:%d,%s   j:%d,%s\n", i, ss[i], j, ss[j])
			}
		}
	}
	return ans
}

// https://github.com/Ocyss
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	nums := make([]int, n)
	ss := make([]string, n)
	for i := range nums {
		// Fscan(in, &nums[i])
		Fscan(in, &ss[i])
	}
	Fprint(out, start(n, nums, ss))

	// 如果是多组数据，请务必加上这段保险 —— 已经无法统计，有多少位竞赛选手在漏读数据上损失大把分数（包括我）
	// 此外，如果不小心写成了 scan(v) 而不是 scan(&v)，由于未传入指针，不会读入任何数据，这样本地运行一次就能立马发现错误
	//leftData, _ := io.ReadAll(in)
	//if s := strings.TrimSpace(string(leftData)); s != "" {
	//	panic("有未读入的数据：\n" + s)
	//}
}

func main() { run(os.Stdin, os.Stdout) }
