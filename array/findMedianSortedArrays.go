package main

import "fmt"

//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	length1 := len(nums1)
	length2 := len(nums2)

	totalLength := length1 + length2

	if totalLength <= 1 {
		if length1 == 1 {
			return float64(nums1[0])
		} else if length2 == 1{
			return float64(nums2[0])
		} else {
			return 0.0
		}
	}

	middle := totalLength / 2
	pointer1, pointer2, step := 0, 0, 0
	sumArray := make([]int, middle + 1)

	for ; step <= middle && pointer1 < length1 && pointer2 < length2; {
		if nums1[pointer1] <= nums2[pointer2] {
			sumArray[step] = nums1[pointer1]
			pointer1 ++
		} else if nums1[pointer1] > nums2[pointer2] {
			sumArray[step] = nums2[pointer2]
			pointer2 ++
		}
		step ++
	}


	for ; step <= middle && pointer1 < length1; {
		sumArray[step] = nums1[pointer1]
		pointer1 ++
		step ++
	}

	for ; step <= middle && pointer2 < length2; {
		sumArray[step] = nums2[pointer2]
		pointer2 ++
		step ++
	}

	if totalLength% 2 == 0 {
		// 取 totalLength / 2 处 + totalLength / 2 + 1 处的平均值
		return float64(sumArray[middle - 1] + sumArray[middle]) / 2

	} else {
		// 取 totalLength / 2 处的值
		return float64(sumArray[middle])
	}
}

func main() {
	value := findMedianSortedArrays([]int{3, 4}, []int{})
	fmt.Println(value)
}
