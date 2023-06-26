package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type pair struct {
	i, v int
}

func CF987C(_r io.Reader, _w io.Writer) {
	const inf int = 1e9
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	p := make([]pair, n)
	for i := range p {
		fmt.Fscan(in, &(p[i].i))
	}
	for i := range p {
		fmt.Fscan(in, &(p[i].v))
	}
	ans := inf
	for j := 1; j < n; j++ {
		l := inf
		for i := j - 1; i >= 0; i-- {
			if p[i].i < p[j].i {
				l = min(l, p[i].v)
			}
		}
		r := inf
		for i := j + 1; i < n; i++ {
			if p[i].i > p[j].i {
				r = min(r, p[i].v)
			}
		}
		ans = min(ans, l+r+p[j].v)
	}

	if ans == inf {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, ans)
	}
}

func main() { CF987C(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(x []int) (a int) {
	for _, v := range x {
		a += v
	}
	return
}
