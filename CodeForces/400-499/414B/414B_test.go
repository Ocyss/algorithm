package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF414B https://codeforces.com/problemset/problem/414/B
// https://codeforces.com/problemset/status/414/problem/B
// https://github.com/Ocyss
func TestCF414B(t *testing.T) {
	// just copy from website
	rawText := `
input
1478 194
output
5`
	testutil.AssertEqualCase(t, rawText, 0, CF414B)
}
