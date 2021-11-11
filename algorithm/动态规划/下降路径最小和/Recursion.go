package main

import "math"

var memo [][]int

func minFallingPathSumRecursion(matrix [][]int) int {
	n := len(matrix)
	res := math.MaxInt64
	// 添加备忘录
	memo = make([][]int, n)
	for i:=0; i<n; i++ {
		memo[i] = make([]int, n)
		for j:=0; j<n; j++ {
			memo[i][j] = 6666666
		}
	}

	for j:=0; j<n; j++ {
		res = minNums(res, dp(matrix, n-1, j))
	}

	return res
}

func dp(matrix [][]int, i,j int) int{
	if i<0 || j<0 || i>=len(matrix)|| j>= len(matrix[0]) {
		return 99999
	}

	if i==0 {
		return matrix[0][j]
	}

	if memo[i][j] != 6666666 {
		return memo[i][j]
	}

	memo[i][j] = matrix[i][j] + min(dp(matrix, i-1, j-1), dp(matrix, i-1, j), dp(matrix, i-1, j+1))

	return memo[i][j]
}