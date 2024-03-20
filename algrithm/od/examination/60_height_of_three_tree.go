package examination

import (
	"fmt"
)

type TreeNode struct {
	val    int
	left   *TreeNode
	middle *TreeNode
	right  *TreeNode
}

func HeightOfThreeTree() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	var root *TreeNode
	//输入节点并插入树中
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		root = insertNode(root, nums[i])
	}
	fmt.Println(getHeight(root))
}

// 递归插入节点
func insertNode(root *TreeNode, x int) *TreeNode {
	if root == nil {
		return &TreeNode{val: x}
	}
	curVal := root.val
	if x < curVal-500 {
		root.left = insertNode(root.left, x)
	} else if x > curVal+500 {
		root.right = insertNode(root.right, x)
	} else {
		root.middle = insertNode(root.middle, x)
	}
	return root
}

func getHeight(root *TreeNode) int {
	//递归出口
	if root == nil {
		return 0
	}
	//递归式
	return 1 + max(max(getHeight(root.left), getHeight(root.middle)), getHeight(root.right))
}
