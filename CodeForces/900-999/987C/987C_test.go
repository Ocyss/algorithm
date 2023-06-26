package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// TestCF987C https://codeforces.com/problemset/problem/987/C
// https://codeforces.com/problemset/status/987/problem/C
// https://github.com/Ocyss
func TestCF987C(t *testing.T) {
	// just copy from website
	rawText := `
input
5
2 4 5 4 10
40 30 20 10 40
output
90
input
3
100 101 100
2 4 5
output
-1
input
10
1 2 3 4 5 6 7 8 9 10
10 13 11 14 15 12 13 13 18 13
output
33
input
10
802030518 598196518 640274071 983359971 71550121 96204862 799843967 446173607 796619138 402690754
23219513 68171337 12183499 5549873 73542337 66661387 79397647 34495917 31413076 50918417
output
85904709
`
	testutil.AssertEqualCase(t, rawText, 0, CF987C)
}
