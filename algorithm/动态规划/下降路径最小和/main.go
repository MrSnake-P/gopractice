package main

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if a >= b && c >= b {
		return b
	} else {
		return c
	}
}

func minNums(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minFallingPathSum(matrix [][]int) int {
	lines, colums := len(matrix), len(matrix[0])
	dp := make([][]int, lines)
	for i:=0; i<lines; i++ {
		dp[i] = make([]int, colums)
		for j:=0; j<colums; j++ {
			dp[i][j] = 6666666
		}
	}

	dp[0][0] = matrix[0][0]

	//for j:=0; j<lines; j++ {
	//	dp[0][j] = matrix[0][j]
	//}
	//
	//for i:=1; i<lines; i++ {
	//	dp[i][0] = matrix[i][0] + min(dp[i-1][0],66666666, dp[i-1][1])
	//	dp[i][colums-1] = matrix[i][colums-1] + min(dp[i-1][colums-1],66666666, dp[i-1][colums-2])
	//}



	for i:=0; i<lines; i++ {
		for j:=0; j<colums; j++ {
			if i-1 < 0  {
				dp[0][j] = matrix[0][j]
				continue
			}
			if j-1 < 0 {
				dp[i][0] = matrix[i][0] + min(dp[i-1][j],66666666, dp[i-1][j+1])
			} else if j+1 > colums-1 {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j],66666666, dp[i-1][j-1])
			} else {
				dp[i][j] = matrix[i][j] + min(dp[i-1][j], dp[i-1][j-1], dp[i-1][j+1])
			}

		}
	}

	res := 6666666666666

	for j:=0; j<colums; j++ {
		res = minNums(res, dp[lines-1][j])
	}

	return res
}

func main() {
	//fmt.Println(min(4, 4, 3))
	//fmt.Println(min(-19, 6666, 57))
	//m := [][]int{{-19, 57}, {-40, -5}}
	//m1 := [][]int{{2,1,3}, {6,5,4}, {7,8,9}}
	//a := minFallingPathSum(m)
	//b := minFallingPathSumRecursion(m1)
	//fmt.Println(b)
	slice()
}
