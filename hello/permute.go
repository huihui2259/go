package main

// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
var combinationRes [][]int
var combinationTmp []int
var mark []bool

func permute(nums []int) [][]int {
	combinationRes = [][]int{}
	for i := 0; i < len(nums); i++ {
		mark = append(mark, false)
	}
	premuteDfs(nums)
	return combinationRes
}
func premuteDfs(nums []int) {
	if len(combinationTmp) == len(nums) {
		combinationRes = append(combinationRes, append([]int{}, combinationTmp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if mark[i] {
			continue
		}
		tmp = append(tmp, nums[i])
		mark[i] = true
		premuteDfs(nums)
		mark[i] = false
		tmp = tmp[:len(nums)]
	}
}
