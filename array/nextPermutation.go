package main

import (
	"fmt"
	"sort"
)

//实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
//
//如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//
//必须 原地 修改，只允许使用额外常数空间。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/next-permutation
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func nextPermutation(nums []int)  {
	length := len(nums)
	if length <= 1 {
		fmt.Println(nums)
	} else {
		tmpI, tmpJ := -1, -1
		for i:= length - 1; i > 0; i -- {
			for j := i - 1; j >= 0 ; j -- {
				if nums[j] < nums[i] && (i < tmpI || tmpI == -1) {
					tmpI, tmpJ = i, j
				}
			}
		}
		if tmpI != tmpJ {
			nums[tmpI], nums[tmpJ] = nums[tmpJ], nums[tmpI]
			sort.Ints(nums[tmpI:])
			fmt.Println(nums)
			return
		}
		sort.Ints(nums)
		fmt.Println(nums)
	}
}

func main() {
	nextPermutation([]int{1, 2, 3})
}