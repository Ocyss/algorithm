package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF883I https://codeforces.com/problemset/problem/883/I
// https://codeforces.com/problemset/status/883/problem/I
// https://github.com/Ocyss
func TestCF1832A(t *testing.T) {
	// just copy from website
	rawText := `
input
3
codedoc
gg
aabaa
output
YES
NO
NO
`
	testutil.AssertEqualCase(t, rawText, 0, CF1832A)
}
