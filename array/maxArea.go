package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func maxArea(height []int) int {
	left,right := 0,len(height)-1
	maxContain := 0
	for left != right {
		heiLeft,heiRight := height[left],height[right]
		containerVal := min(heiLeft,heiRight)*(right-left)
		if containerVal > maxContain {
			maxContain = containerVal
		}
		if heiLeft > heiRight {
			right --
		} else {
			left ++
		}
	}
	return maxContain
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
//2021-04-27 23:06:50.4359783 +0800 CST m=+0.022013601
//705634720
//2021-04-27 23:06:53.9869929 +0800 CST m=+3.573028201
func main() {
	bytes, _ := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\int.txt")

	var ints []int
	json.Unmarshal(bytes, &ints)
	fmt.Println(time.Now())
	//fmt.Println(maxArea(ints))
	fmt.Println(maxArea([]int{1, 2, 3, 4}))
	fmt.Println(time.Now())
}
