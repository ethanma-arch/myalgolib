package main

func solve(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)

	left, right, top, bottom := 0, n-1, 0, m-1

	for left <= right && top <= bottom {

	}

	for i := 0; i < n-1; i++ {
		res = append(res, matrix[0][i])
	}

	for i := 0; i < m-1; i++ {
		res = append(res, matrix[i][i])
	}
}
