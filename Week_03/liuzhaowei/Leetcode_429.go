package main

/**
 * N叉树的层序遍历
 * https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
 * 1.使用map存储返回的接口，key：层级，v：数组
 * 2.递归遍历N叉树
 */
var levelMap map[int][]int

func levelOrder(root *Node) [][]int {
	levelMap = make(map[int][]int, 0)
	levelTraversal(0, root)
	var back [][]int = make([][]int, 0)
	for i := 0; i < len(levelMap); i++ {
		back = append(back, levelMap[i])
	}
	return back
}

func levelTraversal(level int, root *Node) {
	if root != nil {
		if r, ok := levelMap[level]; ok {
			levelMap[level] = append(r, root.Val)
		} else {
			levelMap[level] = []int{root.Val}
		}
		for _, node := range root.Children {
			levelTraversal(level+1, node)
		}
	}
}
