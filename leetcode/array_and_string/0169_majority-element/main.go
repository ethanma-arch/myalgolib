package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(nums []int) int {
	n := len(nums)
	candidate := nums[0]
	count := 1

	for i := 1; i < n; i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1 // 这里建议重置为 1
			continue  // 重置后直接进入下一轮
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func readLineAsInts(r *bufio.Reader) ([]int, error) {
	line, err := r.ReadString('\n')

	line = strings.TrimSpace(line)

	if len(line) == 0 && err != nil {
		return nil, err
	}

	parts := strings.Fields(line)

	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		num, _ := strconv.Atoi(p)
		nums = append(nums, num)
	}
	return nums, nil
}

func main() {
	defer out.Flush()

	nums, err := readLineAsInts(in)
	if err != nil || len(nums) == 0 {
		return
	}
	result := solve(nums)
	fmt.Fprintf(out, "%d\n", result)
}
