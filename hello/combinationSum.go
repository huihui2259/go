package main

import "sort"

// 输入: candidates = [2,3,5], target = 8
// 输出: [[2,2,2,2],[2,3,3],[3,5]]
var res [][]int
var tmp []int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	dfs(candidates, target, 0, 0)
	return res
}

func dfs(candidates []int, target int, sum, index int) {
	if sum == target {
		res = append(res, append([]int{}, tmp...))
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		sum += candidates[i]
		tmp = append(tmp, candidates[i])
		dfs(candidates, target, sum, i)
		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
	}
}
