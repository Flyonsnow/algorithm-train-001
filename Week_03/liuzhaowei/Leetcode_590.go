package main

/**
 * N叉树的后序遍历
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
 * 1.先遍历叶子节点
 * 2.遍历根节点
 */
type Node struct {
	Val      int
	Children []*Node
}

var afertResult []int

func postorder(root *Node) []int {
	afertResult = make([]int, 0)
	afterTraversal(root)
	return afertResult
}

func afterTraversal(root *Node) {
	if root != nil {
		for _, node := range root.Children {
			afterTraversal(node)
		}
		afertResult = append(afertResult, root.Val)
	}
}
