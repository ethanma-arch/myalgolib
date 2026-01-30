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

func solve(nums []int) (steps int) {
	n := len(nums)
	maxReach := 0
	end := 0

	for i := 0; i < n-1; i++ {
		maxReach = max(maxReach, nums[i]+i)
		if i == end {
			steps++
			end = maxReach
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readLineAsInts(r *bufio.Reader) ([]int, error) {
	line, err := r.ReadString('\n')
	line = strings.TrimSpace(line)
	if err != nil && len(line) == 0 {
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
	if len(nums) == 0 && err != nil {
		return
	}
	n := solve(nums)
	fmt.Fprintf(out, "%d\n", n)
}
