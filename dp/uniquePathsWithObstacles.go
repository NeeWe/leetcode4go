package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	n := len(obstacleGrid)
	m := len(obstacleGrid[0])

	if m == 0 || n == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < n; i ++ {
		if obstacleGrid[i][0] != 1 {
			obstacleGrid[i][0] = -1
		} else {
			break
		}
	}

	for i := 0; i < m; i ++ {
		if obstacleGrid[0][i] != 1 {
			obstacleGrid[0][i] = -1
		} else {
			break
		}
	}

	for i := 1; i < n; i ++ {
		for j := 1; j < m; j ++ {
			tmp1 := obstacleGrid[i][j-1]
			tmp2 := obstacleGrid[i-1][j]

			if tmp1 == 1 {
				tmp1 = 0
			}

			if tmp2 == 1 {
				tmp2 = 0
			}
			if obstacleGrid[i][j] != 1 {
				obstacleGrid[i][j] = tmp1 + tmp2
			}
		}
	}
	result := -obstacleGrid[n-1][m-1]
	if result < 0 {
		return 0
	} else {
		return result
	}
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0},{1, 1},{0, 0}}))
}
