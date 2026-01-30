package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)

	res := make([][]byte, numRows)
	for i := range res {
		res[i] = make([]byte, 0)
	}

	flag := true

	row := 0
	for i := 0; i < n; i++ {
		res[row] = append(res[row], s[i])

		if flag {
			row++
		} else {
			row--
		}

		if row == 0 || row == numRows-1 {
			flag = !flag
		}
	}

	ans := make([]byte, 0, n)
	for _, r := range res {
		ans = append(ans, r...)
	}
	return string(ans)
}

func main() {
	defer out.Flush()

	in.Split(bufio.ScanWords)

	for in.Scan() {
		s := in.Text()
		if !in.Scan() {
			break
		}
		numRows, _ := strconv.Atoi(in.Text())
		res := solve(s, numRows)
		fmt.Fprintln(out, res)
	}
}
