// Package common_lib 提供 ACM 读入等通用方法
package common_lib

import (
	"bufio"
	"strconv"
	"strings"
)

// ReadNextValidLine 从 r 读取直到遇到非空行或 EOF，跳过空行
func ReadNextValidLine(r *bufio.Reader) (string, error) {
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

func StringPartsToInts(line string) []int {
	parts := strings.Fields(line)
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		nums = append(nums, num)
	}
	return nums
}
