package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF1832A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	var v string
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &v)
		set := map[byte]bool{}
		for j := 0; j < len(v)/2; j++ {
			set[v[j]] = true
		}
		if len(set) > 1 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}

func main() { CF1832A(os.Stdin, os.Stdout) }
