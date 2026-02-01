// LeetCode 54 螺旋矩阵
// 输入：若干行，每行若干整数，表示矩阵的一行。输出：按螺旋顺序遍历的元素，一行空格分隔。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"leetcode/common_lib"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)
	left, right, top, bottom := 0, n-1, 0, m-1

	for left <= right && top <= bottom {
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for j := right; j >= left; j-- {
			res = append(res, matrix[bottom][j])
		}
		bottom--
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

func main() {
	defer out.Flush()

	var matrix [][]int
	for {
		line, err := common_lib.ReadNextValidLine(in)
		if len(line) == 0 && err != nil {
			break
		}
		parts := strings.Fields(line)
		row := make([]int, 0, len(parts))
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			row = append(row, num)
		}
		if len(row) == 0 {
			continue
		}
		matrix = append(matrix, row)
		if err != nil {
			break
		}
	}
	if len(matrix) == 0 {
		return
	}
	res := solve(matrix)
	for i, v := range res {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
