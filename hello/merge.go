package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	res = append(res, intervals[0])
	for i := 0; i < len(intervals); i++ {
		tmp := res[len(res)-1]
		if tmp[1] >= intervals[i][0] {
			// 需要合并
			res = res[:len(res)-1]
			newTmp := make([]int, 2)
			newTmp[0] = tmp[0]
			newTmp[1] = mergeMax(tmp[1], intervals[i][1])
			res = append(res, newTmp)
		} else {
			res = append(res, intervals[i])
		}

	}
	return res
}

func mergeMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
