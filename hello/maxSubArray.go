package main

// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func maxSubArray(nums []int) int {
	res := nums[0]
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for k, v := range nums {
		if k == 0 {
			continue
		}
		if dp[k-1] >= 0 {
			dp[k] = dp[k-1] + v
		} else {
			dp[k] = v
		}
		res = max(res, dp[k])
	}

	return res
}
func subArraymax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
