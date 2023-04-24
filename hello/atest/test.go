package main

import (
	"fmt"
	"sort"
)

var res [][]int
var tmp []int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.Ints(candidates)
	dfs(candidates, target, 0, 0)
	return res
}

func dfs(candidates []int, target int, sum, index int) {
	fmt.Println(index)
	tmp = append(tmp, candidates[index])
	fmt.Println(tmp)
	sum += candidates[index]
	fmt.Println(sum)
	if sum == target {
		res = append(res, append([]int{}, tmp...))
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		//sum -= candidates[i]
		dfs(candidates, target, sum, i)
		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
		fmt.Println("---")
		fmt.Println(tmp)
	}
}
