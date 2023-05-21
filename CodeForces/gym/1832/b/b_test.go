package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF883I https://codeforces.com/problemset/problem/883/I
// https://codeforces.com/problemset/status/883/problem/I
// https://github.com/Ocyss
func TestCF1832B(t *testing.T) {
	// just copy from website
	rawText := `
input
6
5 1
2 5 1 10 6
5 2
2 5 1 10 6
3 1
1 2 3
6 1
15 22 12 10 13 11
6 2
15 22 12 10 13 11
5 1
999999996 999999999 999999997 999999998 999999995
output
21
11
3
62
46
3999999986`
	testutil.AssertEqualCase(t, rawText, 0, CF1832B)
}
