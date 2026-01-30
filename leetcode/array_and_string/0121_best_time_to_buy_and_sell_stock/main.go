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

func solve(nums []int) (maxProfit int) {
	n := len(nums)
	dp := make([]int, n)
	lowest := nums[0]

	// dp[i] = max(dp[i-1], nums[i]-lowest)
	for i := 1; i < n; i++ {
		if nums[i] < lowest {
			lowest = nums[i]
		}
		dp[i] = max(dp[i-1], nums[i]-lowest)
	}

	return dp[n-1]
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
	result := solve(nums)

	fmt.Fprintf(out, "%d\n", result)
}
