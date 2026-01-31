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

func solve(haystack, needle string) int {
	n, m := len(haystack), len(needle)

	if m == 0 {
		return 0
	}
	if n < m {
		return -1
	}

	next := buildNext(needle)

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

// 在子串 pattern[0..i] 中，最长相等真前缀与真后缀的长度
func buildNext(p string) []int {
	n := len(p)
	next := make([]int, n)
	// j：当前匹配上的前缀长度，也是下次要比对的位置
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && p[i] != p[j] {
			j = next[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func main() {
	defer out.Flush()

	for {
		haystack, err := readLineAsString(in)
		if len(haystack) == 0 {
			if err != nil {
				break
			}
			continue
		}
		needle, err2 := readLineAsString(in)
		if len(needle) == 0 {
			if err2 != nil {
				break
			}
			continue
		}
		res := solve(haystack, needle)
		fmt.Fprintln(out, res)
		if err == io.EOF {
			break
		}
	}
}
