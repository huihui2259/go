package offer

func myPow(x float64, n int) float64 {
	res := 1.0
	i := n
	for ; i != 0; i /= 2 {
		if i%2 != 0 {
			res *= x
		}
		x *= x
	}
	if n < 0 {
		return 1 / res
	}
	return res
}
