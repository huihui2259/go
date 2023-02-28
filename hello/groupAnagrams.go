package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := map[string][]string{}
	for _, v := range strs {
		arr := []byte(v)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		key := string(arr)
		m[key] = append(m[key], v)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
