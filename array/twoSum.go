package main

import "fmt"

//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/two-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index,value := range nums {
		v, ok := m[target - value]
		if ok {
			return []int {v, index}
		}
		m[value] = index
	}
	return []int{}
}

func main() {
	sum := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(sum)
}
