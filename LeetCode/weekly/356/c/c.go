package main

import "fmt"

// https://github.com/Ocyss
func minimumString(a string, b string, c string) (ans string) {
	ab := f(a, b)
	cab, abc := f(c, ab), f(ab, c)
	ans = min(cab, abc)
	ac := f(a, c)
	bac, acb := f(b, ac), f(ac, b)
	ans = min(ans, bac)
	ans = min(ans, acb)
	bc := f(b, c)
	bca := f(bc, a)
	ans = min(ans, bca)
	cb := f(c, b)
	cba := f(cb, a)
	ans = min(ans, cba)
	return
}

func f(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	// mx := 0
	s := 0
	for i := range a {
		if s >= len(b) {
			break
		}
		if b[s] == a[i] {
			s++
		} else {
			// mx = max(mx, s)
			s = 0
		}
	}
	// mx = max(mx, s)
	fmt.Println(a, b, a+b[s:])
	return a + b[s:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b string) string {
	if len(a) < len(b) {
		return a
	} else if len(a) > len(b) {
		return b
	}
	if a < b {
		return a
	}
	return b
}
