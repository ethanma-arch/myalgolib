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

func solve(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1

	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}

func main() {
	defer out.Flush()
	nums, err := readLineAsInts(in)

	if len(nums) == 0 && err != nil {
		return
	}

	res := solve(nums)
	for i, v := range res {
		fmt.Fprintf(out, "%d", v)
		if i < len(res)-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}
