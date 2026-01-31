package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func readNextValidLine(r *bufio.Reader) (string, error) {
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			return line, err
		}
		if err != nil {
			return "", err
		}
	}
}

func solve(s string) int {
	n := len(s)
	var lastPos [128]int
	for i := range lastPos {
		lastPos[i] = -1
	}

	l := 0
	maxLen := 0

	for r := 0; r < n; r++ {
		char := s[r]
		if lastPos[char] >= l {
			l = lastPos[char] + 1
		}
		lastPos[char] = r
		maxLen = max(maxLen, r-l+1)
	}

	return maxLen

}

func main() {
	defer out.Flush()
	for {
		line, err := readNextValidLine(in)
		if err != nil && len(line) == 0 {
			break
		}
		res := solve(line)
		fmt.Fprintln(out, res)
		if err != nil {
			break
		}
	}
}
