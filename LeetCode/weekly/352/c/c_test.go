// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"testing"

	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, continuousSubarrays, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-352/problems/continuous-subarrays/
// https://leetcode.cn/problems/continuous-subarrays/
