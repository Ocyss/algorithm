// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, minimumBeautifulSubstrings, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-108/problems/partition-string-into-minimum-beautiful-substrings/
// https://leetcode.cn/problems/partition-string-into-minimum-beautiful-substrings/