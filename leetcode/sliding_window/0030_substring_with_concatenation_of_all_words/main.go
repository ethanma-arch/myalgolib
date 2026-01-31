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

func solve(s string, words []string) []int {
	m := len(words)
	n := len(words[0])
	sl := len(s)

	res := []int{}

	need := make(map[string]int)
	for _, word := range words {
		need[word]++
	}

	for i := 0; i < n; i++ {
		l, r := i, i
		window := make(map[string]int)
		matched := 0

		for r+n <= sl {
			word := s[r : r+n]
			r += n

			if _, ok := need[word]; !ok {
				l = r
				window = make(map[string]int)
				matched = 0
			} else {
				window[word]++
				matched++

				for window[word] > need[word] {
					lword := s[l : l+n]
					window[lword]--
					matched--
					l += n
				}

				// 如果有效单词数量等于 words 长度，记录结果
				if matched == m {
					res = append(res, l)
					lword := s[l : l+n]
					window[lword]--
					matched--
					l += n
				}
			}
		}
	}

	return res
}

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

func main() {
	defer out.Flush()

	for {
		line1, err1 := readNextValidLine(in)
		if len(line1) == 0 && err1 != nil {
			break
		}
		s := strings.TrimSpace(line1)

		line2, err2 := readNextValidLine(in)
		if len(line2) == 0 && err2 != nil {
			break
		}
		parts := strings.Fields(line2)
		words := make([]string, 0, len(parts))
		for _, p := range parts {
			words = append(words, p)
		}

		res := solve(s, words)
		fmt.Fprintln(out, res)
	}
}
