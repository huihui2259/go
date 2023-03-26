package main

func sortColors(nums []int) {
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
	for i := pos; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[pos], nums[i] = nums[i], nums[pos]
			pos++
		}
	}
}
