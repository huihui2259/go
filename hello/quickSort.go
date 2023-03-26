package main

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j := low, high
	pivot := nums[(low+high)/2]
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = pivot
	quickSort(nums, low, i-1)
	quickSort(nums, i+1, high)
}
