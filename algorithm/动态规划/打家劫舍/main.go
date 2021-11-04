package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	// 定义dp数组
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for k:=2; k < n+1; k++ {
		dp[k] = maxTwoValue(dp[k-1], nums[k-1] + dp[k-2])
	}
	return dp[n]
}

func maxTwoValue(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main()  {
	fmt.Println(rob([]int{1,2,3,1}))
}