package offer

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return false
	}
	if len(postorder) == 1 {
		return true
	}
	return dfs1(0, len(postorder)-1, postorder)
}

func dfs1(i, j int, order []int) bool {
	if i >= j {
		return true
	}
	rootValue := order[j]
	pos1 := j - 1
	for pos1 >= i && order[pos1] > rootValue {
		pos1--
	}
	pos2 := pos1
	for pos2 >= i && order[pos2] < rootValue {
		pos2--
	}
	if pos2 != i-1 {
		return false
	}
	return dfs1(i, pos1, order) && dfs1(pos1+1, j-1, order)
}
