package main

import "fmt"

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//注意：给定 n 是一个正整数。
func climbStairs(n int) int {

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	a1 := 1
	a2 := 2

	for i := 3; i < n+1; i++ {
		a1, a2 = a2, a1+a2
	}

	return a2
}

func main() {
	fmt.Println(climbStairs(4))
}
