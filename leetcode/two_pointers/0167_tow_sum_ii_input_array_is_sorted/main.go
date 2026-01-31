package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"leetcode/common_lib"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n-1

	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}

func main() {
	defer out.Flush()

	for {
		line, err := common_lib.ReadNextValidLine(in)
		if err != nil {
			break
		}
		parts := strings.Fields(line)
		nums := make([]int, 0, len(parts))
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			nums = append(nums, num)
		}
		targetLine, _ := common_lib.ReadNextValidLine(in)
		targetVal, _ := strconv.Atoi(strings.TrimSpace(targetLine))
		res := solve(nums, targetVal)
		for _, v := range res {
			fmt.Fprintf(out, "%d", v)
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)
	}
}
