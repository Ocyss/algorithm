package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func run(n, k int) int {
	const mod = 1e9 + 7
	var dfs func(int, int, []int) int
	dfs = func(pre, size int, path []int) int {
		if size == n {
			//fmt.Println(path)
			return 1
		}
		res := 0
		for i, j := pre, 1; i <= k; i = pre * j {
			res += dfs(i, size+1, path)
			j++
		}
		return res % mod
	}
	return dfs(1, 0, []int{})
}

func CF414B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &k, &n)
	fmt.Fprint(out, run(n, k))
}

func main() { CF414B(os.Stdin, os.Stdout) }
