package main

// https://github.com/Ocyss
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *testutil.ListNode) *testutil.ListNode {
	if head == nil {
		return head
	}
	if v := dfs(head); v != 0 {
		return &testutil.ListNode{Next: head, Val: v}
	}
	return head
}

func dfs(head *testutil.ListNode) int {
	r := head.Val * 2
	if head.Next != nil {
		r += dfs(head.Next)
	}
	head.Val = r % 10
	return r / 10
}
