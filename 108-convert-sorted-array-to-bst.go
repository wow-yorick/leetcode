package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[l/2]
	root.Left = sortedArrayToBST(nums[:l/2])
	root.Right = sortedArrayToBST(nums[l/2+1:])
	return root
}
