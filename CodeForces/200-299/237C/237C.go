package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var p [1e6 + 1]bool

func init() {
	for i := range p {
		p[i] = true
	}
	p[0], p[1] = false, false
	for i := 2; i < 1e6+1; i++ {
		if p[i] {
			for j := 2; j*i < 1e6+1; j++ {
				p[i*j] = false
			}
		}
	}
}
func CF237C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var a, b, k int
	_, _ = fmt.Fscan(in, &a, &b, &k)
	r := sort.Search(b-a+2, func(l int) bool {
		total := 0
		for i := a; i < a+l; i++ {
			if p[i] {
				total++
			}
		}
		left := a
		for right := a + l; right <= b; left, right = left+1, right+1 {
			if total < k {
				return false
			}
			if p[left] {
				total--
			}
			if p[right] {
				total++
			}

		}
		return total >= k
	})
	if r == b-a+2 {
		fmt.Fprint(out, -1)
	} else {
		fmt.Fprint(out, r)
	}

}

func main() { CF237C(os.Stdin, os.Stdout) }
