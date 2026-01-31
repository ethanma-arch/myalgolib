// LeetCode 76 最小覆盖子串
// 输入：第一行 s，第二行 t。输出：s 中包含 t 中所有字符的最短子串，若无则空串。
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

func solve(s, t string) string {
	sn := len(s)
	if sn == 0 || len(t) == 0 {
		return ""
	}

	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	needCount := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			needCount++
		}
	}

	var window [128]int
	l, r := 0, 0
	matched := 0
	minLen := sn + 1
	start := 0

	for r < sn {
		c := s[r]
		r++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				matched++
			}
		}

		for matched == needCount {
			if r-l < minLen {
				minLen = r - l
				start = l
			}
			c := s[l]
			l++
			if need[c] > 0 {
				if window[c] == need[c] {
					matched--
				}
				window[c]--
			}
		}
	}

	if minLen > sn {
		return ""
	}
	return s[start : start+minLen]
}

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

func main() {
	defer out.Flush()

	for {
		s, err1 := readNextValidLine(in)
		if err1 != nil && len(s) == 0 {
			break
		}
		t, err2 := readNextValidLine(in)
		if err2 != nil && len(t) == 0 {
			break
		}
		res := solve(s, t)
		fmt.Fprintln(out, res)
		if err1 != nil || err2 != nil {
			break
		}
	}
}
