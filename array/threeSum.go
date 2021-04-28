package main

import (
	"fmt"
	"sort"
)

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	firstIndexOfZero := firstIndexOfZero(nums)
	if firstIndexOfZero == -1 {
		return [][]int{}
	}
	results := make([][]int, 0)
	for i := 0; nums[i] <= 0; i++ {
		for j := firstIndexOfZero; j < length-1; j++ {
			for k := j + 1; k < length; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					result := []int{nums[i], nums[j], nums[k]}
					results = append(results, result)
				}
			}
		}
	}
	return results
}

func firstIndexOfZero(nums []int) int {
	index := -1
	length := len(nums)
	for i := 0; i < length && nums[i] < 0; i++ {
		if nums[i] == 0 {
			index = i
		}
	}
	return index
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
