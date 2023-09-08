package data_struct

import (
	"sync"
)

// TreeNode 定义二叉树结构
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type LinkedNode struct {
	Next  *LinkedNode
	Value *TreeNode //value必须是树节点的指针类型，否则出队时会生成新节点，无法遍历原始树
}

type LinkedQueue struct {
	root *LinkedNode //链表队列起点
	size int
	lock sync.Mutex
}

func InitTree() *TreeNode {
	root := &TreeNode{Data: 0}
	root.Left = &TreeNode{Data: 1}
	root.Right = &TreeNode{Data: 2}
	root.Left.Left = &TreeNode{Data: 3}
	root.Left.Right = &TreeNode{Data: 4}
	root.Right.Left = &TreeNode{Data: 5}
	root.Right.Right = &TreeNode{Data: 6}
	return root
}

// PreOrder 前序遍历
func PreOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Data)
	res = append(res, PreOrder(root.Left)...)
	res = append(res, PreOrder(root.Right)...)
	return res
}

// MidOrder 中序遍历
func MidOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, MidOrder(root.Left)...)
	res = append(res, root.Data)
	res = append(res, MidOrder(root.Right)...)
	return res
}

// AfterOrder 后序遍历
func AfterOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, AfterOrder(root.Left)...)
	res = append(res, AfterOrder(root.Right)...)
	res = append(res, root.Data)
	return res
}

// LayerOrder 层次遍历：队列实现
func LayerOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := new(LinkedQueue)
	// 根节点入队
	queue.Add(root)
	for queue.size > 0 {
		elem := queue.Remove()
		res = append(res, elem.Data)
		if elem.Left != nil {
			queue.Add(elem.Left)
		}
		if elem.Right != nil {
			queue.Add(elem.Right)
		}
	}
	return res
}

// Add 入队
func (queue *LinkedQueue) Add(node *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	newNode := &LinkedNode{Value: node}
	if queue.root == nil {
		// 队列为空
		queue.root = newNode
	} else {
		// 队列不为空：遍历队列到队尾
		curNode := queue.root
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = newNode
	}
	queue.size += 1
}

// Remove 出队
func (queue *LinkedQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size == 0 {
		panic("队列为空")
	}
	topNode := queue.root.Value
	next := queue.root.Next
	queue.root = next
	queue.size -= 1
	return topNode
}
