package main

import (
	"fmt"
	"sort"
	"strconv"
)

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}
	lastIndexOfNegativeNum := lastIndexOfNegativeNumber(nums)

	if lastIndexOfNegativeNum < 0 {
		if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 {
			return [][]int{{0, 0, 0}}
		} else {
			return [][]int{}
		}
	}

	results := make([][]int, 0)
	m := make(map[string]bool)
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		for j := i + 1; j < length-1 && nums[i]+nums[j] <= 0; j++ {
			if nums[i]+nums[j] > 0 || nums[i]+nums[j]+nums[length-1] < 0 {
				break
			}
			for k := j + 1; k < length; k++ {
				if nums[k] < 0 {
					continue
				}
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[k])
					exist, _ := m[key]
					if !exist {
						result := []int{nums[i], nums[j], nums[k]}
						results = append(results, result)
						m[key] = true
					}
				} else if sum > 0 {
					break
				}
			}
		}
	}
	return results
}

//lastIndexOfNegativeNumber 寻找有序数组中最后一个负数的位置，-1代表不存在负数
func lastIndexOfNegativeNumber(nums []int) int {
	index := -1
	length := len(nums)
	if nums[0] >= 0 {
		return -1
	}
	for i := 0; i < length; i++ {
		if nums[i]*index < 0 {
			return i - 1
		}
	}

	return length - 1
}

func main() {
	nums := []int{-4, -2, 6, 0, 2}
	sort.Ints(nums)
	//fmt.Println(lastIndexOfNegativeNumber(nums))
	fmt.Println(threeSum(nums))
}
