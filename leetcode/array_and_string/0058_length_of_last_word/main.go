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

func readLineAsString(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	line = strings.TrimSpace(line)
	if len(line) == 0 && err != nil {
		return "", err
	}
	return line, nil
}

func solve(s string) int {
	i := len(s) - 1
	res := 0

	for i >= 0 && s[i] == ' ' {
		i--
	}

	for i >= 0 && s[i] != ' ' {
		res++
		i--
	}

	return res
}

func main() {
	defer out.Flush()

	for {
		line, err := readLineAsString(in)
		if len(line) == 0 {
			if err != nil {
				break
			}
			continue
		}

		res := solve(line)
		fmt.Fprintln(out, res)

		if err == io.EOF {
			break
		}
	}
}
