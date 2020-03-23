package main

import "fmt"

/**
 * 合并两个有序链表
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/
 */

func main() {
	node1 := ListNode{-9, nil}
	node2 := ListNode{3, nil}

	nodea := ListNode{5, nil}
	nodeb := ListNode{7, nil}

	node1.Next = &node2
	nodea.Next = &nodeb

	mergeNode := mergeTwoLists(&node1, &nodea)

	fmt.Println(mergeNode)

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	//排除空情况
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	hasNext := true
	//遍历一个链表，将遍历完的node添加到另一个链表上
	newNode := l2
	for hasNext {
		if newNode == nil {
			hasNext = false
		} else {
			oneNode := newNode
			newNode = newNode.Next
			oneNode.Next = nil
			l1 = addNode(l1, oneNode)
		}
	}
	return l1
}

func addNode(l1 *ListNode, oneNode *ListNode) *ListNode {
	//如果新节点小于当前节点，之间挂在前面
	if oneNode.Val < l1.Val {
		oneNode.Next = l1
		return oneNode
	}
	beforeNode := l1
	c := true
	for c {
		afterNode := beforeNode.Next
		if oneNode.Val >= beforeNode.Val {
			//将tempNode挂载befor和after之间
			if afterNode == nil {
				beforeNode.Next = oneNode
				c = false
			} else if oneNode.Val < afterNode.Val {
				beforeNode.Next = oneNode
				oneNode.Next = afterNode
				c = false
			}
		}
		beforeNode = afterNode
	}
	return l1
}

type ListNode struct {
	Val  int
	Next *ListNode
}
