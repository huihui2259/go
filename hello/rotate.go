package main

func rotate(matrix [][]int) {
	n := len(matrix)

	for l, r := 0, n-1; l < r; {
		for i := 0; i < n; i++ {
			matrix[l][i], matrix[r][i] = matrix[r][i], matrix[l][i]
		}
		l++
		r--
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
