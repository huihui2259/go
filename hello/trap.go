package main

// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6

// 输入：height = [4,2,0,3,2,5]
// 输出：9
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left_height := make([]int, n)
	right_height := make([]int, n)

	left_height[0] = height[0]
	for index, value := range height {
		if index == 0 {
			continue
		}
		left_height[index] = max(left_height[index-1], value)
	}

	right_height[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		right_height[i] = max(right_height[i+1], height[i])
	}
	sum := 0
	for i := 0; i < len(height); i++ {
		sum += min(left_height[i], right_height[i]) - height[i]
	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
