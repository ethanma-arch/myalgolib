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

func readNextValidLine(r *bufio.Reader) (string, error) {
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			return line, nil
		}
		if err != nil {
			return "", err
		}
	}
}

func solve(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	water := 0

	for l < r {
		water = max(water, min(nums[l], nums[r])*(r-l))
		if nums[l] > nums[r] {
			r--
		} else {
			l++
		}
	}
	return water
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	defer out.Flush()

	for {
		line, err := readNextValidLine(in)
		if err != nil {
			break
		}
		parts := strings.Fields(line)
		nums := make([]int, 0, len(parts))

		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			nums = append(nums, num)
		}

		res := solve(nums)
		fmt.Fprintln(out, res)
	}
}
