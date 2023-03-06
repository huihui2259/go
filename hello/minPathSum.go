package main

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range dp {
		tmp := make([]int, len(grid[i]))
		dp[i] = tmp
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
