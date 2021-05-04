package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
func maxSubArray(nums []int) int {
	length := len(nums)
	m := nums[0]
	for i := 1; i < length; i++ {
		nums[i] = max(nums[i] + nums[i - 1], nums[i])
		m = max(nums[i], m)
	}
	return m
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
