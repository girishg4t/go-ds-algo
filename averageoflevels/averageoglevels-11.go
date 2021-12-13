package averageoflevels

import "fmt"

type Tree struct {
	root *TreeNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	var queue []*TreeNode = make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			sum = sum + queue[i].Val
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
		fmt.Println(result)
		queue = queue[size:]
	}
	return result
}

func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &TreeNode{Val: data}
	} else {
		t.root.insert(data)
	}
}

func (n TreeNode) insert(data int) {
	if data <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data}
		} else {
			n.Right.insert(data)
		}
	}
}

func CreateTree(arr []int) {
	var t Tree
	t.insert(3)
	t.insert(9)
	t.insert(20)
	t.insert(15)
	t.insert(7)
	fmt.Println(averageOfLevels(t.root))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return false
	}
	var pqueue []*TreeNode = make([]*TreeNode, 0)
	pqueue = append(pqueue, p)
	var q_queue []*TreeNode = make([]*TreeNode, 0)
	q_queue = append(q_queue, q)
	for len(pqueue) != 0 && len(q_queue) != 0 {
		if len(pqueue) != len(q_queue) {
			return false
		}
		size := len(pqueue)
		for i := 0; i < size; i++ {
			pnode := pqueue[i]
			qnode := q_queue[i]
			if pnode.Val != qnode.Val &&
				((pnode.Left == nil && qnode.Left != nil) ||
					(pnode.Right == nil && qnode.Right != nil)) {
				return false
			}
			if pnode.Left != nil && qnode.Left != nil {
				pqueue = append(pqueue, pnode.Left)
				q_queue = append(q_queue, qnode.Left)
			}
			if pnode.Right != nil && qnode.Right != nil {
				pqueue = append(pqueue, pnode.Right)
				q_queue = append(q_queue, qnode.Right)
			}
		}
		pqueue = pqueue[size:]
		q_queue = q_queue[size:]
	}
	return true
}
