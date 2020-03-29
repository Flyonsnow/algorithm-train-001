package main

import "fmt"

func main() {
	var nums = []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

//循环暴力解
func moveZeroes(nums []int) {
	ct := 0
	for _, v := range nums {
		if v == 0 {
			ct++
		}
	}
	for x := 0; x < ct; x++ {
		for i, v := range nums {
			if v == 0 {
				moveForward(nums, i+1)
			}
		}
	}
	//最后补零
	for i := 1; i <= ct; i++ {
		nums[len(nums)-i] = 0
	}
}

//向前移动  [0,0,1,3,12]
func moveForward(nums []int, index int) {
	for i := index; i < len(nums); i++ {
		nums[i-1] = nums[i]
	}
}
