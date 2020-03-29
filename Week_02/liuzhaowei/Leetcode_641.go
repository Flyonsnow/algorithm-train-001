package main

import "fmt"

/**
 * 双端队列
 * https://leetcode-cn.com/problems/design-circular-deque/
 */

// 1.使用slice实现
type MyCircularDeque struct {
	sl []int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {

	var sl = make([]int, 0, k)

	return MyCircularDeque{sl: sl}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	s := make([]int, 0, cap(this.sl))
	s = append(s, value)

	this.sl = append(s, this.sl...)

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.sl = append(this.sl, value)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	s := make([]int, len(this.sl)-1, cap(this.sl))
	copy(s, this.sl[1:])
	this.sl = s

	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.sl = this.sl[:len(this.sl)-1]
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.sl[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.sl[len(this.sl)-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if len(this.sl) == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if len(this.sl) == cap(this.sl) {
		return true
	}

	return false
}

/**
["MyCircularDeque","insertFront","insertLast","insertLast","getFront","deleteLast","getRear","insertFront",
"deleteFront","getRear","insertLast","isFull"]
[[3],[8],[8],[2],[],[],[],[9],[],[],[2],[]]
*/
func main() {
	obj := Constructor(3)
	param_1 := obj.InsertFront(8)
	param_2 := obj.InsertLast(8)
	param_3 := obj.InsertLast(2)
	param_4 := obj.GetFront()
	param_5 := obj.DeleteLast()
	param_6 := obj.GetRear()
	param_7 := obj.InsertFront(9)
	param_8 := obj.DeleteFront()
	param_9 := obj.GetRear()
	param_10 := obj.InsertLast(2)
	param_11 := obj.IsFull()

	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
	fmt.Println(param_7)
	fmt.Println(param_8)
	fmt.Println(param_9)
	fmt.Println(param_10)
	fmt.Println(param_11)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
