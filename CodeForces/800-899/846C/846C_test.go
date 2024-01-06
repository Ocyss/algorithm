package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/main/testutil"
)

// TestCF846C https://codeforces.com/problemset/problem/846/C
// https://codeforces.com/problemset/status/846/problem/C
// https://github.com/Ocyss
func TestCF846C(t *testing.T) {
	// just copy from website
	rawText := `
input
3
-1 2 3
output
0 1 3
input
4
0 0 -1 0
output
0 0 0
input
1
10000
output
1 1 1`
	testutil.AssertEqualCase(t, rawText, 0, CF846C)
}
