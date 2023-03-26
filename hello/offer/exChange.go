package offer

// 输入；1 2 3 4
func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]%2 != 0 {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		for i < j && nums[j]%2 == 0 {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	return nums
}
