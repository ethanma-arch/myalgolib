// LeetCode 68 Text Justification
// 输入：第一行空格分隔的单词（即 []string 的一行表示），第二行 maxWidth
// 输出：每行一条 justified 后的字符串
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func solve(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return nil
	}
	var result []string
	i := 0
	for i < len(words) {
		j := i + 1
		totalLen := len(words[i])
		for j < len(words) && totalLen+1+len(words[j]) <= maxWidth {
			totalLen += 1 + len(words[j])
			j++
		}
		numWords := j - i
		if j == len(words) || numWords == 1 {
			// 最后一行或只有一个词：左对齐，末尾补空格到 maxWidth
			line := strings.Join(words[i:j], " ")
			line += strings.Repeat(" ", maxWidth-len(line))
			result = append(result, line)
		} else {
			// 完全对齐：totalLen 已含单词间单空格，再分配 (maxWidth-totalLen) 的额外空格
			extraTotal := maxWidth - totalLen
			numGaps := numWords - 1
			baseExtra := extraTotal / numGaps
			leftExtra := extraTotal % numGaps
			var sb strings.Builder
			for k := i; k < j; k++ {
				if k > i {
					n := 1 + baseExtra // 1 个必选空格 + 额外
					if (k - i - 1) < leftExtra {
						n++
					}
					sb.WriteString(strings.Repeat(" ", n))
				}
				sb.WriteString(words[k])
			}
			result = append(result, sb.String())
		}
		i = j
	}
	return result
}

func main() {
	defer out.Flush()

	for {
		//line1, err := readLineAsString(in)
		//if err != nil && len(line1) == 0 {
		//	break
		//}
		//line2, err2 := readLineAsString(in)
		//if err2 != nil && len(line2) == 0 {
		//	break
		//}
		//words := strings.Fields(line1)
		//maxWidth, _ := strconv.Atoi(line2)
		//lines := solve(words, maxWidth)
		//for _, s := range lines {
		//	fmt.Fprintln(out, s)
		//}
		//if err == io.EOF || err2 == io.EOF {
		//	break
		//}
		line1, err1 := readNextValidLine(in)
		if len(line1) == 0 {
			if err1 != nil {
				break
			}
		}

		line2, err2 := readNextValidLine(in)
		if len(line2) == 0 {
			if err2 != nil {
				break
			}
		}

		words := strings.Fields(line1)
		maxWidth, _ := strconv.Atoi(line2)
		res := solve(words, maxWidth)
		for _, s := range res {
			fmt.Fprintln(out, s)
		}
		fmt.Fprintln(out, " ")

		if err1 == io.EOF || err2 == io.EOF {
			break
		}
	}
}
