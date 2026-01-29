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

func solve(nums []int) int {
	n := len(nums)
	slow, fast := 1, 1

	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func readLineAsInts(r *bufio.Reader) ([]int, error) {
	// 1. 读取整行直到遇到换行符 '\n'
	line, err := r.ReadString('\n')
	if err != nil && len(line) == 0 {
		return nil, err
	}

	// 2. 去除首尾空白字符 (包括 \n, \r, 空格)
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}
	// 3. 使用 Fields 按空白分割 (Fields 会自动忽略连续的空格)
	// 不要用 strings.Split(line, " ")，因为如果两个数字间有多个空格会出错
	parts := strings.Fields(line)

	// 4. 转换
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
