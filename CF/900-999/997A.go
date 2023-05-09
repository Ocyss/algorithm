package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/Ocyss
func CF997A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() { CF997A(os.Stdin, os.Stdout) }
