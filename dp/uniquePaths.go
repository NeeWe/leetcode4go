package main

import "fmt"

//一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/unique-paths
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func uniquePaths(m int, n int) int {
	// f(m, n) = f(m, n - 1) + f(m - 1, n)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, n)

	for i := 0; i < n; i ++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i ++ {
		dp[i][0] = 1
	}

	for i := 0; i < m; i ++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i ++ {
		for j := 1; j < m; j ++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[n-1][m-1]
}

func main() {
	fmt.Println(uniquePaths(1, 1))
}
