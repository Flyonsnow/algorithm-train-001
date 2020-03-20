package main

/**
 *	https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
 *	移除重复元素
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//数组长度
	count := 1
	temp := nums[0]
	for i, _ := range nums {
		if nums[i] > temp {
			nums[count] = nums[i]
			count = count + 1
		}
		temp = nums[i]
	}
	return count
}
