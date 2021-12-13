package dfs

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) int {
	var stack []*TreeNode = make([]*TreeNode, 0)
	stack = append(stack, root)
	sum := 0
	for len(stack) != 0 {
		size := len(stack)
		for i := size - 1; i < 0; i-- {
			sum = sum + stack[i].Val
			node := stack[i]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		stack = stack[size:]
	}
	return sum
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max = 0
	return depth(root, max)
}

func depth(node *TreeNode, max int) int {
	leftMax := depth(node.Left, max)
	rightMax := depth(node.Right, max)
	max = int(math.Max(float64(leftMax+rightMax), float64(max)))
	return max + 1
}
