package main

import (
	"bufio"
	"fmt"
	"io"
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
			return line, nil
		}
		if err != nil {
			return "", err
		}
	}
}

func solve(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	l, r := 0, n-1

	for l < r {
		for l < n && !isValid(s[l]) {
			l++
		}
		for r >= 0 && !isValid(s[r]) {
			r--
		}

		if l < n && r >= 0 && s[l] != s[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}

func isValid(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}

func main() {
	defer out.Flush()
	for {
		line, err := readNextValidLine(in)

		if err == io.EOF {
			break
		}

		res := solve(line)
		fmt.Fprintln(out, res)
	}
}
