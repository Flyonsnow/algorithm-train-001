package main

/**
 * N叉树的前序遍历
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
 * 1.先遍历根节点
 * 2.遍历叶子节点
 */
var preResult []int

func preorder(root *Node) []int {
	preResult = make([]int, 0)
	preTraversal(root)
	return preResult
}

func preTraversal(root *Node) {
	if root != nil {
		preResult = append(preResult, root.Val)
		for _, node := range root.Children {
			preTraversal(node)
		}
	}
}
