package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF883I https://codeforces.com/problemset/problem/883/I
// https://codeforces.com/problemset/status/883/problem/I
// https://github.com/Ocyss
func TestCF1832C(t *testing.T) {
	// just copy from website
	rawText := `
input
4
5
1 3 3 3 7
2
4 2
4
1 1 1 1
7
5 4 2 1 0 0 4
output
2
2
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1832C)
}
