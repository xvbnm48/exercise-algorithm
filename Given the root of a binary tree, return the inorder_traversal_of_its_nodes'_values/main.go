package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int

	cur := root
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
			continue
		}

		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

func main() {
	var root TreeNode
	root.Val = 1
	root.Left = nil
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}
	fmt.Println(inorderTraversal(&root)) // [1 3 2]
}
