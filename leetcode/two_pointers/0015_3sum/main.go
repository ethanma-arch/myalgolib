package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func solve(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			return res
		}

		for i > 0 && nums[i] == nums[i-1] {
			i++
		}

		l, r := i+1, n-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})

				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				// 记得增加左右指针
				l++
				r--
			}
		}
	}
	return res
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
