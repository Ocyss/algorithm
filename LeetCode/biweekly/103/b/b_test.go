// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, findThePrefixCommonArray, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-103/problems/find-the-prefix-common-array-of-two-arrays/
// https://leetcode.cn/problems/find-the-prefix-common-array-of-two-arrays/
