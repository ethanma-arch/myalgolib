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

func solve(nums []int) (res int) {
	n := len(nums)
	slow, fast := 2, 2

	for fast < n {
		if nums[fast] != nums[fast-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func readLineAsInts(r *bufio.Reader) ([]int, error) {
	line, err := r.ReadString('\n')
	if err != nil && len(line) == 0 {
		return nil, err
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}
	parts := strings.Fields(line)
	nums := make([]int, 0, len(parts))
	for _, s := range parts {
		num, _ := strconv.Atoi(s)
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

	newLength := solve(nums)

	fmt.Fprintf(out, "%d\n", newLength)
	for i := 0; i < newLength; i++ {
		fmt.Fprint(out, nums[i])
		if i < newLength-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}
