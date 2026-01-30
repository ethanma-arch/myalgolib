// 堆排序，ACM 格式：stdin 一行空格分隔整数，stdout 输出排序后的一行。
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

func heapify(nums []int, n, i int) {
	largest := i
	left, right := 2*i+1, 2*i+2

	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}

func buildHeap(nums []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

func heapsort(nums []int) {
	n := len(nums)
	buildHeap(nums, n)
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
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

	heapsort(nums)

	for i, v := range nums {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
