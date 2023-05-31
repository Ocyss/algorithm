package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF237C https://codeforces.com/contest/237/problem/C
// https://codeforces.com/problemset/status/237/problem/C
// https://github.com/Ocyss
func TestCF237C(t *testing.T) {
	// just copy from website
	rawText := `
input
2 4 2
output
3
input
6 13 1
output
4
input
1 4 3
output
-1
`
	testutil.AssertEqualCase(t, rawText, 0, CF237C)
}
