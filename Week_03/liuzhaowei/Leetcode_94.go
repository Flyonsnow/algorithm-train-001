package main

/**
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 * 中序遍历
 */
//type TreeNode struct {
//	Val int
//	Left,Right *TreeNode
//}

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	traversal(root)
	return result
}

func traversal(root *TreeNode) {
	if root != nil {
		traversal(root.Left)
		result = append(result, root.Val)
		traversal(root.Right)
	}
}

//func main() {
//	t := TreeNode{1,nil,nil}
//	result := inorderTraversal(&t)
//
//	fmt.Println(result)
//}
