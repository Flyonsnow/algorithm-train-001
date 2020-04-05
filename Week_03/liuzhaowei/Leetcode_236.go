package main

/**
 * 二叉树的最近公共祖先
 * https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 * 1.递归遍历p节点，找到p节点的路径
 * 2.递归遍历q节点，找到q节点的路径
 * 3.找到p和q的最近公共节点
 */

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := treeTraversal(root, p.Val)
	qPath := treeTraversal(root, q.Val)
	c := 0
	for i := 0; i < len(pPath); i++ {
		if i < len(qPath) {
			if pPath[i] == qPath[i] {
				c = pPath[i]
			}
		}
	}
	return &TreeNode{c, nil, nil}
}

func treeTraversal(root *TreeNode, target int) []int {

	if root != nil {
		if root.Val == target {
			return []int{root.Val}
		}
		lf := treeTraversal(root.Left, target)
		rg := treeTraversal(root.Right, target)

		if lf != nil {
			var back []int = make([]int, 0)
			back = append(back, root.Val)
			back = append(back, lf...)
			return back
		}
		if rg != nil {
			var back []int = make([]int, 0)
			back = append(back, root.Val)
			back = append(back, rg...)
			return back
		}
	}
	return nil
}

func main() {
	root := &TreeNode{3, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{1, nil, nil}
	root.Left.Left = &TreeNode{6, nil, nil}
	root.Left.Right = &TreeNode{2, nil, nil}
	root.Left.Right.Left = &TreeNode{7, nil, nil}
	root.Left.Right.Right = &TreeNode{4, nil, nil}
	root.Right.Left = &TreeNode{0, nil, nil}
	root.Right.Right = &TreeNode{8, nil, nil}

	p := &TreeNode{0, nil, nil}
	q := &TreeNode{8, nil, nil}
	lowestCommonAncestor(root, p, q)
}
