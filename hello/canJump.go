package main

func canJump(nums []int) bool {
	jump := nums[0]
	for k, v := range nums {
		if k <= jump {
			jump = max(jump, k+v)
			if jump >= len(nums) {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
