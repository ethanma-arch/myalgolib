package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func init() {
	// bufio.Scanner 默认按行读取。
	// 对于 ACM 模式读取数字，必须设置为按单词(空格/换行)分割。
	in.Split(bufio.ScanWords)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func readInt() int {
	in.Scan()
	num, _ := strconv.Atoi(in.Text())
	return num
}

func main() {
	defer out.Flush()

	if !in.Scan() {
		return
	}
	mStr := in.Text()
	m, _ := strconv.Atoi(mStr)

	n := readInt()

	nums1 := make([]int, m+n)
	for i := 0; i < m; i++ {
		nums1[i] = readInt()
	}

	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		nums2[i] = readInt()
	}

	merge(nums1, m, nums2, n)

	for i := 0; i < len(nums1); i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, nums1[i])
	}
	fmt.Fprintln(out)
}
