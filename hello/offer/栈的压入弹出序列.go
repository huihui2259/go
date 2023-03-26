package offer

func validateStackSequences(pushed []int, popped []int) bool {
	res := []int{}
	i := 0
	for _, value := range pushed {
		res = append(res, value)
		for len(res) > 0 && res[len(res)-1] == popped[i] {
			i++
			res = res[:len(res)-1]
		}
	}
	if i == len(popped) {
		return true
	}
	return false
}
