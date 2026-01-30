// 快速排序，ACM 格式：stdin 一行空格分隔整数，stdout 输出排序后的一行。
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

// partition Lomuto：以 nums[r] 为 pivot，返回 pivot 最终下标
func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quicksort(nums, l, p-1)
	quicksort(nums, p+1, r)
}

func readLineAsInts() ([]int, error) {
	line, err := in.ReadString('\n')
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

	nums, err := readLineAsInts()
	if err != nil || len(nums) == 0 {
		return
	}

	quicksort(nums, 0, len(nums)-1)

	for i, v := range nums {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
