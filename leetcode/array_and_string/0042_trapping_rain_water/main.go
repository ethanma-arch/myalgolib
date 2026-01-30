package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(nums []int) int {
	n := len(nums)
	lMax, rMax := nums[0], nums[n-1]
	lp, rp := 1, n-2
	res := 0

	for lp <= rp {
		if lMax < rMax {
			if nums[lp] < lMax {
				res += lMax - nums[lp]
			} else {
				lMax = nums[lp]
			}
			lp++
		} else {
			if nums[rp] < rMax {
				res += rMax - nums[rp]
			} else {
				rMax = nums[rp]
			}
			rp--
		}
	}
	return res
}

func main() {
	defer out.Flush()

	for {
		nums, err := readLineAsInts(in)
		if len(nums) == 0 {
			if err != nil {
				break
			}
			continue
		}

		ans := solve(nums)
		fmt.Fprintf(out, "%d\n", ans)
		if err == io.EOF {
			break
		}
	}
}
