package main

/**
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
 * 前序遍历
 */

var result2 []int

func preorderTraversal(root *TreeNode) []int {
	result2 = make([]int, 0)
	traversal2(root)
	return result
}

func traversal2(root *TreeNode) {
	if root != nil {
		result2 = append(result2, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
}

//func main() {
//	t := TreeNode{1,nil,nil}
//	result := inorderTraversal(&t)
//
//	fmt.Println(result)
//}
