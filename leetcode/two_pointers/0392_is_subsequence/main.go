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

func solve(s, t string) bool {
	sp, tp := 0, 0
	sn, tn := len(s), len(t)

	for sp < sn && tp < tn {
		if s[sp] == t[tp] {
			sp++
		}
		tp++
	}
	return sp == sn
}

func main() {
	defer out.Flush()

	for {
		line, err := readNextValidLine(in)
		if err == io.EOF {
			break
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			break
		}
		s := parts[0]
		t := parts[1]
		res := solve(s, t)
		fmt.Fprintln(out, res)
	}
}
