package main

import (
	"fmt"
	"sort"
)

var res [][]int
var tmp []int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	dfs2(candidates, target, 0, 0)
	return res
}

func dfs2(candidates []int, target int, sum, index int) {
	fmt.Println(tmp)
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
		dfs2(candidates, target, sum, i)
		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
	}
}
