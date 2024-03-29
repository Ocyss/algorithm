// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, findPrefixScore, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-102/problems/find-the-score-of-all-prefixes-of-an-array/
// https://leetcode.cn/problems/find-the-score-of-all-prefixes-of-an-array/
