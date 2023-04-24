package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 6, 7}
	b := 7
	c := combinationSum(a, b)
	fmt.Println(c)
}
