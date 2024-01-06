// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, findValueOfPartition, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-350/problems/find-the-value-of-the-partition/
// https://leetcode.cn/problems/find-the-value-of-the-partition/
