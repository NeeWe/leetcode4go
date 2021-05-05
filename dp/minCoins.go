package main

import (
	"../util"
	"fmt"
)

const UnReachable = 99999

var dp []int

func minCoins(target int) int {

	if target < 8 {
		dp = make([]int, 8)
	} else {
		dp = make([]int, target+1)
	}
	dp[0] = UnReachable
	dp[1] = UnReachable
	dp[2] = 1
	dp[3] = UnReachable
	dp[4] = 2
	dp[5] = 1
	dp[6] = 3
	dp[7] = 1

	result := UnReachable

	if target < 0 {
		result = UnReachable
	} else if target < 8 {
		result = dp[target]
	} else {
		for i := 8; i <= target; i++ {
			decrice2 := dp[i-2] + 1
			decrice5 := dp[i-5] + 1
			decrice7 := dp[i-7] + 1
			result = util.Min(util.Min(decrice2, decrice5), decrice7)
			if result >= UnReachable {
				dp[i] = UnReachable
			} else {
				dp[i] = result
			}
		}
	}

	if result >= UnReachable {
		return -1
	} else {
		return result
	}
}

func main() {
	target := 139
	fmt.Println(minCoins(target))
}
