package dfs

import "math"

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt
	maxGain(root)
	return maxSum
}

// 获得当前节点的最大贡献值
func maxGain(node *TreeNode) int {
	//空节点的贡献值为0
	if node == nil {
		return 0
	}
	//递归计算左右子节点的最大贡献值
	leftGain := max(maxGain(node.Left), 0)
	rightGain := max(maxGain(node.Right), 0)
	//当前节点的最大路径和 取决于该节点的值与左右子节点的最大贡献值
	priceCurPath := node.Val + leftGain + rightGain
	//更新最大路径和
	maxSum = max(maxSum, priceCurPath)

	//返回当前节点的最大贡献值
	return node.Val + max(leftGain, rightGain)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
