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

func solve(ratings []int) (res int) {
	n := len(ratings)
	candies := make([]int, n)
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		// 注意边界条件
		if ratings[i] > ratings[i+1] {
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	for _, v := range candies {
		res += v
	}

	return res
}

func main() {
	defer out.Flush()

	ratings, err := readLineAsInts(in)
	if len(ratings) == 0 && err != nil {
		return
	}

	res := solve(ratings)
	fmt.Fprintf(out, "%d\n", res)
}
