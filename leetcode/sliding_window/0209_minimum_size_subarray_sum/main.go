package main

import (
	"bufio"
	"fmt"
	"leetcode/common_lib"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func solve(target int, nums []int) int {
	n := len(nums)
	l, r := 0, 0
	currentSum := 0
	res := n

	for r < n {
		currentSum += nums[r]
		r++

		for currentSum >= target {
			if res > r-l {
				res = r - l
			}
			currentSum -= nums[l]
			l++
		}
	}

	if res == n {
		return 0
	}

	return res
}

func main() {
	defer out.Flush()

	for {
		line1, err1 := common_lib.ReadNextValidLine(in)
		if err1 != nil {
			break
		}
		target, _ := strconv.Atoi(line1)

		line2, err2 := common_lib.ReadNextValidLine(in)
		if len(line2) == 0 && err2 != nil {
			break
		}
		parts := strings.Fields(line2)
		nums := make([]int, 0, len(parts))
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			nums = append(nums, num)
		}

		res := solve(target, nums)
		fmt.Fprintln(out, res)
	}
}
