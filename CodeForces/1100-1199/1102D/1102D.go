package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func run(n int, s string) string {
	set := ['3'][]int{}
	p := n / 3
	for i, v := range s {
		set[v] = append(set[v], i)
	}
	k := []byte(s)
	for len(set['2']) > p {
		if len(set['0']) < p {
			k[set['2'][0]] = '0'
			set['0'] = append(set['0'], set['2'][0])
		} else if len(set['1']) < p {
			k[set['2'][0]] = '1'
			set['1'] = append(set['1'], set['2'][0])
		}
		set['2'] = set['2'][1:]
	}
	for len(set['1']) > p {
		if len(set['0']) < p {
			k[set['1'][0]] = '0'
			set['0'] = append(set['0'], set['1'][0])
			set['1'] = set['1'][1:]
		} else if len(set['2']) < p {
			i := len(set['1']) - 1
			k[set['1'][i]] = '2'
			set['2'] = append(set['2'], set['1'][i])
			set['1'] = set['1'][:i]
		}
	}
	for len(set['0']) > p {
		i := len(set['0']) - 1
		if len(set['2']) < p {
			k[set['0'][i]] = '2'
			set['2'] = append(set['2'], set['0'][i])
		} else if len(set['1']) < p {
			k[set['0'][i]] = '1'
			set['1'] = append(set['1'], set['0'][i])
		}
		set['0'] = set['0'][:i]
	}
	return string(k)
}

func CF1102D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var s string
	_, _ = fmt.Fscan(in, &n, &s)
	fmt.Fprint(out, run(n, s))
}

func main() { CF1102D(os.Stdin, os.Stdout) }
