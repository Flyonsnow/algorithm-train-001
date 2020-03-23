package main

/**
 * 两数之和
 * https://leetcode-cn.com/problems/two-sum/
 */
//一次hash
func twoSum2(nums []int, target int) []int {
	var back []int
	m := make(map[int]int)
	for i, v := range nums {
		complete := target - v
		if v, ok := m[complete]; ok {
			back = append(back, v)
			back = append(back, i)
			break
		}
		m[v] = i
	}
	return back
}

//暴力解，两次for循环
func twoSum(nums []int, target int) []int {
	var back []int
	for i, v := range nums {
		back = []int{}
		back = append(back, i)
		second := target - v

		for j, k := range nums {
			if second == k && i != j {
				back = append(back, j)
				break
			}
		}

		if len(back) == 2 {
			break
		}
	}
	return back
}
