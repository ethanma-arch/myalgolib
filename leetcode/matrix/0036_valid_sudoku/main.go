package main

import (
	"bufio"
	"fmt"
	"leetcode/common_lib"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(board [][]byte) bool {
	var rows, cols [9][9]int
	var box [3][3][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			idx := board[i][j] - '1'
			if idx < 0 || idx > 8 {
				continue
			}
			rows[i][idx]++
			cols[j][idx]++
			box[i/3][j/3][idx]++
			if rows[i][idx] > 1 || cols[j][idx] > 1 || box[i/3][j/3][idx] > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	defer out.Flush()
	for {
		board := make([][]byte, 9)
		for i := 0; i < 9; i++ {
			line, err := common_lib.ReadNextValidLine(in)
			if len(line) == 0 && err != nil {
				return
			}
			parts := strings.Fields(line)
			if len(parts) != 9 {
				return
			}
			row := make([]byte, 9)
			for j, p := range parts {
				if p == "." {
					row[j] = '.'
				} else {
					row[j] = p[0]
				}
			}
			board[i] = row
		}
		ok := solve(board)
		fmt.Fprintln(out, ok)
	}
}
