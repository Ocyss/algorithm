// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/Ocyss/OI/generator/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, countSteppingNumbers, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-356/problems/count-stepping-numbers-in-range/
// https://leetcode.cn/problems/count-stepping-numbers-in-range/
