package main

import "math"

/**
 * 接雨水
 * https://leetcode-cn.com/problems/trapping-rain-water
 */

//暴力解优化，将leftMax和rightMax保存一个slice中
func trap2(height []int) int {
	sum := 0
	size := len(height)
	var leftMax = make([]int, size, size)
	var rightMax = make([]int, size, size)

	//从左往右，计算第i个左边的最大值
	for j := 1; j < size; j++ {
		leftMax[j] = int(math.Max(float64(leftMax[j-1]), float64(height[j-1])))
	}
	//从右往左，计算第i个右边的最大值
	for j := size - 2; j >= 0; j-- {
		rightMax[j] = int(math.Max(float64(rightMax[j+1]), float64(height[j+1])))
	}

	for i := 1; i < size-1; i++ {
		//取左右较小的值作为最低边界
		min := int(math.Min(float64(leftMax[i]), float64(rightMax[i])))
		if min > height[i] {
			sum += min - height[i]
		}
	}

	return sum
}

func main() {
	pa := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap2(pa)
}
