package main

import (
	"bufio"
	"fmt"
	"leetcode/common_lib"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(s, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return []int{}
	}

	var res []int

	var pCount, windowCount [26]int

	// 1. 初始化
	// 统计 p 的字符，以及 s 中第一个窗口的字符
	for i := 0; i < pLen; i++ {
		pCount[p[i]-'a']++
		windowCount[s[i]-'a']++
	}

	if pCount == windowCount {
		res = append(res, 0)
	}

	// 2. 开始滑动窗口
	// i 是滑动前窗口的最左侧下标
	// i+1 是滑动后窗口的最左侧下标
	for i := 0; i < sLen-pLen; i++ {
		// [出窗口]: 移除当前窗口最左边的字符 s[i]
		windowCount[s[i]-'a']--
		// [入窗口]: 加入新窗口最右边的字符 s[i+pLen]
		windowCount[s[i+pLen]-'a']++

		// [检查]: Go 语言支持数组直接比较
		// 如果计数完全一致，说明是异位词
		if pCount == windowCount {
			res = append(res, i+1)
		}
	}

	return res
}

func main() {
	defer out.Flush()

	for {
		line1, err1 := common_lib.ReadNextValidLine(in)
		if len(line1) == 0 && err1 != nil {
			return
		}
		line1 = strings.TrimSpace(line1)
		line2, err2 := common_lib.ReadNextValidLine(in)
		if len(line2) == 0 && err2 != nil {
			return
		}
		line2 = strings.TrimSpace(line2)

		res := solve(line1, line2)
		fmt.Fprintln(out, res)
	}
}
