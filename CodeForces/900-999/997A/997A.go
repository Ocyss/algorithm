package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CF997A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	var s string
	_, _ = fmt.Fscan(in, &n, &x, &y, &s)
	fmt.Println(n, x, y, s)
}

func main() { CF997A(os.Stdin, os.Stdout) }
