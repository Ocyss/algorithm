package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// https://github.com/Ocyss
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	q := [][2]*TreeNode{{root, nil}}
	root.Val = 0
	for len(q) > 0 {
		temp := q
		q = nil
		ans := 0
		for _, v := range temp {
			ans += v[0].Val
		}
		f := make([]int, len(temp))
		for i, v := range temp {
			f[i] = ans
			if v[1] != nil && v[1].Left != nil {
				f[i] -= v[1].Left.Val
			}
			if v[1] != nil && v[1].Right != nil {
				f[i] -= v[1].Right.Val
			}
			if v[0].Left != nil {
				q = append(q, [2]*TreeNode{v[0].Left, v[0]})
			}
			if v[0].Right != nil {
				q = append(q, [2]*TreeNode{v[0].Right, v[0]})
			}
		}
		for i, v := range temp {
			v[0].Val = f[i]
		}
	}
	return root
}
