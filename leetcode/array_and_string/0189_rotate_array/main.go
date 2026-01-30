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

func solve(nums []int, k int) {
	n := len(nums)

	if n == 1 {
		return
	}

	k = k % n

	swap(nums, 0, n-1)
	swap(nums, 0, k-1)
	swap(nums, k, n-1)
}

func swap(nums []int, a, b int) {
	for a < b {
		nums[a], nums[b] = nums[b], nums[a]
		a++
		b--
	}
}

func readLineAsInts(r *bufio.Reader) ([]int, error) {
	// 读取直到换行符
	line, err := r.ReadString('\n')

	// 去除首尾空白（包括换行符）
	line = strings.TrimSpace(line)

	// 如果读取到空行且 err 不为空（EOF），返回 nil
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

	// 1. 读取第一行：数组
	nums, err := readLineAsInts(in)
	if err != nil && len(nums) == 0 {
		return
	}

	// 2. 读取第二行：k 值
	// 虽然 k 只有一个数，但为了方便，我们还是按行读，取第一个即可
	kLine, err := readLineAsInts(in)
	if err != nil && len(kLine) == 0 {
		return
	}
	k := kLine[0]

	// 3. 执行逻辑
	solve(nums, k)

	// 4. 输出结果
	for i, v := range nums {
		fmt.Fprintf(out, "%d", v)
		if i < len(nums)-1 {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}
