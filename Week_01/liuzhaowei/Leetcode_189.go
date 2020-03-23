package main

/**
 * https://leetcode-cn.com/problems/rotate-array/
 * 旋转数组
 */

//1.for 循环k次，移动数组
func rotate(nums []int, k int) {
	l := len(nums)
	t := k % l
	for i := 0; i < t; i++ {
		temp := nums[l-1]
		for j := l - 1; j >= 0; j-- {
			if j == 0 {
				nums[j] = temp
			} else {
				nums[j] = nums[j-1]
			}
		}
	}
}
