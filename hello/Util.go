package main

import (
	"fmt"
)

// 一些辅助函数

//打印二维切片
func PrintSlice(slice [][]int) {
	for _, v := range slice {
		fmt.Println(v)
	}
}
