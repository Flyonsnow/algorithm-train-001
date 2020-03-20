package main

/**
 * 合并两个有序数组
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 * 输出: [1,2,2,3,5,6]
 */

/**
 * 解题思路：
 * 1.遍历nums2，拿到nums2中的元素
 * 2.将nums2中的元素放到nums1对应的位置上
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	length := m
	for i := 0; i < n; i++ {
		length = addItem(nums1, nums2[i], length)
	}

}

func addItem(nums1 []int, item int, length int) int {
	if length == 0 {
		nums1[0] = item
		return length + 1
	}

	for i := 0; i < length; i++ {
		if item < nums1[i] {
			move(nums1, 0, length)
			nums1[0] = item
			break
		} else if item >= nums1[i] {
			if i+1 == length {
				nums1[i+1] = item
				break
			} else if item < nums1[i+1] {
				move(nums1, i+1, length)
				nums1[i+1] = item
				break
			}
		}
	}
	return length + 1
}

func move(nums1 []int, start int, length int) {
	for i := 0; i < length-start; i++ {
		nums1[length-i] = nums1[length-i-1]
	}
}
