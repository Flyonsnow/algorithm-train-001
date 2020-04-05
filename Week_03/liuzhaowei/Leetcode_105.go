package main

/**
 * 从前序与中序遍历序列构造二叉树
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 *
 */

var idxMap map[int]int

var pre_idx int = 0

func buildTree(preorder []int, inorder []int) *TreeNode {
	idxMap = make(map[int]int, 0)

	for i, v := range inorder {
		idxMap[v] = i
	}
	return help(0, len(inorder), preorder, inorder)
}

func help(inLeft, inRight int, preorder []int, inorder []int) *TreeNode {
	if inLeft == inRight {
		return nil
	}

	rootVal := preorder[pre_idx]
	root := TreeNode{rootVal, nil, nil}

	index := idxMap[rootVal]

	pre_idx++
	root.Left = help(inLeft, index, preorder, inorder)
	root.Right = help(index+1, inRight, preorder, inorder)
	return &root
}

func main() {

}
