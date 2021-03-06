package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	findPath(root, root.Val, sum, []int{root.Val}, &res)
	return res
}

func findPath(t *TreeNode, s, sum int, vals []int, ts *[][]int) {
	if t.Left == nil && t.Right == nil && s == sum {
		*ts = append(*ts, vals)
		return
	}
	if t.Left != nil {
		findPath(t.Left, s + t.Left.Val, sum, append([]int{}, append(vals, t.Left.Val)...), ts)
	}
	if t.Right != nil {
		findPath(t.Right, s + t.Right.Val, sum, append([]int{}, append(vals, t.Right.Val)...), ts)
	}
}