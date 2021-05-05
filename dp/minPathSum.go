package main

import "fmt"
import "../util"

//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	for i := 1; i < m; i ++ {
		grid[i][0] = grid[i - 1][0] + grid[i][0]
	}

	for i := 1; i < n; i ++ {
		grid[0][i] = grid[0][i - 1] + grid[0][i]
	}

	// f(x, y) = grid[x-1][y-1] + min(f(x-1, y),f(x, y - 1))

	for i := 1; i < m; i ++ {
		for j := 1; j < n; j ++ {
  			grid[i][j] = grid[i][j] + util.Min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]

}

func main() {
	fmt.Println(minPathSum([][]int{{1,2}}))
}