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
	// 必须去空格，否则会保留 \n
	line = strings.TrimSpace(line)
	if len(line) == 0 && err != nil {
		return "", err
	}
	return line, nil
}

func solve(s string) string {
	bytes := []byte(s)
	n := len(bytes)
	slow := 0

	for i := 0; i < n; i++ {
		if bytes[i] == ' ' {
			continue
		}

		if slow != 0 {
			bytes[slow] = ' '
			slow++
		}

		for i < n && bytes[i] != ' ' {
			bytes[slow] = bytes[i]
			slow++
			i++
		}
	}

	bytes = bytes[:slow]
	reverseBytes(bytes, 0, len(bytes)-1)

	start := 0
	// 注意点：这里 i <= len(bytes)，为了处理最后一个单词的结束边界
	for i := 0; i <= len(bytes); i++ {
		// 当 i 到达末尾 或者 遇到空格时，说明上一个单词结束了
		if i == len(bytes) || bytes[i] == ' ' {
			reverseBytes(bytes, start, i-1)
			start = i + 1
		}
	}
	return string(bytes)
}

func reverseBytes(bytes []byte, i, j int) {
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
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
