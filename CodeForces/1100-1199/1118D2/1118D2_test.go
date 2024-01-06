package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/main/testutil"
)

// TestCF1118D2 https://codeforces.com/problemset/problem/1118/D2
// https://codeforces.com/problemset/status/1118/problem/D2
// https://github.com/Ocyss
func TestCF1118D2(t *testing.T) {
	// just copy from website
	rawText := `
input
5 8
2 3 1 1 2
output
4
input
7 10
1 3 4 2 1 4 2
output
2
input
5 15
5 5 5 5 5
output
1
input
5 16
5 5 5 5 5
output
2
input
5 26
5 5 5 5 5
output
-1
input
100 550
12 7 8 16 13 6 12 6 10 12 13 10 6 12 9 8 5 13 7 13 5 14 10 13 9 6 14 14 6 11 13 13 8 3 8 13 12 8 13 8 14 10 15 8 12 8 10 13 13 13 7 8 12 9 7 10 16 10 10 4 9 9 11 8 8 13 8 15 11 8 9 6 6 16 12 11 9 8 10 9 12 8 11 9 8 10 7 13 6 13 10 9 15 9 3 11 5 17 7 13
output
9`
	testutil.AssertEqualCase(t, rawText, 0, CF1118D2)
}
