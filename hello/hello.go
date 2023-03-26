package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] < nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {

	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	b := minPathSum(a)
	fmt.Println(b)
}
