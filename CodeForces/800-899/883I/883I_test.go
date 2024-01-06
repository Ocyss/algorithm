package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/main/testutil"
)

// TestCF883I https://codeforces.com/problemset/problem/883/I
// https://codeforces.com/problemset/status/883/problem/I
// https://github.com/Ocyss
func TestCF883I(t *testing.T) {
	// just copy from website
	rawText := `
input
5 2
50 110 130 40 120
output
20
input
4 1
2 3 4 1
output
0
`
	testutil.AssertEqualCase(t, rawText, 0, CF883I)
}
