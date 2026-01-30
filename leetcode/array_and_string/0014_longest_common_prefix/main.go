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

func readLineAsStrings(r *bufio.Reader) ([]string, error) {
	line, err := r.ReadString('\n')
	line = strings.TrimSpace(line)
	if len(line) == 0 && err != nil {
		return nil, err
	}
	parts := strings.Fields(line)
	strs := make([]string, 0, len(parts))
	for _, p := range parts {
		strs = append(strs, p)
	}
	return strs, nil
}

func solve(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	m := len(strs[0])

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	defer out.Flush()

	for {
		strs, err := readLineAsStrings(in)
		if len(strs) == 0 {
			if err != nil {
				break
			}
			continue
		}
		res := solve(strs)
		fmt.Fprintln(out, res)

		if err == io.EOF {
			break
		}
	}
}
