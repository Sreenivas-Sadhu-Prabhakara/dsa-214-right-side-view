package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// RightSideView returns the node visible from the right at each level:
// DFS visiting RIGHT first, so the first arrival at each depth is the
// rightmost — the mirror of the left view.
func RightSideView(root *TreeNode) []int {
	view := []int{}
	var walk func(n *TreeNode, depth int)
	walk = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}
		if depth == len(view) {
			view = append(view, n.Val)
		}
		walk(n.Right, depth+1)
		walk(n.Left, depth+1)
	}
	walk(root, 0)
	return view
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(2, nil, tn(5, nil, nil)), tn(3, nil, tn(4, nil, nil)))
	fmt.Println(RightSideView(root))
}
