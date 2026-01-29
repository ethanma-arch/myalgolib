package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func init(){
	in.Split(bufio.ScanWords)
}

func solve(nums []int, val int) int {
	n := len(nums)
	slow, fast := 0, 0
	
	for fast < n {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func readInt() int {
	in.Scan()
	num, _ := strconv.Atoi(in.Text())
	return num
}

func main(){
	defer out.Flush()

	n := readInt()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	val := readInt()
	
	newLength := solve(nums, val)

	fmt.Fprintf(out, "%d\n", newLength)

	for i := 0; i < newLength; i++ {
		fmt.Fprintf(out, "%d", nums[i])
		if i < newLength-1 {
			fmt.Fprintf(out, " ")
		}
	}
	fmt.Fprintf(out, "\n")
}
