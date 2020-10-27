package main

// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var values []int
	inorderWalk(root, &values)
	prev := 0
	for i, val := range values {
		if i != 0 && prev >= val {
			return false
		}
		prev = val
	}
	return true
}

func inorderWalk(node *TreeNode, save *[]int) {
	if node != nil {
		inorderWalk(node.Left, save)
		*save = append(*save, node.Val)
		inorderWalk(node.Right, save)
	}
}
