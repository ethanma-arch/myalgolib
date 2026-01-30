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

func solve(nums []int) bool {
	n := len(nums)
	maxReach := 0

	for i := 0; i < n; i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, nums[i]+i)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readLineAsInts(r *bufio.Reader) ([]int, error) {
	line, err := r.ReadString('\n')
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
	if err != nil && len(nums) == 0 {
		return
	}
	b := solve(nums)
	fmt.Fprintf(out, "%t\n", b)
}
