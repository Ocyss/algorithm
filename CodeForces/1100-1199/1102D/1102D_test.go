package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF1102D https://codeforces.com/problemset/problem/1102/D
// https://codeforces.com/problemset/status/1102/problem/D
// https://github.com/Ocyss
func TestCF1102D(t *testing.T) {
	// just copy from website
	rawText := `
input
3
121
output
021
input
6
000000
output
001122
input
6
211200
output
211200
input
6
120110
output
120120
`
	testutil.AssertEqualCase(t, rawText, 0, CF1102D)
}
